package models

import (
	"Service-Content/constants"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MenuModel struct {
	BaseModel  `json:"-" orm:"-"`
	Id         int64     `json:"id" orm:"column(id)"`
	Pid        int64     `json:"pid" orm:"column(pid)"`
	Url        string    `json:"url" orm:"column(url)"`
	Name       string    `json:"name" orm:"column(name)"`
	Icon       string    `json:"icon" orm:"column(icon)"`
	Status     int8      `json:"status" orm:"column(status)"`
	CreateTime string    `json:"create_time" orm:"column(create_time)"`
	UpdateTime string    `json:"update_time" orm:"column(update_time)"`
	O          orm.Ormer `json:"-" orm:"-"`
}

func NewMenuModel() *MenuModel {
	return &MenuModel{O: orm.NewOrm()}
}

func (m *MenuModel) TableName() string {
	return TableName(constants.MenuTableName)
}

func (m *MenuModel) List(filter map[string]interface{}, page, pageSize int) ([]*MenuModel, int64, error) {
	query := m.O.QueryTable(m.TableName())

	if len(filter) > constants.DefaultZero {
		for key, value := range filter {
			query = query.Filter(key, value)
		}
	}

	total, err := query.Count()
	if ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "MenuModel.List", "query.Count", err)
		return nil, constants.DefaultZero, err
	}

	list := make([]*MenuModel, constants.DefaultZero)

	query = query.Limit(pageSize).Offset((page - 1) * pageSize)
	if _, err = query.All(&list); ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "MenuModel.All", "query.Limit", err)
		return nil, constants.DefaultZero, err
	}
	return list, total, nil
}
