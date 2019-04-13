package routers

import (
	"Service-Content/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// 用户模块
	InitUserRouter()
}

func InitUserRouter() {

	ns := beego.NewNamespace("/user",
		beego.NSRouter("/detail", new(controllers.UserController), "*:Detail"),
	)

	beego.AddNamespace(ns)
}
