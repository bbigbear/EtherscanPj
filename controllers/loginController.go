package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {

	//this.StartNotificationTask()
	//this.TplName = "index.tpl"
	this.TplName = "login.tpl"
}

func (this *LoginController) LoginAction() {

	fmt.Println("点击登录按钮")
	uname := this.Input().Get("inputAccount")
	pwd := this.Input().Get("inputPassword")

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		fmt.Println("登录成功")
		//存session
		this.SetSession("islogin", 1)
		this.ajaxMsg("登录成功", MSG_OK)
	} else {
		fmt.Println("账户密码错误")
		this.ajaxMsg("账户密码错误", MSG_ERR)
	}
}

func (this *LoginController) Logout() {

	fmt.Println("点击推出按钮")
	this.DelSession("islogin")
	this.TplName = "login.tpl"

}
