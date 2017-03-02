package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	//"fmt"
	"fmt"
	"strings"
	"os/user"
	"HXBlog/models"
)


type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get()  {
	self.TplName = "admin/login.html"
}
//
func (self *LoginController) Login() {
	//username := self.GetString("username")
	username := strings.TrimSpace(self.GetString("username"))
	password := strings.TrimSpace(self.GetString("password"))
	remember := self.GetString("")

	if username != "" && password != "" {
		var user models.User
		user.Username = username
		if user.Read("user_name") != nil || user.Read("user_password") != nil{
			self.Data["errmsg"] = "帐号或密码错误"
		}
	}
}