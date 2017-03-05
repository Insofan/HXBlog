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
	
	title := self.Input().Get("articleTitle")
	tag := "Go"
	content := self.Input().Get("content")
	fmt.Printf("mdtitle %s\n", title)
	fmt.Printf("mdtag %s\n", tag)
	models.CreateMarkdown(title, tag, content)
}