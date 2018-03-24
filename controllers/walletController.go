package controllers

import (
	"EtherscanPj/models"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type WalletController struct {
	BaseController
}

func (this *WalletController) Get() {
	//wallet
	o := orm.NewOrm()
	moniter := new(models.Monitior)
	num, err := o.QueryTable(moniter).Count()
	if err != nil {
		fmt.Println("err!")
	}
	this.Data["count"] = num

	this.TplName = "wallet.tpl"
}
func (this *WalletController) GetWalletMonitor() {
	//获取钱包地址
	wallet_address := this.Input().Get("wallet_address")
	this.Data["wallet_address"] = wallet_address
	//获取投资者名字
	o := orm.NewOrm()
	var lists []orm.ParamsList
	monitor := new(models.Monitior)
	num, err := o.QueryTable(monitor).Filter("Address", wallet_address).ValuesList(&lists, "Name", "Phone")
	if err != nil {
		fmt.Println("err!")
	}
	fmt.Println("num", num, lists[0][0])
	//	for _, row := range lists {
	//		fmt.Println("Name:Age:", row[0], row[1])
	//	}
	this.Data["name"] = lists[0][0]
	this.Data["phone"] = lists[0][1]
	this.TplName = "wallet_monitor.tpl"
}

//获取钱包监控数据
func (this *WalletController) GetWalletMonitorData() {
	o := orm.NewOrm()
	//o.Using("db")
	var maps []orm.Params
	var reslut []map[string]interface{}
	out := make(map[string]interface{})
	num, err := o.Raw("select * from data").Values(&maps)
	if err != nil {
		fmt.Println("err!")
	}
	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
		out["token"] = "SAY"
		reslut = append(reslut, out)
	}
	fmt.Println("num", num, reslut)
	this.ajaxList("获取数据成功", 0, num, reslut)
}

func (this *WalletController) GetIncrease() {

	this.TplName = "wallet_increase.tpl"
}

func (this *WalletController) GetPie() {

	this.TplName = "wallet_piechart.tpl"
}
