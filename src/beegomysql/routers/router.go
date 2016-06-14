package routers

import (
	"beegomysql/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user/register/:username", &controllers.MainController{}, "get:UserRegister")
}
