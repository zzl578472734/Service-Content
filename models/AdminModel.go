package models

import (
	"Service-Content/constants"
	"Service-Content/utils/redisclient"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type AdminModel struct {
	Id            int64     `json:"id" orm:"column(id)"`
	Account       string    `json:"account" orm:"column(account)"`
	Name          string    `json:"name" orm:"column(name)"`
	Phone         string    `json:"phone" orm:"column(phone)"`
	Email         string    `json:"email" orm:"column(email)"`
	Password      string    `json:"password" orm:"column(password)"`
	Salt          string    `json:"salt" orm:"column(salt)"`
	RoleId        int64     `json:"role_id" orm:"column(role_id)"`
	Status        int8      `json:"status" orm:"column(status)"`
	LastLoginIp   string    `json:"last_login_ip" orm:"column(last_login_ip)"`
	LastLoginTime string    `json:"last_login_time" orm:"column(last_login_time)"`
	CreateTime    string    `json:"create_time" orm:"column(create_time)"`
	UpdateTime    string    `json:"update_time" orm:"column(update_time)"`
	O             orm.Ormer `json:"-" orm:"-"`
	BaseModel     `json:"-" orm:"-"`
}

func NewAdminModel() *AdminModel {
	return &AdminModel{O: orm.NewOrm()}
}

func (m *AdminModel) TableName() string {
	return TableName(constants.AdminTableName)
}

/**
 * 获取列表信息
 */
func (m *AdminModel) List(filter map[string]interface{}, page, pageSize int) ([]*AdminModel, int64, error) {
	query := m.O.QueryTable(m.TableName())

	if len(filter) > constants.DefaultZero{
		for key,value := range filter{
			query = query.Filter(key, value)
		}
	}

	total,err := query.Count()
	if ormErr(err) != nil{
		logs.Error(constants.DefaultErrorTemplate, "AdminModel.List", "query.Count", err)
		return nil,constants.DefaultZero,err
	}

	list := make([]*AdminModel, constants.DefaultZero)
	query = query.Limit(pageSize).Offset((page - 1) * pageSize)
	if _,err := query.All(&list); ormErr(err) != nil{
		logs.Error(constants.DefaultErrorTemplate, "AdminModel.List", "query.All", err)
		return nil,constants.DefaultZero,err
	}
	return list,total,nil
}

/**
 * 根据账号获取管理员信息
 */
func (m *AdminModel) GetByAccount(account string) (*AdminModel, error) {
	query := m.O.QueryTable(m.TableName())

	admin := new(AdminModel)

	query = query.Filter("account", account)
	if err := query.One(admin); ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "AdminModel.GetByAccount", "query.One", err)
		return nil, err
	}
	return admin, nil
}

/**
 * 更新操作
 */
func (m *AdminModel) Update(filter map[string]interface{}, data orm.Params) (int64, error) {
	query := m.O.QueryTable(m.TableName())

	if len(filter) > constants.DefaultZero {
		for key, value := range filter {
			query = query.Filter(key, value)
		}
	}

	id, err := query.Update(data)
	if ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "AdminModel.Update", "query.Update", err)
		return constants.DefaultZero, nil
	}
	return id, nil
}

/**
 * 根据id获取列表
 */
func (m *AdminModel) GetById(id int64) (*AdminModel, error) {
	query := m.O.QueryTable(m.TableName())

	admin := new(AdminModel)

	query = query.Filter("id", id)

	if err := query.One(admin); ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "AdminModel.GetById", "query.One", err)
		return nil, err
	}

	return admin, nil
}

/**
 * 添加操作
 */
func (m *AdminModel) Insert(data *AdminModel) (int64, error) {
	id, err := m.O.Insert(data)
	if ormErr(err) != nil {
		logs.Error(constants.DefaultErrorTemplate, "AdminModel.Insert", "O.Insert", err)
		return constants.DefaultZero, err
	}
	return id, nil
}

func (m *AdminModel) DeleteCache() {
	key1 := fmt.Sprintf("%s%d", constants.AdminCacheGetAdmin, m.Id)

	redisclient.DeleteCache(key1)
}
