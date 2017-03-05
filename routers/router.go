package routers

import (
	"HXBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
	
	//登录路由
	
	//beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{},"get:Get")
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	
	//后台路由
	beego.Router("/admin", &controllers.AdminController{}, "get:Get")
	beego.Router("/admin", &controllers.AdminController{}, "post:Post")
	beego.Router("/admin/create-article-markdown", &controllers.CreateArticleMarkdownController{}, "get:Get")
	beego.Router("/admin/create-article-markdown", &controllers.CreateArticleMarkdownController{}, "post:Post")
	beego.Router("/admin/create-article-normal", &controllers.CreateArticleNormalController{})
	beego.Router("/admin/articles", &controllers.ArticlesController{})
	
}
