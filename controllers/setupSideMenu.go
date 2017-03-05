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
	//var tags = [4]string{"Go", "Swift", "ObjC", "Other"}
	controller.Data["Tag"], err = models.GetTags()
	
	//controller.Data["Tag"] = tags
	if err != nil {
		beego.Error(err)
	}
}
