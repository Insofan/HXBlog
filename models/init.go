package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
//
//#MYSQL 地址
//dbhost = localhost
//
//#MYSQL 端口
//dbport = 3306
//
//#MYSQL 密码
//dbpassword =
var (
	dbuser     string = "root"
	dbpassword string = "admin"
	dbname     string = "db_hxblog"
)

//注册DB
func RegisterDB() {
	
	conn := dbuser + ":" + dbpassword + "@/" + dbname + "?charset=utf8&loc=Local"
	orm.RegisterModel( new(User), new(Category))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", conn, 10, 10)
}

func TableName(str string) string {
	return str
}
