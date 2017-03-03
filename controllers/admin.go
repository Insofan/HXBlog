package controllers

import (
	"github.com/astaxie/beego"
	"HXBlog/models"
)

type AdminController struct {
	beego.Controller
}

func (self *AdminController) Get() {
	SetupSideMenu(self.Controller)
	
	self.TplName = "admin/admin.html"
}

func (self *AdminController) LogOut()  {
	var user models.User
	user.Logout()
}
