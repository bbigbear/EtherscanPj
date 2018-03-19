package controllers

import (
	"EtherscanPj/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/tidwall/gjson"
)

const (
	MSG_OK            = 200
	MSG_ERR_Param     = 400
	MSG_ERR_Verified  = 401
	MSG_ERR_Authority = 403
	MSG_ERR_Resources = 404
	MSG_ERR           = 500
)

type StockholderController struct {
	beego.Controller
}

func (this *StockholderController) Get() {

	this.TplName = "early_warn.tpl"
	//	this.TplName = "index.tpl"
}

//新增投资者
func (this *StockholderController) AddMonitor() {
	this.TplName = "add_monitor.tpl"
}

func (this *StockholderController) AddMonitorAction() {
	var m models.Monitior
	//	var result models.ResultMoitor
	list := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &m)
	//获取parms
	//	req := httplib.Post("http://testapi.friendfun.org:8088/api/addmonitoring")
	//	req.Param("userid", m.Userid)
	//	req.Param("contract", m.Contract)
	//	req.Param("address", m.Address)
	//	req.Param("time", m.Time)
	//	req.Param("value", m.Value)

	//	err := req.ToJSON(&result)
	//	if err != nil {
	//		fmt.Println("add monitor err", err.Error())
	//		this.ajaxMsg("新增失败", MSG_ERR_Resources)
	//	}
	o := orm.NewOrm()
	num, err := o.Insert(&m)
	if err != nil {
		fmt.Println("insert monitor err")
	}
	list["id"] = num
	this.ajaxList("新增成功", MSG_OK, num, list)
	return

}

func (this *StockholderController) GetEarlyWarn() {
	//获取action
	action := this.Input().Get("action")
	fmt.Println("earlywarn_action:", action)
	if action != "" {
		if action == "start" {
			sn := this.Input().Get("sn")
			sp := this.Input().Get("sp")
			//get float
			n, err := strconv.ParseFloat(sn, 64)
			if err != nil {
				fmt.Println("get sn err")
			}
			p, err := strconv.ParseFloat(sp, 64)
			if err != nil {
				fmt.Println("get sp err")
			}
			fmt.Println("earlywarn_info:", n, p)
			//this.StopTimeTask()
			//this.StartTimeTask(n, p)
		} else if action == "stop" {
			//this.StopTimeTask()
		}

	}
	fmt.Println("earlywarn_info:", action)
	this.TplName = "early_warn.tpl"
}

func (this *StockholderController) GetNotifcationMessage() {
	this.TplName = "notifcation_message.tpl"
}

func (this *StockholderController) GetStockHolder() {
	this.TplName = "investor_manage.tpl"
}

//获取成员list
func (this *StockholderController) GetStockHolderData() {
	fmt.Println("获取stockholder消息信息")
	o := orm.NewOrm()
	var maps []orm.Params
	st := new(models.Stockholder)
	query := o.QueryTable(st)
	//查询数据库
	num, err := query.Values(&maps)
	if err != nil {
		//log4go.Stdout("获取消息失败", err.Error())
		this.ajaxMsg("获取消息失败", MSG_ERR_Resources)
	}
	fmt.Println("get stockholder reslut num:", num)
	this.ajaxList("获取消息成功", 0, num, maps)
	return
}

//获取成员list
func (this *StockholderController) GetMonitorData() {
	fmt.Println("获取MonitorData消息信息")
	//	var m map[string]interface{}
	//	req := httplib.Get("http://testapi.friendfun.org:8088/api/getmonitorings")
	//	str, err := req.Bytes()
	//	if err != nil {
	//		fmt.Println("get monitor data err", err.Error())
	//	}
	//	fmt.Println(string(str))
	//	result := gjson.GetBytes(str, "monitorings")
	//	json.Unmarshal(str, &m)

	//	this.ajaxList("获取消息成功", 0, 1, m["monitorings"])
	o := orm.NewOrm()
	var maps []orm.Params
	m := new(models.Monitior)
	query := o.QueryTable(m)
	//查询数据库
	num, err := query.Values(&maps)
	if err != nil {
		//log4go.Stdout("获取投资者失败", err.Error())
		this.ajaxMsg("获取投资者失败", MSG_ERR_Resources)
	}
	fmt.Println("get monitor_list reslut num:", num)
	this.ajaxList("获取投资者成功", 0, num, maps)
	return
}

//删除
func (this *StockholderController) DelMonitorData() {
	fmt.Println("删除消息数据")
	o := orm.NewOrm()
	m := new(models.Monitior)
	//list := make(map[string]interface{})
	id := this.Input().Get("id")
	fmt.Println("del id:", id)
	idList := strings.Split(id, ",")
	fmt.Println("idList:", idList)
	id_len := len(idList) - 1
	var idIntList []int64
	for i := 0; i < id_len; i++ {
		idd, err := strconv.ParseInt(idList[i], 10, 64)
		if err != nil {
			//log4go.Stdout("delmulti string转int 失败", err.Error())
			fmt.Println("delmulti string转int 失败", err.Error())
		}
		idIntList = append(idIntList, idd)
	}
	fmt.Println("idIntList:", idIntList)
	num, err := o.QueryTable(m).Filter("Id__in", idIntList).Delete()
	if err != nil {
		//log4go.Stdout("删除消息失败", err.Error())
		this.ajaxMsg("删除消息失败", MSG_ERR_Resources)
	}
	fmt.Println("del multimonitor reslut num:", num)
	if num == 0 {
		this.ajaxMsg("删除消息失败", MSG_ERR_Param)
	}
	//list["data"] = maps
	this.ajaxMsg("删除消息成功", MSG_OK)
	return
}

//获取消息list
func (this *StockholderController) GetNotifcationData() {
	fmt.Println("获取消息信息")
	o := orm.NewOrm()
	var maps []orm.Params
	notifcation := new(models.Notifcation)
	query := o.QueryTable(notifcation)
	//查询数据库
	num, err := query.Values(&maps)
	if err != nil {
		//log4go.Stdout("获取消息失败", err.Error())
		this.ajaxMsg("获取消息失败", MSG_ERR_Resources)
	}
	fmt.Println("get notifcation_list reslut num:", num)
	this.ajaxList("获取消息成功", 0, num, maps)
	return
}

//删除
func (this *StockholderController) DelNotifcationData() {
	fmt.Println("删除消息数据")
	o := orm.NewOrm()
	notifcation := new(models.Notifcation)
	//list := make(map[string]interface{})
	id := this.Input().Get("id")
	fmt.Println("del id:", id)
	idList := strings.Split(id, ",")
	fmt.Println("idList:", idList)
	id_len := len(idList) - 1
	var idIntList []int64
	for i := 0; i < id_len; i++ {
		idd, err := strconv.ParseInt(idList[i], 10, 64)
		if err != nil {
			//log4go.Stdout("delmulti string转int 失败", err.Error())
			fmt.Println("delmulti string转int 失败", err.Error())
		}
		idIntList = append(idIntList, idd)
	}
	fmt.Println("idIntList:", idIntList)
	num, err := o.QueryTable(notifcation).Filter("Id__in", idIntList).Delete()
	if err != nil {
		//log4go.Stdout("删除消息失败", err.Error())
		this.ajaxMsg("删除消息失败", MSG_ERR_Resources)
	}
	fmt.Println("del multinotifcation reslut num:", num)
	if num == 0 {
		this.ajaxMsg("删除消息失败", MSG_ERR_Param)
	}
	//list["data"] = maps
	this.ajaxMsg("删除消息成功", MSG_OK)
	return
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

//ajax返回
func (this *StockholderController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

//定时
func (this *StockholderController) StartTimeTask(n float64, p float64) {
	//定期上架
	fmt.Println("定时执行")
	//获取apikey
	apikey := beego.AppConfig.String("apikey")
	//ordertime

	//查询设置的数值
	//	ew := new(models.EarlyWarn)
	//	var maps_ew []orm.Params
	//	num, err := o.QueryTable(ew).Values(&maps_ew)
	//	for _, m := range maps_ew {

	//	}

	//每隔20分钟执行一次
	tk1 := toolbox.NewTask("tk1", " 0 */20 * * * *", func() error {
		fmt.Println("20分钟执行一次")
		//		nt := time.Now().Format("2016-01-02 15:04:05")
		//		s, err := time.ParseInLocation("2006-01-02 15:04:05", nt, time.Local)
		//		if err != nil {
		//			log4go.Stdout("转化秒数失败", err.Error())
		//		}
		//获取list1 的数据
		o := orm.NewOrm()
		sh := new(models.Stockholder)
		var maps []orm.Params
		var sh_info models.Stockholder
		num, err := o.QueryTable(sh).Values(&maps)
		if err != nil {
			fmt.Println("获取list失败")
		}
		fmt.Println("get list num", num)
		for _, m := range maps {
			//遍历列表中每一项
			var balance models.Balance
			url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=tokenbalance&contractaddress=0xa9ec9f5c1547bd5b0247cf6ae3aab666d10948be&address=%s&tag=latest&apikey=%s", m["ADDRESS"], apikey)
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
			//fmt.Println(string(res))
			json.Unmarshal(res, &balance)
			//get value
			result, err := strconv.ParseFloat(m["NUM"].(string), 64)
			if err != nil {
				fmt.Println("get result err")
			}
			fmt.Println("result", result)
			last_result, err := strconv.ParseFloat(balance.Result, 64)
			if err != nil {
				fmt.Println("get last_result err")
			}
			fmt.Println("last_result", last_result)
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
				if (value > n) || (value_percent > p) {
					//根据地址获取网站交易信息
					//var transaction models.Transaction
					//var result_info models.Results
					tx_url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=&sort=asc&apikey=%s", m["ADDRESS"], apikey)
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
					//json.Unmarshal(res, &transaction)
					var hash string
					var timestamp int64
					result_info := gjson.GetBytes(res, "result")
					//fmt.Println("result_info", result_info)
					if result_info.Exists() {
						result1 := result_info.Array()
						length := len(result1)
						fmt.Println("len", length)
						time1 := result1[length-1].Get("timeStamp").String()
						hash = result1[length-1].Get("hash").String()
						fmt.Println("time1", time1)
						fmt.Println("hash", hash)
						timestamp, err = strconv.ParseInt(time1, 10, 64)
						if err != nil {
							fmt.Println("get timestamp err")
						}
						fmt.Println("timestamp", timestamp)
					}
					//打印transaction
					//fmt.Println("result_info", result_info)
					//时间戳转日期
					dataTimeStr := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05") //设置时间戳 使用模板格式化为日期字符串
					//交易超过0.1或者超过30%，存入消息中
					var nm models.Notifcation
					nm.Target = m["ADDRESS"].(string)
					nm.Style = "单笔交易"
					nm.Num = value
					nm.Percent = value_percent
					nm.Time = dataTimeStr
					nm.Hash = hash
					//插入数据库
					num, err := o.Insert(&nm)
					if err != nil {
						fmt.Println("新增失败", err.Error())
					}
					fmt.Println("insert num:", num)
				}
			}
			//sleep
			//			time.Sleep(0.5 * time.Second)
		}
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
}
func (this *StockholderController) StopTimeTask() {
	fmt.Println("删除定时")
	toolbox.DeleteTask("tk1")
	//toolbox.StopTask()
}
