package controllers

import (
	"EtherscanPj/models"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	//"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/shopspring/decimal"
)

type StockholderController struct {
	BaseController
}

func (this *StockholderController) Get() {

	//this.StartNotificationTask()
	//this.TplName = "index.tpl"
	this.TplName = "early_warn.tpl"
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
	this.TplName = "early_warn.tpl"
}

func (this *StockholderController) EarlyWarnAction() {
	//获取action
	fmt.Println("通知状态按钮")
	var status models.Status
	//list := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &status)
	status.Id = 1
	o := orm.NewOrm()
	num, err := o.Update(&status, "Status", "Svalue")
	if err != nil {
		fmt.Println("通知状态更新失败err", err.Error())
	}
	fmt.Println("更新成功num", num)
	this.ajaxMsg("保存成功", MSG_OK)
	return
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

//关闭通知
func (this *StockholderController) StopNotificationTask() {
	fmt.Println("关闭通知")
	toolbox.DeleteTask("tk1")
}

//开启通知
func StartNotificationTask() {
	fmt.Println("开启通知")
	//不过期
	bm, _ := cache.NewCache("memory", `{"interval":0}`)
	//15s 刷新一次
	tk1 := toolbox.NewTask("tk1", "0/15 * * * * *", func() error {
		o := orm.NewOrm()
		o.Using("db")
		var maps []orm.Params
		num, err := o.Raw("select * from Token order by id desc limit 1").Values(&maps)
		//err := o.Raw("select * from Token order by id desc limit 1", 1).QueryRow(&token)
		if err != nil {
			fmt.Println("err!")
		}
		if bm.IsExist("id") {
			if bm.Get("id") == maps[0]["id"].(string) {
				fmt.Println("相等,不更新")
				return nil
			} else {
				fmt.Println("不等,更新")
				bm.Delete("id")
			}
		} else {
			bm.Put("id", maps[0]["id"].(string), 0)
		}
		constact_address := maps[0]["contractAddress"].(string)
		to_address := maps[0]["toAddress"].(string)
		from_address := maps[0]["fromAddress"].(string)

		value, err := decimal.NewFromString(maps[0]["value"].(string))
		if err != nil {
			fmt.Println("err!")
		}
		//获取list1 的数据
		o1 := orm.NewOrm()
		monitor := new(models.Monitior)
		status := new(models.Status)
		var maps_monitor []orm.Params

		num1, err := o1.QueryTable(monitor).Values(&maps_monitor)
		if err != nil {
			fmt.Println("获取list失败")
		}
		fmt.Println("get list num", num1, maps_monitor)
		for _, m := range maps_monitor {
			address := m["Address"].(string)
			fmt.Println("address", address)
			if m["Contract"].(string) == constact_address {
				fmt.Println("contracct", m["Contract"].(string), constact_address)
				if address == to_address || address == from_address {
					//判断是否开启通知状态
					var status_data models.Status
					o1.QueryTable(status).Filter("Id", 1).One(&status_data)
					if err != nil {
						fmt.Println("查找status失败")
					}
					//将数据存入data中
					var token_data models.Data
					token_data.BlockNumber = maps[0]["blockNumber"].(string)
					token_data.ContractAddress = constact_address
					token_data.FromAddress = from_address
					t, _ := time.Parse("2006-01-02 15:04:05", maps[0]["timestamp"].(string))
					token_data.Timestamp = t
					token_data.ToAddress = to_address
					token_data.TransactionHash = maps[0]["transactionHash"].(string)
					token_data.Value = maps[0]["value"].(string)
					if status_data.Status == "on" {
						token_data.Status = "警告"
						//获取交易value
						value1, err := decimal.NewFromString(status_data.Svalue)
						if err != nil {
							fmt.Println("err!")
						}
						fmt.Println("num:", num, maps, value.LessThan(value1))
						if value.LessThan(value1) {
							fmt.Println("不进行推送")
						} else {
							fmt.Println("进行推送")
							var notify models.Notifcation
							notify.Hash = maps[0]["transactionHash"].(string)
							notify.Num = maps[0]["value"].(string)
							notify.Style = "单笔交易"
							t1, _ := time.Parse("2006-01-02 15:04:05", maps[0]["timestamp"].(string))
							notify.Time = t1
							notify.Target = address
							num, err := o1.Insert(&notify)
							if err != nil {
								fmt.Println("isnert err!")
							}
							fmt.Println("insert id", num)
						}
					}
					num2, err := o1.Insert(&token_data)
					if err != nil {
						fmt.Println("插入失败")
					}
					fmt.Println("插入data成功num", num2)

				}
			}
		}
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
}
