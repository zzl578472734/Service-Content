package initial

import (
	"Service-Content/models"
	"Service-Content/utils/redisclient"
	"github.com/astaxie/beego"
)

func init() {
	initRedis()
	models.InitDB()
	InitLog()
}

func initRedis()  {

	redisMode,err := beego.AppConfig.Int("redisMode")
	if err != nil{
		panic(err)
	}
	redisPort,err := beego.AppConfig.Int("redis.port")
	redisDatabase,err := beego.AppConfig.Int("redis.database")
	if err != nil{
		panic(err)
	}
	redisPassword := beego.AppConfig.String("redis.password")
	redisHost := beego.AppConfig.String("redis.host")

	config := &redisclient.Config{
		Port:	redisPort,
		Database: redisDatabase,
		Password: redisPassword,
		Host: redisHost,
	}
	redisclient.InitRedis(redisMode, config)
}