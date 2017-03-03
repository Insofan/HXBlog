package controllers

import (
	"github.com/astaxie/beego"
	"HXBlog/models"
	"fmt"
)

type CreateArticleMarkdownController struct {
	beego.Controller
}

func (self *CreateArticleMarkdownController) Get() {
	SetupSideMenu(self.Controller)
	self.TplName = "admin/create-article-markdown.html"
}

func (self *CreateArticleMarkdownController) Post()  {
	
	fmt.Println("Logoutsucess")
	self.Redirect("/login", 302)
	var user models.User
	user.Logout()
	
}