package main

import (
	_ "Service-Content/initial"
	_ "Service-Content/routers"
	"github.com/astaxie/beego"
	"Service-Content/controllers"
	"github.com/astaxie/beego/logs"
)

func main() {
	beego.ErrorController(new(controllers.ErrorController))

	runmode := beego.AppConfig.String("runmode")
	logs.Info("Service-Content is run mode:", runmode)

	beego.Run()
}
