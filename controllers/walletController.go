package controllers

type WalletController struct {
	BaseController
}

func (this *WalletController) Get() {

	this.TplName = "wallet.tpl"
}

func (this *WalletController) GetIncrease() {

	this.TplName = "wallet_increase.tpl"
}

func (this *WalletController) GetPie() {

	this.TplName = "wallet_piechart.tpl"
}
