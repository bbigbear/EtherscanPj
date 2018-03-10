package controllers

import (
	"EtherscanPj/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
)

type StockholderController struct {
	beego.Controller
}

func (this *StockholderController) Get() {
	o := orm.NewOrm()
	var maps []orm.Params
	sh := new(models.Stockholder)
	num, err := o.QueryTable(sh).Values(&maps)
	if err != nil {
		fmt.Println("error")
	}
	//fmt.Println(maps)
	fmt.Println("get sh reslut num:", num)
	this.ajaxList("获取列表数据成功", 0, num, maps)
	//	this.TplName = "index.tpl"

}

func (this *StockholderController) GetEarlyWarn() {
	this.TimeTask()
	this.TplName = "early_warn.tpl"
}

func (this *StockholderController) GetNotifcationMessage() {
	this.TplName = "notifcation_message.tpl"
}

//ajax返回 列表
func (this *StockholderController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["message"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

//定时
func (this *StockholderController) TimeTask() {
	//定期上架
	fmt.Println("定时执行")
	//ordertime
	o := orm.NewOrm()
	sh := new(models.Stockholder)
	var maps []orm.Params
	var sh_info models.Stockholder
	//查询设置的数值
	//	ew := new(models.EarlyWarn)
	//	var maps_ew []orm.Params
	//	num, err := o.QueryTable(ew).Values(&maps_ew)
	//	for _, m := range maps_ew {

	//	}

	//每隔10分钟执行一次
	tk1 := toolbox.NewTask("tk1", " 0 */10 * * * *", func() error {
		fmt.Println("10分钟执行一次")
		//		nt := time.Now().Format("2016-01-02 15:04:05")
		//		s, err := time.ParseInLocation("2006-01-02 15:04:05", nt, time.Local)
		//		if err != nil {
		//			log4go.Stdout("转化秒数失败", err.Error())
		//		}
		num, err := o.QueryTable(sh).Values(&maps)
		if err != nil {
			fmt.Println("获取list失败")
		}
		fmt.Println("get list num", num)
		for _, m := range maps {
			//遍历列表中每一项
			var balance models.Balance
			url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=tokenbalance&contractaddress=0xa9ec9f5c1547bd5b0247cf6ae3aab666d10948be&address=%s&tag=latest&apikey=", m["ADDRESS"])
			//fmt.Println("url:", url)
			r, err := http.Get(url)
			if err != nil {
				fmt.Println("http.Get err", err)
			}
			res, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println("ioutil.ReadAll(r.Body) ", err)
			}
			defer r.Body.Close()
			fmt.Println(string(res))
			json.Unmarshal(res, &balance)
			//get value
			result, err := strconv.ParseFloat(m["NUM"].(string), 64)
			if err != nil {
				fmt.Println("get result err")
			}
			last_result, err := strconv.ParseFloat(balance.Result, 64)
			if err != nil {
				fmt.Println("get last_result err")
			}
			if result != last_result {
				//跟新数据库
				//				id, err := strconv.Atoi(m["ID"].(string))
				//				if err != nil {
				//					fmt.Println("get id err")
				//				}
				id := m["ID"].(int64)
				sh_info.ID = id
				sh_info.NUM = balance.Result
				num, err := o.Update(&sh_info, "NUM")
				if err != nil {
					fmt.Println("get num", num)
				}
				var value, value_percent float64

				if last_result > result {
					value = (last_result - result) / 1000000000000000000
					value_percent = value / result * 100

				} else {
					value = (result - last_result) / 1000000000000000000
					value_percent = value / result * 100
				}
				if (value > 0.1) || (value_percent > 10) {
					//根据地址获取网站交易信息
					var transaction models.Transaction
					tx_url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=1&sort=asc&apikey=", m["ADDRESS"])
					fmt.Println("tx_url:", tx_url)
					r, err := http.Get(tx_url)
					if err != nil {
						fmt.Println("http.Get err", err)
					}
					res, err := ioutil.ReadAll(r.Body)
					if err != nil {
						fmt.Println("ioutil.ReadAll(r.Body) ", err)
					}
					defer r.Body.Close()
					//fmt.Println(string(res))
					json.Unmarshal(res, &transaction)
					//时间戳转日期
					dataTimeStr := time.Unix(transaction.TimeStamp, 0).Format("2006-01-02 15:04:05") //设置时间戳 使用模板格式化为日期字符串
					//交易超过0.1或者超过30%，存入消息中
					var nm models.Notifcation
					nm.Target = m["ADDRESS"].(string)
					nm.Style = "单笔交易"
					nm.Num = value
					nm.Percent = value_percent
					nm.Time = dataTimeStr
					nm.Hash = transaction.Hash
					//插入数据库
					num, err := o.Insert(&nm)
					if err != nil {
						fmt.Println("新增失败", err.Error())
					}
					fmt.Println("insert num:", num)
				}
			}
		}
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()

}
