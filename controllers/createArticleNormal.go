package controllers

import "github.com/astaxie/beego"

type CreateArticleNormalController struct {
	beego.Controller
}

func (c *CreateArticleNormalController) Get() {
	c.TplName = "admin/create-article-normal.html"
}
