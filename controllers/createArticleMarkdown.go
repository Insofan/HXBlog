package controllers

import (
	"github.com/astaxie/beego"
)

type CreateArticleMarkdownController struct {
	beego.Controller
}

func (self *CreateArticleMarkdownController) Get() {
	SetupSideMenu(self.Controller)
	self.TplName = "admin/create-article-markdown.html"
}
