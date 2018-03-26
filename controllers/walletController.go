package controllers

import (
	"EtherscanPj/models"
	"fmt"
	"sort"
	"strconv"

	"time"

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
	//获取钱包地址
	wallet_address := this.Input().Get("wallet_address")
	this.Data["wallet_address"] = wallet_address
	o := orm.NewOrm()
	//o.Using("db")
	var maps []orm.Params
	var reslut []map[string]interface{}

	num, err := o.Raw("select * from data where from_address=? or to_address=?", wallet_address, wallet_address).Values(&maps)
	if err != nil {
		fmt.Println("err!")
		this.ajaxMsg("获取数据失败", MSG_ERR_Param)
	}
	fmt.Println("maps", maps)
	for _, m := range maps {
		fmt.Println("m", m)
		out := make(map[string]interface{})
		for k, v := range m {
			//fmt.Println("k", k)
			//fmt.Println("v", v)
			out[k] = v
		}
		out["token"] = "SAY"
		fmt.Println("out", out)
		reslut = append(reslut, out)
		fmt.Println("reslut", reslut)
	}
	fmt.Println("num", num, reslut)
	this.ajaxList("获取数据成功", 0, num, reslut)
}

func (this *WalletController) GetIncrease() {
	//获取一周的数据
	o := orm.NewOrm()
	var maps []orm.Params
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	num, err := o.Raw("select * from wallet_num WHERE timestamp between DATE_SUB(?,INTERVAL 1 WEEK) and ?", nowtime, nowtime).Values(&maps)
	if err != nil {
		fmt.Println("err!")
	}
	fmt.Println("num", num, maps)
	this.Data["maps"] = maps

	this.TplName = "wallet_increase.tpl"
}

func (this *WalletController) GetPie() {
	//获取占比前10名
	o := orm.NewOrm()
	data := new(models.Data)
	var maps []orm.Params
	var list []float64
	var reslut_list []float64
	num, err := o.QueryTable(data).Values(&maps)
	if err != nil {
		fmt.Println("err!")
	}
	fmt.Println("num", num)
	for _, m := range maps {
		if m["Percent"].(string) != "" {
			v, err := strconv.ParseFloat(m["Percent"].(string), 64)
			if err != nil {
				fmt.Println("err")
			}
			list = append(list, v)
		}
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(list)))
	list1 := this.RemoveRepBySlice(list)
	//sort.Float64s(list)
	fmt.Println("list", list1)
	for i := 0; i < 10; i++ {
		reslut_list = append(reslut_list, list1[i])
	}
	this.Data["list1"] = reslut_list
	//根据

	this.TplName = "wallet_piechart.tpl"
}

func (this *WalletController) GetPieData() {
	//获取占比前10名
	o := orm.NewOrm()
	data := new(models.Data)
	//下次改为data
	notify := new(models.Notifcation)
	var maps []orm.Params
	var list []float64
	//var reslut_list []float64
	var reslut_maps []map[string]interface{}
	num, err := o.QueryTable(data).Values(&maps)
	if err != nil {
		fmt.Println("err!")
	}
	fmt.Println("num", num)
	for _, m := range maps {
		if m["Percent"].(string) != "" {
			v, err := strconv.ParseFloat(m["Percent"].(string), 64)
			if err != nil {
				fmt.Println("err")
			}
			list = append(list, v)
		}
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(list)))
	list1 := this.RemoveRepBySlice(list)
	//sort.Float64s(list)
	fmt.Println("list", list1)
	for i := 0; i < 10; i++ {
		//reslut_list = append(reslut_list, list1[i])
		num, err := o.QueryTable(notify).Filter("Percent", strconv.FormatFloat(list1[i], 'f', -1, 64)).Values(&maps)
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println("num", num)
		for _, m := range maps {
			out := make(map[string]interface{})
			for k, v := range m {
				out[k] = v
			}
			//fmt.Println("out", out)
			reslut_maps = append(reslut_maps, out)
			//fmt.Println("reslut", reslut)
		}
	}
	//this.Data["list1"] = reslut_list
	//根据
	this.ajaxList("获取数据成功", 0, num, reslut_maps)
}
