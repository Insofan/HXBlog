package controllers

import (
	"github.com/astaxie/beego"
	"HXBlog/models"
)

func SetupSideMenu(controller beego.Controller) {
	//if models.CheckAccount(controller.Ctx) == false{
	//
	//	controller.Redirect("/login", 302)
	//}
	var err error
	controller.Data["Tag"], err = models.GetTags()
	if err != nil {
		beego.Error(err)
	}
}
