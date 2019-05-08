package services

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"Service-Content/models"
	"github.com/astaxie/beego/context"
)

type RoleService struct {
	BaseService
	ctx   *context.Context
	model *models.RoleModel
}

func NewRoleService(ctx *context.Context) *RoleService {
	return &RoleService{ctx: ctx, model: models.NewRoleModel()}
}

func (s *RoleService) GetById(id int64) (*models.RoleModel, *errors.ErrMsg) {
	if id <= constants.DefaultZero {
		return nil,errors.ErrParam
	}

	role,err := s.model.GetById(id)
	if err != nil{
		return nil,errors.ErrQueryError
	}

	if role == nil || role.Status != constants.RoleStatusSuccess{
		return nil,nil
	}

	return role,nil
}
