package controllers

import (
	"encoding/json"
	"fmt"
	"poultry/models"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type SpecialistLoginController struct {
	MainController
}

type SpecialistLoginErrorJson struct {
	SpecialistResponse string
	SpecialistUsername string
	SpecialistPassword string
}

var S_Username string
var spresponsejson SpecialistLoginErrorJson

func (this *SpecialistLoginController) Get() {
	this.SpecialistAuth()
}

func (this *SpecialistLoginController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	username := dataform["user"]
	password := dataform["pass"]

	o := orm.NewOrm()
	o.Using("default")
	h := username.(string)
	p := password.(string)

	//check if theusername exists
	exist := o.QueryTable("pro").Filter("Username", username).Exist()
	pro := models.Pro{Username: h}
	err := o.Read(&pro, "Username")
	err = bcrypt.CompareHashAndPassword([]byte(pro.Password), []byte(p))
	if exist == true && err == nil {
		session := this.StartSession()
		session.Set("UserID", h)
		S_Username = pro.Username
		spresponsejson.SpecialistResponse = "/display-mssg"
		obj, _ := json.Marshal(spresponsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
			//send ajax response that the password entered is invalid
			if exist == false {
				spresponsejson.SpecialistUsername = "incorrect"
			}
			if err != nil {
				spresponsejson.SpecialistPassword = "incorrect"
			}
			obj, _ := json.Marshal(spresponsejson)
			this.Ctx.Output.SetStatus(300)
			this.Ctx.Output.Header("Content-Type", "application/json")
			this.Ctx.Output.Body(obj)
		}

	}

