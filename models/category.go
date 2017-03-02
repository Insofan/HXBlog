package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

//结构体Category
type Category struct {
	Id           int64
	Title        string
	Created      time.Time `orm:"index"`
	Updated      time.Time `orm:"index"`
	ArticleCount int64
}

func GetCategories() ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	
	qs := o.QueryTable("Category")
	_, err := qs.All(&categories)
	fmt.Printf("cates %s",qs)
	
	return categories, err
}