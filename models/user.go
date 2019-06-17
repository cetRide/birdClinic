package models

import (
	"github.com/astaxie/beego/orm"
)

type Farmer_account struct {
	Id       string `orm:"pk"`
	Username string
	Email    string `orm:"unique"`
	Password string `valid:"Required"`
}

type Pro struct{
	Id string `orm:"pk"`
	Username string `orm:"unique"`
	First_name string
	Last_name string
	Email string `orm:"unique"`
	Phone string `orm:"unique"`
	District string
	County string 
	Country string
	Password string
}


func init() {
	orm.RegisterModel(new(Farmer_account), new(Pro))
}
