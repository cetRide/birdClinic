package controllers

import (
	"encoding/json"
	"fmt"
	"poultry/models"

	"github.com/astaxie/beego/orm"
	"github.com/twinj/uuid"
)

type ProSignUpController struct {
	MainController
}

type ProErrorJson struct {
	ResS        string
	SId         string
	SUsername   string
	SFirst_name string
	SLast_name  string
	SEmail      string
	SPhone      string
	SDistrict   string
	SCounty     string
	SCountry    string
	SPassword   string
	ScPassword  string
}

var sjson ProErrorJson

func (this *ProSignUpController) Get() {
	this.SpecialistAuth()
}

func (this *ProSignUpController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	username := dataform["username"].(string)
	first := dataform["first"].(string)
	last := dataform["last"].(string)
	email := dataform["email"].(string)
	phone := dataform["phone"].(string) //
	district := dataform["district"].(string)
	county := dataform["county"].(string)
	country := dataform["country"].(string)
	password := dataform["pass"].(string)         //
	copassword := dataform["copassword"].(string) //

	o := orm.NewOrm()
	o.Using("default")

	exist := o.QueryTable("pro").Filter("Username", username).Exist()
	exist1 := o.QueryTable("pro").Filter("Email", email).Exist()
	exist2 := o.QueryTable("pro").Filter("Phone", phone).Exist()
	if exist == false && exist1 == false && exist2 == false && password == copassword && IsValid(password) == true && PhoneValid(phone) == true && EmailValid(email) == true {
		password, _ := HashPassword(password) //hash the submitted password
		//insert into the database the details
		pro := models.Pro{Username: username, First_name: first, Last_name: last, Email: email, Phone: phone, District: district, County: county, Country: country, Password: password}
		registration_uid := uuid.NewV4().String()
		pro.Id = registration_uid
		_, err := o.Insert(&pro)
		if err != nil {
			fmt.Println("Internal server error")
		}
		sjson.ResS = "/specialist-messages" //redirect to log in page
		obj, _ := json.Marshal(sjson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
		//send ajax response that the password entered is invalid
		if exist == true {
			sjson.SUsername = "exist"
		}
		if exist1 == true {
			sjson.SEmail = "exist"
		}
		if exist2 == true {
			sjson.SPhone = "exist"
		}

		if first == "" {
			sjson.SFirst_name = "empty"
		}

		if last == "" {
			sjson.SLast_name = "empty"
		}
		if email == "" {
			sjson.SEmail = "empty"
		}

		if username == "" {
			sjson.SUsername = "empty"
		}

		if district == "" {
			sjson.SDistrict = "empty"
		}

		if county == "" {
			sjson.SCounty = "empty"
		}
		if country == "" {
			sjson.SCountry = "empty"
		}

		if password != copassword {
			sjson.SPassword = "nomatch"
			sjson.ScPassword = "nomatch"
		}

		if IsValid(password) == false {
			sjson.SPassword = "incorrect"
		}

		if PhoneValid(phone) == false {
			sjson.SPhone = "wrong"
		}

		obj, _ := json.Marshal(sjson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

}
