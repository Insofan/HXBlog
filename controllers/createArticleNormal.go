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
	
	//创建markdown文章
	title := self.Input().Get("articleTitle")
	tag := self.Input().Get("tags")
	content := self.Input().Get("content")
	
	fmt.Printf("createnormalpost%s",content)
	models.CreateNormal(title, tag, content)
	
}