package main

import (
	//"EtherscanPj/controllers"
	"EtherscanPj/models"
	_ "EtherscanPj/routers"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	DBConnection()
	RegisterModel()
	//定时
	//controllers.StartNotificationTask()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	//controllers.StartNotificationTask()
	//controllers.WriteAddressNum()
	//controllers.ScanData()
	//controllers.UpdateBalance()
	//controllers.UpdatePercent()
	//controllers.UpdateTransactionStatus()

	beego.Run()
}

func DBConnection() {
	fmt.Println("初始化数据库")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:qwe!23@/etherscan?charset=utf8", 30, 30)
	orm.RegisterDataBase("db", "mysql", "root:h0bc1NUm@tcp(47.100.38.37:3306)/blockchain?charset=utf8", 30, 30)
}

func RegisterModel() {
	fmt.Println("注册数据库模型")
	orm.RegisterModel(new(models.Balance), new(models.Stockholder), new(models.Notifcation), new(models.Monitior), new(models.Data), new(models.Status), new(models.WalletNum))

}
