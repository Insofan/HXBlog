package controllers

import (
	"github.com/astaxie/beego"
	"HXBlog/models"
)

func SetupSideMenu(controller beego.Controller) {
	var err error
	controller.Data["Category"], err = models.GetCategories()
	if err != nil {
		beego.Error(err)
	}
}
