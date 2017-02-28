package controllers

import "github.com/astaxie/beego"

type CreateArticleMarkdownController struct {
	beego.Controller
}

func (c *CreateArticleMarkdownController) Get() {
	c.TplName = "admin/create-article-markdown.html"
}