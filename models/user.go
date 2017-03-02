package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type User struct {
	Id       int64
	Username string `orm:"unique;size(15)"`
	Password string `orm:"size(32)"`
}

func (m *User) TableName() string {
	return TableName("user")
}

func (m *User) CheckAccount(username, password string) error {
	o := orm.NewOrm()
	user := User{}
	qs := o.QueryTable("user")
	err := qs.Filter("username", username).Filter("password", password).One(&user)
	
	if err != nil {
		fmt.Println("查询不到")
		return err
	} else {
		return nil
	}
}
