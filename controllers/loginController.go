package controllers

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {

	//this.StartNotificationTask()
	//this.TplName = "index.tpl"
	this.TplName = "login.tpl"
}
