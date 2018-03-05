package routers

import (
	"EtherscanPj/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/search", &controllers.MainController{}, "*:SearchResult")
}
