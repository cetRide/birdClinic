package controllers

import (
	"github.com/astaxie/beego"

)

type MainController struct {
	beego.Controller
}

var flash = beego.NewFlash()
func (this *MainController) Prepare() {

	this.Data["HeadStyles"] = []string{
		"/static/css/mdb/bootstrap.min.css",
		"/static/css/mdb/mdb.min.css",
		"/static/css/mdb/style.min.css",
		"/static/css/jquery.transfer.css",
		"/static/css/a.css",
		"/static/css/icon_font/css/icon_font.css",
	
		}

	this.Data["HeadScripts"] = []string{
		"/static/js/mdb-js/jquery-3.3.1.min.js",
		"/static/js/mdb-js/mdb.min.js",
		"/static/js/mdb-js/bootstrap.min.js",
		"/static/js/mdb-js/popper.min.js",
		"/static/js/axios.min.js",
		
		
	}
}

//landing page render
func (this *MainController) Get() {
	this.Data["Title"] = "Bird clinic symptom-home"
	this.Layout = "layout.tpl"
	this.TplName = "landing.html"
}

//authentication page render
func (this *MainController) Authentication() {
	this.Data["Title"] = "Get started with bird clinic"
	this.Layout = "layout.tpl"
	this.TplName = "authentication.html"
}

//contact us page render
func (this *MainController) contact_Us() {
	this.Data["Title"] = "Bird clinic- Write to us"
	this.Layout = "layout.tpl"
	this.TplName = "contactus.html"
}

//symptom checker page render
func (this *MainController) Symptom_Checker() {
	this.Data["Title"] = "Bird clinic- Symptom_checker"
	this.Layout = "layout.tpl"
	this.TplName = "symptoms.html"
}

func (this *MainController) SpecialistAuth() {
	this.Data["Title"] = "Bird clinic- get started as a specialist"
	this.Layout = "layout.tpl"
	this.TplName = "specialistauth.html"
}

func (this *MainController) SpecialistMess() {
	this.Data["Title"] = "Bird clinic- view messages"
	this.Layout = "layout.tpl"
	this.TplName = "specialistmessage.html"
}

func (this *MainController) DiseasesResponse() {
	this.Data["Title"] = "Bird clinic- view potential diseases"
	this.Data["Name"] = Diseases
	this.Layout = "layout.tpl"
	this.TplName = "diseaseoutput.html"
}

func (this *MainController) find_pro(view string) {
	this.Data["Title"] = "Bird clinic- search for a specialist"
	// this.Data["Name"] = Diseases
	this.Layout = "layout.tpl"
	this.TplName = view+".html"
}

func (this *MainController) NoResults() {
	this.Data["Title"] = "Bird clinic- the results found"
	this.Data["Search"] = Search
	this.Data["Len"] = Len
	this.Data["Results"] = Pros
	this.Layout = "layout.tpl"
	this.TplName = "noresults.html"
}

func (this *MainController) sendMessage(view string) {
	this.Data["Title"] = "Bird clinic- contact the specialist"
	this.Data["District"] = District1
	this.Data["County"] = County1
	this.Data["Country"] = Country1
	this.Data["FirstName"] = FirstName1
	this.Data["LastName"] = LastName1
	this.Layout = "layout.tpl"
	this.TplName = view +".html"
}

func (this *MainController) Logout() {
	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID != nil {
		// UserID is set and can be deleted
		session.Delete("UserID")
	}
	this.Redirect("/", 302)
}


