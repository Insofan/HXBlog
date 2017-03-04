package controllers

import (
	"github.com/astaxie/beego"
	"HXBlog/models"
	"fmt"
)

type AdminController struct {
	beego.Controller
}

func (self *AdminController) Get() {
	SetupSideMenu(self.Controller)
	
	self.TplName = "admin/admin.html"
}

func (self *AdminController) Post() {
	
	fmt.Println("Logoutsucess")
	
	var user models.User
	user.Logout()
	self.Redirect("/login", 302)
	return
	
}
