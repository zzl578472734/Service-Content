package routers

import (
	"Service-Content/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// 管理员模块
	initAdminRouter()

	// 用户模块
	initUserRouter()

	// 角色模块
	initRoleRouter()
}

func initAdminRouter()  {

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/login", new(controllers.AdminController), "post:Login"),
		beego.NSRouter("/insert", new(controllers.AdminController), "post:Insert"),
	)
	beego.AddNamespace(ns)
}

func initUserRouter() {

	ns := beego.NewNamespace("/user",
		beego.NSRouter("/login", new(controllers.UserController), "post:Login"),
		beego.NSRouter("/search", new(controllers.UserController), "post:Search"),
		beego.NSRouter("/insert", new(controllers.UserController), "post:Insert"),
	)

	beego.AddNamespace(ns)
}

func initRoleRouter() {

	ns := beego.NewNamespace("/role",
		beego.NSRouter("/list", new(controllers.RoleController), "post:List"),
		beego.NSRouter("/insert", new(controllers.RoleController), "post:Insert"),
	)

	beego.AddNamespace(ns)
}

