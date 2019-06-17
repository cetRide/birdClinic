package routers

import (
	"poultry/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/authentication", &controllers.MainController{}, "get,post:Authentication")
	beego.Router("/specialist-authentication", &controllers.MainController{}, "get,post:SpecialistAuth")
	beego.Router("/specialist-messages", &controllers.MainController{}, "get,post:SpecialistMess")
	beego.Router("/contact_us", &controllers.MainController{}, "get,post:Contact_Us")
	beego.Router("/symptom_checker", &controllers.MainController{}, "get,post:Symptom_Checker")
	beego.Router("/f-login", &controllers.FarmerLoginController{})
	beego.Router("/f-signup", &controllers.FarmerSignUpController{})
	beego.Router("/s-login", &controllers.SpecialistLoginController{})
	beego.Router("/s-signup", &controllers.ProSignUpController{})
	beego.Router("/symptom_submit", &controllers.TestController{})
	beego.Router("/logout", &controllers.MainController{}, "get,post:Logout")
}

