package models

import "time"

//结构体Article
type Article struct {
	Id       int64
	Uid      int64
	Title    string
	Content  string `orm:"size(5000)"`
	Tag string
	Tags     string
	Created  time.Time `orm:"index"`
	Updated  time.Time `orm:"index"`
	Views    int64     `orm:"index"`
}