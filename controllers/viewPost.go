package controllers

import (
	"github.com/astaxie/beego"
	"HXBlog/models"
)
type ViewPostController struct {
	beego.Controller
}

func (self *ViewPostController) Get() {
	self.TplName = "admin/view-post.html"
	articleId := self.Ctx.Input.Param(":id")
	self.Data["ViewPost"] = models.GetArticle(articleId)
	
}