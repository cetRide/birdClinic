package controllers

import (
	"encoding/json"
	"fmt"
	"poultry/models"
	"github.com/twinj/uuid"
	"github.com/astaxie/beego/orm"
	
)

type FarmerSignUpController struct {
	MainController
}

type MgnSErrorJson struct {
	Feedback  string
	Username  string
	Email     string
	Fpassword string
	Cpassword string
	
}

var mgnresponsejson MgnSErrorJson

func (this *FarmerSignUpController) Get() {
	this.Authentication()
}

func (this *FarmerSignUpController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	username := dataform["username"].(string)
	email := dataform["email"].(string)
	password := dataform["password"].(string)
	cpassword := dataform["cpassword"].(string)

	o := orm.NewOrm()
	o.Using("default")

	exist := o.QueryTable("farmer_account").Filter("Username", username).Exist()
	exist1 := o.QueryTable("farmer_account").Filter("Email", email).Exist()
	fmt.Println(exist)
		if exist == false && exist1 == false && password == cpassword && IsValid(password) == true && EmailValid(email) == true{
		password, _ := HashPassword(password) //hash the submitted password
		//insert into the database the details
		farmer := models.Farmer_account{Username: username, Email: email, Password: password}
		registration_uid := uuid.NewV4().String()

		farmer.Id = registration_uid
	  
		_, err := o.Insert(&farmer)
		fmt.Println(err)
		if err != nil {
			fmt.Println("Internal server error")
		}
		mgnresponsejson.Feedback = "/authentication" //redirect to log in page
		obj, _ := json.Marshal(mgnresponsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
		//send ajax response that the password entered is invalid
		if exist == true {
			mgnresponsejson.Username = "exist"
		}
		if exist1 == true {
			mgnresponsejson.Email = "exist"
		}
		if password != cpassword {
			mgnresponsejson.Fpassword = "notmatch"
			mgnresponsejson.Cpassword = "notmatch"
		}

		if IsValid(password) == false {
			mgnresponsejson.Fpassword = "incorrect"
		}
		if EmailValid(email) == false {
			mgnresponsejson.Email = "invalid"
		}

		obj, _ := json.Marshal(mgnresponsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

}
