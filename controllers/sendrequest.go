package controllers

import (
	"fmt"
	"poultry/models"

	"github.com/astaxie/beego/orm"
	// "log"
	// "net/url"
)

type SendMessageController struct {
	MainController
}

var Email, Username, FirstName1, LastName1, District1, County1, Country1, PhoneNo string

func (this *SendMessageController) SendMessage() {

	// this.sendMessage("sendmessge")
	if this.Ctx.Input.Method() == "GET" {
		email := this.GetString("email")
		fmt.Println("Echoe the email", email)
		o := orm.NewOrm()
		o.Using("default")

		pro := models.Pro{Email: email}
		err := o.Read(&pro, "Email")

		if err == orm.ErrNoRows {
			fmt.Println(err)
			panic(err)
			return
		}
		if err == nil {
			FirstName1 = pro.First_name
			LastName1 = pro.Last_name
			District1 = pro.District
			Country1 = pro.Country
			County1 = pro.County
			PhoneNo = pro.Phone
			Email = pro.Email
			Username = pro.Username
			this.Redirect("/send-message", 302)
		}
	}

}
