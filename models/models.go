package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
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
	dbuser string = "root"
	dbpassword string = "admin"
	dbname string = "hxblog"
)

//结构体Article
type Article struct {
	Id       int64
	Uid      int64
	Title    string
	Content  string `orm:"size(5000)"`
	Category string
	Tags     string
	Created  time.Time `orm:"index"`
	Updated  time.Time `orm:"index"`
	Views    int64     `orm:"index"`
}

//结构体Category
type Category struct {
	Id           int64
	Title        string
	Created      time.Time `orm:"index"`
	Updated      time.Time `orm:"index"`
	ArticleCount int64
}

//注册DB
func RegisterDB()  {
	conn := dbuser + ":" + dbpassword + "@/" + dbname + "?charset=utf8&loc=Local"
	orm.RegisterModel(new(Article), new(Category))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", conn, 10, 10)
	
}
