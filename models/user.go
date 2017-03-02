package models

import "github.com/astaxie/beego/orm"

//import (
//	"github.com/astaxie/beego/orm"
//)

type User struct {
	Id int64
	Username string `orm:"unique;size(15)"`
	Password string `orm:"size(32)"`
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
