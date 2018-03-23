package controllers

//import (
//	"fmt"

//	_ "github.com/Go-SQL-Driver/MySQL"
//	"github.com/astaxie/beego/orm"
//)

type WalletController struct {
	BaseController
}

func (this *WalletController) Get() {
	//	o := orm.NewOrm()
	//	o.Using("db")
	//	//o.Raw(select count(*) from Token)
	//	var maps []orm.Params
	//	num, err := o.Raw("select * from Token").Values(&maps)
	//	if err == nil && num > 0 {
	//		fmt.Println(len(maps)) // slene
	//	}

	this.TplName = "wallet.tpl"
}

func (this *WalletController) GetIncrease() {

	this.TplName = "wallet_increase.tpl"
}

func (this *WalletController) GetPie() {

	this.TplName = "wallet_piechart.tpl"
}
