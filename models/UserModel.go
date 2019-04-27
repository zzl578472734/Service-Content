package models

import (
	"Service-Content/constants"
	"Service-Content/utils/redisclient"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"Service-Content/errors"
)

type UserModel struct {
	BaseModel
	Id            int64     `json:"id" orm:"column(id)"`
	Phone         string    `json:"phone" orm:"column(phone)"`
	Status        int8      `json:"status" orm:"column(status)"`
	Username      string    `json:"username" orm:"column(username)"`
	LastLoginIp   string    `json:"last_login_ip" orm:"column(last_login_ip)"`
	LastLoginTime string    `json:"last_login_time" orm:"column(last_login_time)"`
	Password      string    `json:"password" orm:"column(password)"`
	CreateTime    string    `json:"create_time" orm:"column(create_time)"`
	UpdateTime    string    `json:"update_time" orm:"column(update_time)"`
	O             orm.Ormer `json:"-" orm:"-"`
}

func NewUserModel() *UserModel {
	return &UserModel{O: orm.NewOrm()}
}

func (m *UserModel) TableName() string {
	return TableName(constants.UserTableName)
}

func (m *UserModel) GetByPhone(phone string) (*UserModel, error) {
	query := m.O.QueryTable(m.TableName())

	user := new(UserModel)

	query = query.Filter("phone", phone)
	if err := query.One(user); ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "UserModel.GetByPhone", "query.One", err)
		return nil, err
	}
	return user, nil
}

func (m *UserModel) GetByUsername(username string) (*UserModel, error) {

	query := m.O.QueryTable(m.TableName())

	user := new(UserModel)

	query = query.Filter("username", username)
	if err := query.One(user); ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "UserModel.GetByUsername", "query.One", err)
		return nil, err
	}
	return user, nil
}

func (m *UserModel) Update(filter map[string]interface{}, param orm.Params) (int64, error) {

	if len(filter) <= constants.DefaultZero{
		return constants.DefaultZero, errors.ErrUpdateForbid
	}

	query := m.O.QueryTable(m.TableName())

	for key,value := range filter{
		query = query.Filter(key, value)
	}

	effectId,err := query.Update(param)
	if ormErr(err) != nil{
		logs.Warn(constants.DefaultErrorTemplate, "UserModel.Update", "query.Update", err)
		return constants.DefaultZero, err
	}
	return effectId,nil
}

func (m *UserModel) GetById(id int64) (*UserModel, error) {

	query := m.O.QueryTable(m.TableName())

	user := new(UserModel)

	if err := query.Filter("id", id).One(user); ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "UserModel.GetById", "One", err)
		return nil, err
	}

	return user, nil
}

/**
 * 数据插入
 */
func (m *UserModel) Insert(user *UserModel) (int64, error) {
	id, err := m.O.Insert(user)
	if ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "UserModel.Insert", "Insert", err)
		return constants.DefaultZero, err
	}
	return id, nil
}

/**
 * 删除缓存
 */
func (m *UserModel) ReleaseCache() {
	key := fmt.Sprintf("%s%d", constants.UserDetailCacheKey, m.Id)

	redisclient.DeleteCache(key)
}
