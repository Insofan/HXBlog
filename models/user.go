package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/context"
	"fmt"
)

type User struct {
	Id       int64
	Username string `orm:"unique;size(15)"`
	Password string `orm:"size(32)"`
	IsLogin  int8 `orm:"bit default(0)"`
}

var currentUsername string

func (m *User) TableName() string {
	return TableName("user")
}

func (m *User) LoginValidation(username, password string) error {
	o := orm.NewOrm()
	user := User{}
	qs := o.QueryTable("user")
	err := qs.Filter("username", username).Filter("password", password).One(&user)
	
	if err != nil {
		fmt.Println("查询不到用户")
		return err
	} else {
		currentUsername = username
		_, err := o.QueryTable("user").Filter("username", username).Update(orm.Params{
			"IsLogin": 1,
		})
		if err != nil {
			fmt.Println("查询不到用户")
			return err
		}
		
		return nil
	}
}

func (m *User) Logout() {
	o := orm.NewOrm()
	o.QueryTable("user").Filter("username", currentUsername).Update(orm.Params{
		"IsLogin": 0,
	})
}

func CheckAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	//fmt.Printf("username and password : %s %s", username, password)
	if err != nil {
		fmt.Println("get error")
		return false
	}
	
	username := ck.Value
	
	return username == currentUsername
	
}
