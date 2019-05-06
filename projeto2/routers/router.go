package routers

import (
	"projeto2/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/handle", &controllers.UserController{})
}
