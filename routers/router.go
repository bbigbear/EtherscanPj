package routers

import (
	"EtherscanPj/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.StockholderController{})
	beego.Router("/getearlywarn", &controllers.StockholderController{}, "*:GetEarlyWarn")
	beego.Router("/getnotifcationmessage", &controllers.StockholderController{}, "*:GetNotifcationMessage")
	beego.Router("/getnotifcationdata", &controllers.StockholderController{}, "*:GetNotifcationData")
	beego.Router("/getstockholder", &controllers.StockholderController{}, "*:GetStockHolder")
	beego.Router("/getstockholderdata", &controllers.StockholderController{}, "*:GetStockHolderData")
	beego.Router("/delnotifcationdata", &controllers.StockholderController{}, "*:DelNotifcationData")
	beego.Router("/test", &controllers.MainController{})
	beego.Router("/search/balance", &controllers.MainController{}, "*:GetBalanceResult")
	beego.Router("/search/transaction", &controllers.MainController{}, "*:GetTransactionResult")
}
