package models

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"Service-Content/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
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

func OrmErr(err error) error {
	if err != orm.ErrNoRows {
		return err
	}
	return nil
}

func (m *BaseModel) SetCache(key string, value interface{}, expiration ...time.Duration) error {
	if key == constants.DefaultEmptyString || len(key) <= constants.DefaultZero {
		return errors.ErrCacheKey
	}

	if value == nil {
		return errors.ErrCacheValue
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		logs.Info(constants.DefaultErrorTemplate, "BaseModel.SetCache", "Marshal", err)
		return err
	}
	if len(bytes) <= constants.DefaultZero {
		return nil
	}

	expire := constants.DefaultCacheExpire
	if len(expiration) > constants.DefaultZero {
		expire = expiration[0]
	}

	err = utils.RedisClient.Set(key, bytes, expire).Err()
	if err != nil {
		logs.Info(constants.DefaultErrorTemplate, "BaseModel.SetCache", "Set", err)
		return err
	}
	return nil
}

func (m *BaseModel) GetCache(key string, value interface{}) error {
	if key == constants.DefaultEmptyString || len(key) <= constants.DefaultZero {
		return errors.ErrCacheKey
	}

	if value == nil {
		return errors.ErrCacheValue
	}

	bytes, err := utils.RedisClient.Get(key).Bytes()
	if utils.RedisErr(err) != nil {
		logs.Info(constants.DefaultErrorTemplate, "BaseModel.GetCache", "Get", err)
		return err
	}

	if len(bytes) <= constants.DefaultZero {
		return nil
	}

	err = json.Unmarshal(bytes, value)
	if err != nil {
		logs.Info(constants.DefaultErrorTemplate, "BaseModel.GetCache", "Unmarshal", err)
		return err
	}
	return nil
}

func (m *BaseModel) DeleteCache(keys ...string) error {
	if len(keys) <= constants.DefaultZero {
		return errors.ErrCacheKey
	}

	err := utils.RedisClient.Del(keys...).Err()
	if err != nil {
		logs.Info(constants.DefaultErrorTemplate, "BaseModel.DeleteCache", "Del", err)
		return err
	}
	return nil
}
