package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//结构体Category
type Tag struct {
	Id      int64
	Go      string
	Swift   string
	ObjC    string
	Other   string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views   int64     `orm:"index"`
}

func GetTags() ([]*Tag, error) {
	o := orm.NewOrm()
	cates := make([]*Tag, 0)
	
	qs := o.QueryTable("Tag")
	_, err := qs.All(&cates)
	
	return cates, err
}