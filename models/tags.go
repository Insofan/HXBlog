package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//结构体Category
type Tag struct {
	Id      int64
	Title   string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Total   int64
}

func GetTags() ([]*Tag, error) {
	o := orm.NewOrm()
	cates := make([]*Tag, 0)
	
	qs := o.QueryTable("Tag")
	_, err := qs.All(&cates)
	
	return cates, err
}