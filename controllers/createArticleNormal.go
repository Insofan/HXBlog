package controllers

import (
	"github.com/astaxie/beego"
	"HXBlog/models"
	"fmt"
)

type CreateArticleNormalController struct {
	beego.Controller
}

func (self *CreateArticleNormalController) Get() {
	SetupSideMenu(self.Controller)
	self.TplName = "admin/create-article-normal.html"
}

func (self *CreateArticleNormalController) Post() {
	
	fmt.Println("Logoutsucess")
	self.Redirect("/login", 302)
	var user models.User
	user.Logout()
	
}