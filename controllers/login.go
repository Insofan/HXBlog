package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	//"fmt"
)


type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get()  {
	//if self.Input().Get("exit") == "true" {
	//	self.Ctx.SetCookie("username", "", -1, "/")
	//	self.Ctx.SetCookie("password","",-1, "/")
	//	self.Redirect("/", 302)
	//	return
	//}
	c.TplName = "admin/login.html"
}
//
func (c *LoginController) Post() {
	//username := c.Input().Get("username")
	//password := c.Input().Get("password")
	//autoLogin := c.Input().Get("autoLogin") == "on"
}