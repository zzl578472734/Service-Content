package services

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"Service-Content/models"
	"github.com/astaxie/beego/context"
	"strings"
)

type AdminService struct {
	BaseService
	ctx *context.Context
}

func NewAdminService(ctx *context.Context) *AdminService {
	return &AdminService{ctx: ctx}
}

func (s *AdminService) GetMenu(roleId int64) ([]*models.MenuModel, *errors.ErrMsg) {
	if roleId <= constants.DefaultZero {
		return nil, errors.ErrParam
	}

	roleService := NewRoleService(s.ctx)
	role, errMsg := roleService.GetById(roleId)
	if errMsg != nil {
		return nil, errMsg
	}
	if role == nil ||
		role.Id <= constants.DefaultZero ||
		role.MenuIds == constants.DefaultEmptyString {
		return nil, errors.ErrRoleNotExists
	}

	ids := strings.Split(role.MenuIds, ",")
	menuService := NewMenuService(s.ctx)
	list, errMsg := menuService.GetByIds(ids)
	if errMsg != nil {
		return nil, errMsg
	}
	return list, nil
}
