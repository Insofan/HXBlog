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

func (self *CreateArticleMarkdownController) Post() {
	fmt.Println("mdcreate")
	
	//创建markdown文章
	title := self.Input().Get("articleTitle")
	tag := self.Input().Get("tags")
	content := self.Input().Get("content")
	models.CreateMarkdown(title, tag, content)
}