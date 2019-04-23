package models

import (
	"Service-Content/constants"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type UserModel struct {
	BaseModel
	Id         int64     `json:"id" orm:"column(id)"`
	Status     int8      `json:"status" orm:"column(status)"`
	CreateTime string    `json:"create_time" orm:"column(create_time)"`
	UpdateTime string    `json:"update_time" orm:"column(update_time)"`
	O          orm.Ormer `json:"-" orm:"-"`
}

func NewUserModel() *UserModel {
	return &UserModel{O: orm.NewOrm()}
}

func (m *UserModel) TableName() string {
	return TableName(constants.UserTableName)
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

func (m *UserModel) Insert(user *UserModel) (int64, error) {
	id, err := m.O.Insert(user)
	if ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "UserModel.Insert", "Insert", err)
		return constants.DefaultZero, err
	}
	return id, nil
}

func (m *UserModel) ReleaseCache()  {
	key := fmt.Sprintf("%s%d", constants.UserDetailCacheKey, m.Id)

	m.DeleteCache(key)
}