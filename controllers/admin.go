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
	
	var err error
	self.Data["Categories"], err = models.GetCategories()
	if err != nil {
		beego.Error(err)
	}
	fmt.Printf("adminget %s",self.Data["Categories"])
	
	self.TplName = "admin/admin.html"
}