package models

import (
	"github.com/astaxie/beego/orm"
	"time"
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

type Symptom struct{
	Id string `orm:"pk"`
	Name string
	Disease_id string
}

type Disease struct {
	Id string `orm:"pk"`
	Name string
	Disease_id string
}

type Message struct{
	Id string `orm:"pk"`
	Username string
	Email string
	Message string
	Subject string
	SendDate time.Time
}

func init() {
	orm.RegisterModel(new(Farmer_account), new(Pro), new(Symptom), new(Disease), new(Message))
}
