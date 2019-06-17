package controllers

import (
	"encoding/json"
	"fmt"
	"poultry/models"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type FarmerLoginController struct {
	MainController
}

type FarmerLoginErrorJson struct {
	Messages       string
	FarmerUsername  string
	FarmerPassword string
}

var responsejson FarmerLoginErrorJson

func (this *FarmerLoginController) Get() {
	this.Authentication()
}

func (this *FarmerLoginController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	username := dataform["username"]
	password := dataform["password"]

	o := orm.NewOrm()
	o.Using("default")
	h := username.(string)
	p := password.(string)

	//check if theusername exists
	exist := o.QueryTable("farmer_account").Filter("Username", username).Exist()
	farmer := models.Farmer_account{Username: h}
	err := o.Read(&farmer, "Username")
	err = bcrypt.CompareHashAndPassword([]byte(farmer.Password), []byte(p))
	if exist == true && err == nil {
		session := this.StartSession()
		session.Set("UserID", h)
	
		
		responsejson.Messages = "/symptom_checker" 
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
		//send ajax response that the password entered is invalid
		if exist == false {
			responsejson.FarmerUsername = "invalid"
		}
		if err != nil {
			responsejson.FarmerPassword = "invalid"
		}
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

}
