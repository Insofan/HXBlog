package controllers

import "github.com/astaxie/beego"

type CreateArticleController struct {
	beego.Controller
}

func (c *CreateArticleController) Get() {
	c.TplName = "admin/create-article.html"
}