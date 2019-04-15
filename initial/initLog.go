package initial

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"github.com/astaxie/beego"
)

func InitLog() {

	config := fmt.Sprintf(`{"filename": "%s", "daily": %t, "maxdays": %d, "rotate": %t, "level":%d}`,
		beego.AppConfig.String("log.filename"),
		beego.AppConfig.DefaultBool("log.daily", true),
		beego.AppConfig.DefaultInt("log.maxdays", 7),
		beego.AppConfig.DefaultBool("log.rotate", true),
		beego.AppConfig.DefaultInt("log.level", 7))

	logs.SetLogger(logs.AdapterMultiFile, config)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}
