package routers

import (
	"EtherscanPj/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/search/balance", &controllers.MainController{}, "*:GetBalanceResult")
	beego.Router("/search/transaction", &controllers.MainController{}, "*:GetTransactionResult")
}
