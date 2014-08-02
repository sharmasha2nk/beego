package routers

import (
	"beego-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
     beego.Router("/user/?:name", &controllers.UserController{}, "get:GetUser")
     beego.SetStaticPath("/data", "data")
}
