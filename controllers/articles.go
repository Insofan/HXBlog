package controllers

import (
	"github.com/astaxie/beego"
	"HXBlog/models"
	"fmt"
)

type ArticlesController struct {
	beego.Controller
}

func (self *ArticlesController) Get() {
	var err error
	//取得所有文章
	self.Data["Articles"], err = models.GetArticles()
	if err != nil {
		beego.Error(err)
	}
	SetupSideMenu(self.Controller)
	self.TplName = "admin/articles.html"
}

func (self *ArticlesController) Post() {
	
	fmt.Println("Logoutsucess")
	self.Redirect("/login", 302)
	var user models.User
	user.Logout()
	
}