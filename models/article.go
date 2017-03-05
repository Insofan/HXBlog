package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/russross/blackfriday"
	"math/rand"
	"github.com/astaxie/beego"
)

//结构体Article
type Article struct {
	Id      int64
	Uid     int64 `orm:"unique"`
	Title   string `orm:"size(50)"`
	Content string `orm:"size(5000)"`
	Tag     string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views   int64     `orm:"index"`
}

func CreateMarkdown(title, tag, content string) error {
	o := orm.NewOrm()
	var articleContent = string(blackfriday.MarkdownCommon([]byte(content)))
	article := &Article{
		Uid:   int64(rand.Intn(100000)),
		Title: title,
		Content: articleContent,
		Tag:     "Go",
		Created: time.Now(),
		Updated: time.Now(),
		Views:   0,
	}
	
	_, err := o.Insert(article)
	if err != nil {
		beego.Error(err)
		return err
	}
	
	return err
}

func GetArticles() ([]*Article, error) {
	o := orm.NewOrm()
	articles := make([]*Article, 0)
	
	qs := o.QueryTable("Article")
	_, err := qs.All(&articles)
	
	return articles, err
}