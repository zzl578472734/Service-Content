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

	// 素材管理
	initMaterialRouter()
}


func initAdminRouter()  {

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/login", new(controllers.AdminController), "post:Login"),
		beego.NSRouter("/insert", new(controllers.AdminController), "post:Insert"),
		beego.NSRouter("/active", new(controllers.AdminController), "post:Active"),
		beego.NSRouter("/disable", new(controllers.AdminController), "post:Disable"),
		beego.NSRouter("/modifyPassword", new(controllers.AdminController), "post:ModifyPassword"),
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

}

func initMaterialRouter()  {

	ns := beego.NewNamespace("/material",
		beego.NSRouter("/list", new(controllers.MaterialController), "post:List"),
	)

	beego.AddNamespace(ns)
}
