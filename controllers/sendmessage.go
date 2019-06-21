package controllers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"poultry/models"
"time"
	"github.com/astaxie/beego/orm"
	"github.com/gomail"
	"github.com/twinj/uuid"
)

type WriteToProController struct {
	MainController
}

type MessageJson struct {
	Messge  string
	Heading string
	Body    string
}

var res MessageJson

func (this *WriteToProController) Get() {
	this.sendMessage("sendmessge")
}

func (this *WriteToProController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	subject := dataform["subject"].(string)
	message := dataform["message"].(string)

	o := orm.NewOrm()
	o.Using("default")
	mssg := models.Message{Username: Username, Email: F_email, Message: message, Subject: subject}
	registration_uid := uuid.NewV4().String()
	mssg.Id = registration_uid
	dt := time.Now()
	dt.Format("2006-01-02")
	mssg.SendDate = dt 
	_, err := o.Insert(&mssg)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Internal server error")
		panic(err)
		return
	}
	m := gomail.NewMessage()
	m.SetHeader("From", F_email)
	m.SetHeader("To", Email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "cetokola2015@gmail.com", "password2015Gmail")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = d.DialAndSend(m)
	if err != nil {
		fmt.Println("the error is", err)
	res.Messge = "/symptom_checker"
		obj, _ := json.Marshal(res)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
		panic(err)
	}

	if err == nil {
		res.Messge = "/symptom_checker"
		obj, _ := json.Marshal(res)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

	if subject == "" {
		res.Heading = "empty"
	}

	if message == "" {
		res.Body = "empty"
	}

	obj, _ := json.Marshal(res)
	this.Ctx.Output.SetStatus(300)
	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body(obj)
}
