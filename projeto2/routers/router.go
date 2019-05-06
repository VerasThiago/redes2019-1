package routers

import (
	"projeto2/controllers"
	"projeto2/models"

	"github.com/astaxie/beego"
)

func init() {

	socket := models.Socket{
		Ip:   "codeforces.com",
		Port: 80,
	}
	socket.Init()

	beego.Router("/", &controllers.MainController{})
	beego.Router("/handle", &controllers.UserController{S: &socket})
}
