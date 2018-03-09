package controllers

import (
	"EtherscanPj/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//o := orm.NewOrm()
	//	var balance models.Balance
	//	url := "https://api.etherscan.io/api?module=account&action=balance&address=0x6f46cf5569aefa1acc1009290c8e043747172d89&tag=latest&apikey=YourApiKeyToken"
	//	r, err := http.Get(url)
	//	if err != nil {
	//		fmt.Println("http.Get err", err)
	//	}
	//	res, err := ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		fmt.Println("ioutil.ReadAll(r.Body) ", err)
	//	}
	//	defer r.Body.Close()
	//	fmt.Println(string(res))
	//	json.Unmarshal(res, &balance)

	//	this.Data["balance"] = balance.Result

	//插入数据库
	//num, err := o.Insert(&balance)
	//if err != nil {
	//	fmt.Println("新增失败", err.Error())
	//}
	//fmt.Println("num:", num)

	this.TplName = "index.tpl"

}
func (this *MainController) GetTransactionResult() {
	//	o := orm.NewOrm()
	//	var maps []orm.Params
	var url string
	var address string
	//获取apikey
	apikey := beego.AppConfig.String("apikey")
	//获取address
	address = this.Input().Get("address")
	fmt.Println("address:", address)
	if address == "" {
		address = "0x6f46cf5569aefa1acc1009290c8e043747172d89"
	}
	//var transaction models.Transaction
	var transaction map[string]interface{}

	//url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%s&tag=latest&apikey=%s", address, apikey)
	url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=2&sort=asc&apikey=%s", address, apikey)

	fmt.Println("url:", url)
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err", err)
	}
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll(r.Body) ", err)
	}
	defer r.Body.Close()
	fmt.Println(string(res))
	json.Unmarshal(res, &transaction)
	//get address
	//	transaction.Address = address

	//查询
	//	t := new(models.Transaction)

	//	//插入数据库
	//	num, err := o.Insert(&transaction)
	//	if err != nil {
	//		fmt.Println("新增失败", err.Error())
	//	}
	//	fmt.Println("insert num:", num)

	//	//获取不重复数据
	//	num1, err := o.QueryTable(t).Filter("Address", address).Distinct().Values(&maps)
	//	if err != nil {
	//		fmt.Println("获取数据失败")
	//	}
	//	fmt.Println("get data num", num1)

	this.ajaxList("查询成功", 0, 1, transaction["result"])

}

func (this *MainController) GetBalanceResult() {
	o := orm.NewOrm()
	var maps []orm.Params
	var url string
	var address string
	//list := make([]interface{}, 0)
	//获取apikey
	apikey := beego.AppConfig.String("apikey")
	//获取address
	address = this.Input().Get("address")
	fmt.Println("address:", address)
	//	if address == "" {
	//		address = "0x6f46cf5569aefa1acc1009290c8e043747172d89"
	//	}
	var balance models.Balance

	url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=tokenbalance&contractaddress=0xa9ec9f5c1547bd5b0247cf6ae3aab666d10948be&address=%s&tag=latest&apikey=%s", address, apikey)
	fmt.Println("url:", url)
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err", err)
	}
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll(r.Body) ", err)
	}
	defer r.Body.Close()
	fmt.Println(string(res))
	json.Unmarshal(res, &balance)
	//get address
	balance.Address = address

	//查询
	b := new(models.Balance)
	exist := o.QueryTable(b).Filter("Address", address).Exist()
	if exist {
		num, err := o.Update(&balance)
		fmt.Println("updata num", num)
		if err != nil {
			fmt.Println("更新失败")
		}
	} else {
		//插入数据库
		num, err := o.Insert(&balance)
		if err != nil {
			fmt.Println("新增失败", err.Error())
		}
		fmt.Println("insert num:", num)
	}
	//获取数据
	num, err := o.QueryTable(b).Filter("Address", address).Values(&maps)
	if err != nil {
		fmt.Println("获取数据失败")
	}
	fmt.Println("get data num", num)
	//maps = balance
	//cstr, _ := json.Marshal(&balance)
	//log.Println(url, c, string(cstr))
	//balance=models.Balance{}

	this.Data["balance"] = balance.Result
	//list["data"] = string(res)
	//list = append(list, balance)

	this.ajaxList("查询成功", 0, num, maps)

}

//ajax返回 列表
func (this *MainController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["message"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}
