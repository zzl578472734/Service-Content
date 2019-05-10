package models

import (
	"Service-Content/constants"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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
		return
	}

	orm.RegisterModel(
		new(UserModel),
		new(AdminModel),
		new(RoleModel),
	)

	// 判断是否是生产环境，打印对应的sql
	runMode := beego.AppConfig.String("runmode")
	if runMode != beego.PROD {
		orm.Debug = true
	}
}

func TableName(tableName string) string {
	databasePrefix := beego.AppConfig.DefaultString("database.prefix", "")
	if len(databasePrefix) > constants.DefaultZero {
		trueTableName := fmt.Sprintf("%s%s", databasePrefix, tableName)
		return trueTableName
	}
	return tableName
}

func ormErr(err error) error {
	if err != orm.ErrNoRows {
		return err
	}
	return nil
}
