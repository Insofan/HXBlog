package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//结构体Category
type Category struct {
	Id      int64
	Title   string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Total   int64
}

func GetCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	
	qs := o.QueryTable("Category")
	_, err := qs.All(&cates)
	
	return cates, err
}