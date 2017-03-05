package controllers

import (
	"github.com/astaxie/beego"
	//"HXBlog/models"
	"fmt"
)

func SetupSideMenu(controller beego.Controller) {
	//if models.CheckAccount(controller.Ctx) == false{
	//
	//	controller.Redirect("/login", 302)
	//}
	var tags = [4]string{"Go", "Swift", "ObjC", "Other"}
	controller.Data["Tag"] = tags
	fmt.Println("sidemenu%s", tags)
	
	//controller.Data["Tag"] = tags

}
