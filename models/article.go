package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//结构体Article
type Article struct {
	Id      int64
	Uid     int64
	Title   string
	Content string `orm:"size(5000)"`
	Tag     string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views   int64     `orm:"index"`
}

func GetArticles() ([]*Article, error) {
	o := orm.NewOrm()
	articles := make([]*Article, 0)
	
	qs := o.QueryTable("Article")
	_, err := qs.All(&articles)
	
	return articles, err
}