package controllers

import "github.com/astaxie/beego"

type CreateArticleNormalController struct {
	beego.Controller
}

func (self *CreateArticleNormalController) Get() {
	SetupSideMenu(self.Controller)
	self.TplName = "admin/create-article-normal.html"
}
