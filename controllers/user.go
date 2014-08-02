package controllers

import (
	"github.com/astaxie/beego"
) 

type UserController struct {
	beego.Controller
}

func (this *UserController) GetUser() {
	this.Data["Name"] = this.Ctx.Input.Param(":name")
	this.Data["Email"] = "sharma.sha2nk@gmail.com"
	this.TplNames = "user.tpl"
}