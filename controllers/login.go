package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"HXBlog/models"
	"fmt"
	"os/user"
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	self.TplName = "admin/login.html"
}

//
func (self *LoginController) Login() {
	username := strings.TrimSpace(self.GetString("username"))
	password := strings.TrimSpace(self.GetString("password"))
	fmt.Printf("login controller %s %s", username, password)
	
	var user models.User
	
	if user.LoginValidation(username, password) != nil {
		self.Redirect("/login", 302)
		fmt.Println("找不到")
	} else {
		fmt.Println("loginright")
		//self.TplName = "admin/admin.html"
		self.Redirect("/admin", 302)
	}
}

