package controllers

import (
	"fmt"
	//"time"

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
	o := orm.NewOrm()
	o.Using("db")
	var maps []orm.Params
	//nowtime := time.Now().Format("2006-01-02 15:04:05")
	num, err := o.Raw("select * from Token WHERE timestamp between DATE_SUB('2018-03-19 13:00:00',INTERVAL 1 HOUR) and '2018-03-19 13:00:00'").Values(&maps)
	//num, err := o.Raw("select * from Token WHERE timetamp > DATE_SUB('2018-03-19 12:00:00',INTERVAL 1 HOUR)").Values(&maps)
	if err != nil {
		fmt.Println("err!")
	}
	fmt.Println("num", num, maps)
	this.ajaxList("获取数据成功", 0, num, maps)
}
