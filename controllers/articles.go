package controllers

import "github.com/astaxie/beego"

type ArticlesController struct {
	beego.Controller
}

func (c *ArticlesController) Get() {
	c.TplName = "admin/articles.html"
}
