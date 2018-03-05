package controllers

import (
	"EtherscanPj/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
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
func (this *MainController) SearchResult() {
	//o := orm.NewOrm()
	//var maps []orm.Params
	var url string
	list := make([]interface{}, 0)
	//获取apikey
	apikey := beego.AppConfig.String("apikey")
	//获取address
	address := this.Input().Get("address")
	fmt.Println("address:", address)
	var balance models.Balance
	if address != "" {
		url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%s&tag=latest&apikey=%s", address, apikey)
	} else {
		url = "https://api.etherscan.io/api?module=account&action=balance&address=0x6f46cf5569aefa1acc1009290c8e043747172d89&tag=latest&apikey=YourApiKeyToken"
	}
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

	//maps = balance
	//cstr, _ := json.Marshal(&balance)
	//log.Println(url, c, string(cstr))
	//balance=models.Balance{}

	this.Data["balance"] = balance.Result
	//list["data"] = string(res)
	list = append(list, balance)
	//插入数据库
	//num, err := o.Insert(&balance)
	//if err != nil {
	//	fmt.Println("新增失败", err.Error())
	//}
	//fmt.Println("num:", num)
	this.ajaxList("查询成功", 200, 1, list)

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
