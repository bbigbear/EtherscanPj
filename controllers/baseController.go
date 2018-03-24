package controllers

import (
	"EtherscanPj/models"
	"fmt"
	"log"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
)

const (
	MSG_OK            = 200
	MSG_ERR_Param     = 400
	MSG_ERR_Verified  = 401
	MSG_ERR_Authority = 403
	MSG_ERR_Resources = 404
	MSG_ERR           = 500
)

type BaseController struct {
	beego.Controller
}

//ajax返回
func (this *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

//ajax返回 列表
func (this *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["message"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

// 通过两重循环过滤重复元素
func (this *BaseController) RemoveRepBySlice(slc []string) []string {
	result := []string{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 图片接口
func (this *BaseController) PutFileImg() {
	h, err := this.GetFiles("file")
	fmt.Println("文件名称", h[0].Filename)
	fmt.Println("文件大小", h[0].Size)
	if err != nil {
		log.Fatal("getfile err ", err)
		this.ajaxMsg(h[0].Filename+"图片上传失败", MSG_ERR_Resources)
	}
	//	defer f.Close()
	path := "static/upload/" + h[0].Filename
	this.SaveToFile("file", path) // 保存位置在 static/upload, 没有文件夹要先创建
	list := make(map[string]interface{})
	list["src"] = path
	list["name"] = h[0].Filename
	list["size"] = h[0].Size
	this.ajaxList("图片上传成功", MSG_OK, 1, list)
}

//将时间化为秒
func (this *BaseController) GetSecs(ordertime string) int64 {
	var s int64
	t, err := time.ParseInLocation("2006-01-02 15:04:05", ordertime, time.Local)
	if err == nil {
		s = t.Unix()
		return s
	} else {
		return -1
	}
}

//获取相差时间
func (this *BaseController) GetMinuteDiffer(server_time, mqtime string) int64 {
	var minute int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", server_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", mqtime, time.Local)
	if err == nil {
		diff := t1.Unix() - t2.Unix()
		minute = diff / 60
		return minute
	} else {
		return -1
	}
}

//数据扫描
func ScanData() {
	fmt.Println("开始扫描")
	o := orm.NewOrm()
	o.Using("db")
	var maps_update []orm.Params

	//获取更新数据
	//count, err := o.Raw("select * from Token order by id desc limit ?", n).Values(&maps_update)
	count, err := o.Raw("select * from Token").Values(&maps_update)
	if err != nil {
		fmt.Println("err!")
	}
	fmt.Println(count)
	//循环数据
	a := 0
	for _, u := range maps_update {
		//获取数据库中数据，进行过滤
		a++
		constact_address := u["contractAddress"].(string)
		to_address := u["toAddress"].(string)
		from_address := u["fromAddress"].(string)
		value, err := decimal.NewFromString(u["value"].(string))
		if err != nil {
			fmt.Println("err!")
		}
		//获取list1 的数据
		o1 := orm.NewOrm()
		monitor := new(models.Monitior)
		status := new(models.Status)
		var maps_monitor []orm.Params

		num1, err := o1.QueryTable(monitor).Values(&maps_monitor)
		if err != nil {
			fmt.Println("获取list失败")
		}
		fmt.Println("get list ok! num", num1)
		for _, m := range maps_monitor {
			address := m["Address"].(string)
			fmt.Println("address", address)
			if constact_address == "0x95408930d6323ac7aa69e6c2cbfe58774d565fa8" || constact_address == "0xa9ec9f5c1547bd5b0247cf6ae3aab666d10948be" {
				fmt.Println("contracct", m["Contract"].(string), constact_address)
				fmt.Println("address:", to_address, from_address, address)
				if address == to_address || address == from_address {
					//判断是否开启通知状态
					//fmt.Println("address:", to_address, from_address, address)
					var status_data models.Status
					o1.QueryTable(status).Filter("Id", 1).One(&status_data)
					if err != nil {
						fmt.Println("查找status失败")
					}
					//将数据存入data中
					var token_data models.Data
					token_data.BlockNumber = u["blockNumber"].(string)
					token_data.ContractAddress = constact_address
					token_data.FromAddress = from_address
					t, _ := time.Parse("2006-01-02 15:04:05", u["timestamp"].(string))
					token_data.Timestamp = t
					token_data.ToAddress = to_address
					token_data.TransactionHash = u["transactionHash"].(string)
					token_data.Value = u["value"].(string)
					if status_data.Status == "on" {
						token_data.Status = "警告"
						//获取交易value
						value1, err := decimal.NewFromString(status_data.Svalue)
						if err != nil {
							fmt.Println("err!")
						}
						//fmt.Println("num:", num1, maps, value.LessThan(value1))
						if value.LessThan(value1) {
							fmt.Println("不进行推送")
						} else {
							fmt.Println("进行推送")
							var notify models.Notifcation
							notify.Hash = u["transactionHash"].(string)
							notify.Num = u["value"].(string)
							notify.Style = "单笔交易"
							t1, _ := time.Parse("2006-01-02 15:04:05", u["timestamp"].(string))
							notify.Timestamp = t1
							notify.Target = m["Name"].(string)
							num, err := o1.Insert(&notify)
							if err != nil {
								fmt.Println("isnert err!")
							}
							fmt.Println("insert id", num)
						}
					}
					num2, err := o1.Insert(&token_data)
					if err != nil {
						fmt.Println("插入失败")
					}
					fmt.Println("插入data成功num", num2)

				}
			}
		}
	}
	fmt.Println("sum scan data complete!sum:", a)
}
