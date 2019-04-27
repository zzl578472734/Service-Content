package routers

import (
	"Service-Content/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// 用户模块
	initUserRouter()
}

func initUserRouter() {

	ns := beego.NewNamespace("/user",
		beego.NSRouter("/login", new(controllers.UserController), "post:Login"),
		beego.NSRouter("/search", new(controllers.UserController), "post:Search"),
		beego.NSRouter("/insert", new(controllers.UserController), "post:Insert"),
	)

	beego.AddNamespace(ns)
}
