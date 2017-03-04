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
	
	beego.Run(":9000")
}

