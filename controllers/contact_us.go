package controllers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/gomail"
)

type Contact_UsController struct {
	MainController
}

type EmailJson struct {
	ContactUs        string
	ContactUsFName   string
	ContactUsLName   string
	ContactUsPhone   string
	ContactUsEmail   string
	ContactUsSubject string
	ContactUsMessage string
}

var emailRes EmailJson

func (this *Contact_UsController) Get() {
	this.contact_Us()
}

func (this *Contact_UsController) Post() {

	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	f_name := dataform["f_name"].(string)
	l_name := dataform["l_name"].(string)
	phone := dataform["phone"].(string)
	email := dataform["email"].(string)
	subject := dataform["subject"].(string)
	message := dataform["message"].(string)

	if EmailValid(email) == true && PhoneValid(phone) {
		textmessage := message +" \n" + "Sender details:"+"\n"+"Name: "+ f_name+" "+l_name+"\n"+"Phone number:"+phone+"."
		m := gomail.NewMessage()
		m.SetHeader("From", email)
		m.SetHeader("To", "cetokola2015@gmail.com")
		// m.SetAddressHeader("Cc", "okolacetric@gmail.com", name)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", textmessage)
		// m.Attach("/home/Alex/lolcat.jpg")
		d := gomail.NewPlainDialer("smtp.gmail.com", 587, "cetokola2015@gmail.com", "password2015Gmail")
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		err := d.DialAndSend(m)
		if err != nil {
			fmt.Println("the error is", err)
			emailRes.ContactUs = "/"
			obj, _ := json.Marshal(emailRes)
			this.Ctx.Output.Header("Content-Type", "application/json")
			this.Ctx.Output.Body(obj)
			panic(err)
		}

		if err == nil {
			emailRes.ContactUs = "/"
			obj, _ := json.Marshal(emailRes)
			this.Ctx.Output.Header("Content-Type", "application/json")
			this.Ctx.Output.Body(obj)
		}
	}
	if EmailValid(email) == false {
		emailRes.ContactUsEmail = "invalid"
	}

	if PhoneValid(phone) == false {
		emailRes.ContactUsEmail = "invalid"
	}

	if email == "" {
		emailRes.ContactUsEmail = "empty"
	}

	if phone == "" {
		emailRes.ContactUsPhone = "empty"
	}

	if subject == "" {
		emailRes.ContactUsSubject = "empty"
	}

	if message == "" {
		emailRes.ContactUsMessage = "empty"
	}

	if f_name == "" {
		emailRes.ContactUsFName = "empty"
	}

	if l_name == "" {
		emailRes.ContactUsLName = "empty"
	}

	obj, _ := json.Marshal(emailRes)
	this.Ctx.Output.SetStatus(300)
	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body(obj)

}
