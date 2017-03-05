package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/russross/blackfriday"
	"github.com/astaxie/beego"
)

//结构体Article
type Article struct {
	Id      int64
	Title   string `orm:"size(50)"`
	Content string `orm:"size(5000)"`
	Tag     string  `orm:"index"`
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views   int64     `orm:"index"`
}

func CreateMarkdown(title, tag, content string) error {
	o := orm.NewOrm()
	var articleContent = string(blackfriday.MarkdownCommon([]byte(content)))
	//unsafe := blackfriday.MarkdownCommon(content)
	article := &Article{
		Title:   title,
		Content: articleContent,
		//Content: content,
		Tag:     tag,
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

func GetAllArticles() ([]*Article, error) {
	o := orm.NewOrm()
	articles := make([]*Article, 0)
	
	qs := o.QueryTable("Article")
	_, err := qs.All(&articles)
	
	return articles, err
}

func GetArticle() (Article) {
	o := orm.NewOrm()
	article := Article{}
	qs := o.QueryTable("Article")
	err := qs.Filter("Id", 5).One(&article)
	if err != nil {
		beego.Error(err)
	}
	
	return article
}