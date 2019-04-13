package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

type BaseModel struct {
}

func InitDB() {

	databaseDriver := beego.AppConfig.String("database.driver")
	databaseUser := beego.AppConfig.String("database.user")
	databasePassword := beego.AppConfig.String("database.password")
	databaseHost := beego.AppConfig.String("database.host")
	databasePort := beego.AppConfig.String("database.port")
	databaseDb := beego.AppConfig.String("database.db")
	databaseCharset := beego.AppConfig.String("database.charset")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		databaseUser,
		databasePassword,
		databaseHost,
		databasePort,
		databaseDb,
		databaseCharset,
	)

	err := orm.RegisterDriver(databaseDriver, orm.DRMySQL)
	err = orm.RegisterDataBase("default", databaseDriver, dataSource)
	if err != nil {
		logs.Info(err)
		return
	}

	orm.RegisterModel(new(UserModel))
}
