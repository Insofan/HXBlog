package main

import (
	_ "HXBlog/routers"
	"github.com/astaxie/beego"
	"HXBlog/models"
	"github.com/astaxie/beego/orm"
)
func init(){
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	
	orm.RunSyncdb("default", false, true)
	
	//因为使用正则路由器，所以下面要设置静态文件地址，请注意IDE或者beego有BUG 将.CSS文件重命名之后才正确读取
	beego.SetStaticPath("/admin/articles/view-post", "views/admin")
	beego.Run(":9000")
}

