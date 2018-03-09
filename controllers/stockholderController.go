package controllers

import (
	"EtherscanPj/models"
	"encoding/json"
	"fmt"
	"io/ioutil"

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
	ew := new(models.EarlyWarn)
	var maps_ew []orm.Params
	num, err := o.QueryTable(ew).Values(&maps_ew)
	for _, m := range maps_ew {

	}

	//每隔10分钟执行一次
	tk1 := toolbox.NewTask("tk1", " 0 */10 * * * *", func() error {
		//fmt.Println("tk1")
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
			url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=tokenbalance&contractaddress=0xa9ec9f5c1547bd5b0247cf6ae3aab666d10948be&address=%s&tag=latest&apikey=", m["ADDRESS"])
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
			if m["Result"] != balance.Result {
				//跟新数据库
				//if
			}
		}
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()

}
