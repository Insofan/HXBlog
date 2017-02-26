package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "admin/login.html"
}