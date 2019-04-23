package initial

import (
	"Service-Content/models"
	"Service-Content/utils"
)

func init()  {
	utils.InitRedis()
	models.InitDB()
	InitLog()
}