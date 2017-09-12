package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       int
	Username string
	Password string
}
func init() {
	orm.RegisterModel(new(Users))
}

func FindUser(username string) (Users, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Username: username}
	err := o.Read(&user, "username")

	return user, err
}