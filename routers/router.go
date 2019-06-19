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
	beego.Router("/search-results", &controllers.MainController{}, "get,post:NoResults")
	beego.Router("/contact_us", &controllers.Contact_UsController{})
	beego.Router("/symptom_checker", &controllers.MainController{}, "get,post:Symptom_Checker")
	beego.Router("/find_pro", &controllers.FindProController{}, "get,post:FindPro")
	beego.Router("/f-login", &controllers.FarmerLoginController{})
	beego.Router("/f-signup", &controllers.FarmerSignUpController{})
	beego.Router("/s-login", &controllers.SpecialistLoginController{})
	beego.Router("/s-signup", &controllers.ProSignUpController{})
	beego.Router("/symptom_submit", &controllers.SymptomController{})
	beego.Router("/potential_risks", &controllers.MainController{}, "get,post:DiseasesResponse")
	beego.Router("/logout", &controllers.MainController{}, "get,post:Logout")
}

