package controllers

import (
	"github.com/astaxie/beego/orm"
)

type DisplayMssgController struct {
	MainController
}
type Message struct {
	Id string
	Email    string
	Message  string
	Subject  string
	SendDate string
}

var Mssg []Message
var Length int

func (this *DisplayMssgController) DisplayMessage() {
	Mssg = nil
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("select id, email, message, subject, send_date from message where username=?", S_Username).QueryRows(&Mssg)
	if err != nil {
		// No result
		panic(err)
		return
	}
	Length = len(Mssg)
	this.Redirect("/specialist-messages", 302)
}

func (this *DisplayMssgController) DeleteMessage() {
	if this.Ctx.Input.Method() == "GET" {
		id := this.GetString("id")
		o := orm.NewOrm()
		o.Using("default")
		_, err := o.QueryTable("message").Filter("id", id).Delete()
		if err != nil {
			panic(err)
			return
		}
		this.Redirect("/display-mssg", 302)

	}
}
