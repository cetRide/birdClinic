package routers

import (
	"poultry/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/symptom_submit", &controllers.TestController{})
}

