package controllers

import (
	"EtherscanPj/models"
	"fmt"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type PmController struct {
	BaseController
}

func (this *PmController) Get() {
	if this.GetSession("islogin") != 1 {
		fmt.Println("未登录")
		this.Redirect("/login", 302)
	}
	//pm
	o := orm.NewOrm()
	moniter := new(models.Monitior)
	num, err := o.QueryTable(moniter).Count()
	if err != nil {
		fmt.Println("err!")
	}
	this.Data["count"] = num
	this.TplName = "pm.tpl"
}

func (this *PmController) RealTimeData() {

	if this.GetSession("islogin") != 1 {
		fmt.Println("未登录")
		this.Redirect("/login", 302)
	}
	this.TplName = "realtime_data.tpl"
}

//获取实时数据
func (this *PmController) GetRealTimeData() {
	//get type
	time_type := this.Input().Get("type")
	o := orm.NewOrm()
	//o.Using("db")
	var maps []orm.Params
	var count int64
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("nowtime", nowtime)
	if time_type == "hour" {
		num, err := o.Raw("select * from data WHERE timestamp between DATE_SUB(?,INTERVAL 1 HOUR) and ?", nowtime, nowtime).Values(&maps)
		//num, err := o.Raw("select * from data WHERE timetamp > DATE_SUB('2018-03-19 12:00:00',INTERVAL 1 HOUR)").Values(&maps)
		if err != nil {
			fmt.Println("err!")
		}
		count = num
		fmt.Println("num", num, time_type)
	} else if time_type == "day" {
		num, err := o.Raw("select * from data WHERE timestamp between DATE_SUB(?,INTERVAL 1 DAY) and ?", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("err!")
		}
		count = num
		fmt.Println("num", num, time_type)
	} else if time_type == "week" {
		num, err := o.Raw("select * from data WHERE timestamp between DATE_SUB(?,INTERVAL 1 WEEK) and ?", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("err!")
		}
		count = num
		fmt.Println("num", num, time_type)
	} else if time_type == "month" {
		num, err := o.Raw("select * from data WHERE timestamp between DATE_SUB(?,INTERVAL 1 MONTH) and ?", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("err!")
		}
		count = num
		fmt.Println("num", num, time_type)
	} else {
		num, err := o.Raw("select * from data").Values(&maps)
		if err != nil {
			fmt.Println("err!")
		}
		count = num
		fmt.Println("num", num)
	}
	this.ajaxList("获取数据成功", 0, count, maps)
}
