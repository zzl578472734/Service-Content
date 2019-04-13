package main

import (
	_ "Service-Content/initial"
	_ "Service-Content/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
