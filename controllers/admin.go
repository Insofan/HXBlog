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

func (self *AdminController) Post()  {
	
	fmt.Println("Logoutsucess")
	self.Redirect("/login", 302)
	var user models.User
	user.Logout()
	
}
