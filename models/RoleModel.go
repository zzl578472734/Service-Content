package models

import (
	"Service-Content/constants"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type RoleModel struct {
	BaseModel  `json:"-" orm:"-"`
	Id         int64     `json:"id" orm:"column(id)"`
	Name       string    `json:"name" orm:"column(name)"`
	MenuIds    string    `json:"menu_ids" orm:"column(menu_ids)"`
	Status     int8      `json:"status" orm:"column(status)"`
	CreateTime string    `json:"create_time" orm:"column(create_time)"`
	UpdateTime string    `json:"update_time" orm:"column(update_time)"`
	O          orm.Ormer `json:"-" orm:"-"`
}

func NewRoleModel() *RoleModel {
	return &RoleModel{O: orm.NewOrm()}
}

func (m *RoleModel) TableName() string {
	return TableName(constants.RoleTableName)
}

func (m *RoleModel) GetById(id int64) (*RoleModel, error) {
	query := m.O.QueryTable(m.TableName())

	role := new(RoleModel)

	query = query.Filter("id", id)
	if err := query.One(role); ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "RoleModel.GetById", "query.One", err)
		return nil, err
	}
	return role, nil
}
