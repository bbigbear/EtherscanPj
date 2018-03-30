package controllers

import (
	"EtherscanPj/models"
	"fmt"
	"sort"
	"strconv"

	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
	//"github.com/shopspring/decimal"
)

type WalletController struct {
	BaseController
}

func (this *WalletController) Get() {
	if this.GetSession("islogin") != 1 {
		fmt.Println("未登录")
		this.Redirect("/login", 302)
	}
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
	if this.GetSession("islogin") != 1 {
		fmt.Println("未登录")
		this.Redirect("/login", 302)
	}
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
	if this.GetSession("islogin") != 1 {
		fmt.Println("未登录")
		this.Redirect("/login", 302)
	}
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
	if this.GetSession("islogin") != 1 {
		fmt.Println("未登录")
		this.Redirect("/login", 302)
	}
	//获取占比前10名
	o := orm.NewOrm()
	data := new(models.Data)
	var maps []orm.Params
	var list []float64
	//var reslut_list []float64
	var name_list []string
	var reslut_name_list []map[string]interface{}
	//	var reslut_num_list []float64
	//notify := new(models.Notifcation)
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
	l := list[0:10]
	list1 := this.RemoveRepBySliceFloat(l)
	for i := 0; i < len(list1); i++ {
		//reslut_list = append(reslut_list, list1[i])
		num, err := o.QueryTable(data).Filter("Percent", strconv.FormatFloat(list1[i], 'f', -1, 64)).Values(&maps)
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println("num", num)
		for _, m := range maps {
			name_list = append(name_list, m["Name"].(string))
		}
	}
	fmt.Println("name_list", name_list)
	fmt.Println("l", l)
	for i := 0; i < 10; i++ {

		out := make(map[string]interface{})
		out["id"] = l[i]
		out["name"] = name_list[i]
		reslut_name_list = append(reslut_name_list, out)
	}
	this.Data["list"] = reslut_name_list
	//	}

	this.TplName = "wallet_piechart.tpl"
}

func (this *WalletController) GetPieData() {
	//获取占比前10名
	o := orm.NewOrm()
	data := new(models.Data)
	//下次改为data
	//notify := new(models.Notifcation)
	var maps []orm.Params
	var list []float64
	//var name_list []string
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
	list1 := this.RemoveRepBySliceFloat(list[0:10])
	//	//判断第10位是否为1
	//	var value_list []string
	//	var name_list []string
	//	if list[9] == float64(1) {
	//		num, err := o.QueryTable(notify).Filter("Percent", "1").Values(&maps)
	//		if err != nil {
	//			fmt.Println("err!")
	//		}
	//		fmt.Println("num", num)
	//		for _, m := range maps {
	//			value_list = append(value_list, m["Num"].(string))
	//			name_list = append(name_list, m["Target"].(string))
	//		}
	//		len1 := len(value_list)
	//		//冒泡法
	//		for i := 0; i < len1; i++ {
	//			for j := i + 1; j < len1; j++ {
	//				n1, err := decimal.NewFromString(value_list[i])
	//				if err != nil {
	//					fmt.Println("err!")
	//				}
	//				n2, err := decimal.NewFromString(value_list[j])
	//				if err != nil {
	//					fmt.Println("err!")
	//				}
	//				if n1.LessThan(n2) {
	//					value_list[i], value_list[j] = value_list[j], value_list[i]
	//					name_list[i], name_list[j] = name_list[j], name_list[i]
	//				}
	//			}
	//		}
	//		fmt.Println("value_list", value_list)
	//		fmt.Println("name_list", name_list)
	//		for i := 0; i < 10; i++ {
	//			//reslut_list = append(reslut_list, list1[i])
	//			num, err := o.QueryTable(notify).Filter("Percent", "1").Filter("Target", name_list[i]).Values(&maps)
	//			if err != nil {
	//				fmt.Println("err")
	//			}
	//			fmt.Println("num", num)
	//			for _, m := range maps {
	//				out := make(map[string]interface{})
	//				for k, v := range m {
	//					out[k] = v
	//				}
	//				//fmt.Println("out", out)
	//				reslut_maps = append(reslut_maps, out)
	//				//fmt.Println("reslut", reslut)
	//			}
	//		}

	//	} else {
	//list1 := this.RemoveRepBySlice(list)
	//sort.Float64s(list)
	//fmt.Println("list", list1)
	for i := 0; i < len(list1); i++ {
		//reslut_list = append(reslut_list, list1[i])
		num, err := o.QueryTable(data).Filter("Percent", strconv.FormatFloat(list1[i], 'f', -1, 64)).Values(&maps)
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

	//	}

	//this.Data["list1"] = reslut_list
	//根据
	this.ajaxList("获取数据成功", 0, num, reslut_maps)
}
