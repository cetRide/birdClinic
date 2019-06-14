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

func (this *MainController) Get() {
	this.Data["Title"] = "iReferral-Welcome"
	this.Layout = "layout.tpl"
		this.TplName = "home.html"
}

func (this *MainController) home(view string) {
	this.Data["Title"] = "iReferral-Welcome"
	this.Layout = "layout.tpl"
	this.TplName = view + ".html"
}
