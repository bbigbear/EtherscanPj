package routers

import (
	"EtherscanPj/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.StockholderController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/addmonitor", &controllers.StockholderController{}, "*:AddMonitor")
	beego.Router("/addmonitor_action", &controllers.StockholderController{}, "post:AddMonitorAction")
	beego.Router("/getmonitordata", &controllers.StockholderController{}, "*:GetMonitorData")
	beego.Router("/delmonitordata", &controllers.StockholderController{}, "post:DelMonitorData")
	beego.Router("/getearlywarn", &controllers.StockholderController{}, "*:GetEarlyWarn")
	beego.Router("/earlywarnaction", &controllers.StockholderController{}, "post:EarlyWarnAction")
	beego.Router("/getnotifcationmessage", &controllers.StockholderController{}, "*:GetNotifcationMessage")
	beego.Router("/getnotifcationdata", &controllers.StockholderController{}, "*:GetNotifcationData")
	beego.Router("/getstockholder", &controllers.StockholderController{}, "*:GetStockHolder")
	beego.Router("/getstockholderdata", &controllers.StockholderController{}, "*:GetStockHolderData")
	beego.Router("/delnotifcationdata", &controllers.StockholderController{}, "*:DelNotifcationData")
	beego.Router("/test", &controllers.MainController{})
	beego.Router("/search/balance", &controllers.MainController{}, "*:GetBalanceResult")
	beego.Router("/search/transaction", &controllers.MainController{}, "*:GetTransactionResult")

	beego.Router("/pm", &controllers.PmController{})
	beego.Router("/realtimedata", &controllers.PmController{}, "*:RealTimeData")
	beego.Router("/getrealtimedata", &controllers.PmController{}, "*:GetRealTimeData")

	beego.Router("/wallet", &controllers.WalletController{})
	beego.Router("/getwalletmonitordata", &controllers.WalletController{}, "*:GetWalletMonitorData")
	beego.Router("/walletmonitor", &controllers.WalletController{}, "*:GetWalletMonitor")
	beego.Router("/walletincrease", &controllers.WalletController{}, "*:GetIncrease")
	beego.Router("/walletpie", &controllers.WalletController{}, "*:GetPie")
	beego.Router("/getwalletpiedata", &controllers.WalletController{}, "*:GetPieData")

}
