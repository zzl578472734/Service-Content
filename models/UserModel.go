package models

import (
	"Service-Content/constants"
	"github.com/astaxie/beego/orm"
)

type UserModel struct {
	BaseModel
	Id     int64 `json:"id" orm:"column(id)"`
	Status int8  `json:"status" orm:"column(status)"`

	O orm.Ormer `json:"-"`
}

func NewUserModel() *UserModel {
	return &UserModel{O: orm.NewOrm()}
}

func (m *UserModel) TableName() string {
	return constants.UserTableName
}

func (m *UserModel) GetById(id int64) (*UserModel, error) {

	query := m.O.QueryTable(m.TableName())

	user := new(UserModel)

	if err := query.One(user); err != nil {
		return nil, err
	}

	return user, nil
}
