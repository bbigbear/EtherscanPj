package controllers

import (
	"fmt"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type PmController struct {
	BaseController
}

func (this *PmController) Get() {

	this.TplName = "realtime_data.tpl"
}

func (this *PmController) RealTimeData() {

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
