package controllers

import (
	"github.com/astaxie/beego"
	"HXBlog/models"
)

func SetupSideMenu(controller beego.Controller) {
	var user models.User
	if user.HasLogin() == 0 {
		controller.Redirect("/login", 302)
	}
	var err error
	controller.Data["Tag"], err = models.GetTags()
	if err != nil {
		beego.Error(err)
	}
}
