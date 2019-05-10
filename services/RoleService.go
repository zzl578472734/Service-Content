package services

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"Service-Content/models"
	"github.com/astaxie/beego/context"
	"Service-Content/utils/redisclient"
)

type RoleService struct {
	BaseService
	ctx   *context.Context
	model *models.RoleModel
}

func NewRoleService(ctx *context.Context) *RoleService {
	return &RoleService{ctx: ctx, model: models.NewRoleModel()}
}

/**
 * 根据角色id获取角色列表
 */
func (s *RoleService) GetById(id int64) (*models.RoleModel, *errors.ErrMsg) {
	if id <= constants.DefaultZero {
		return nil, errors.ErrParam
	}

	role, err := s.model.GetById(id)
	if err != nil {
		return nil, errors.ErrQueryError
	}

	if role == nil || role.Status != constants.RoleStatusSuccess {
		return nil, nil
	}

	return role, nil
}

/**
 * 获取角色列表
 */
func (s *RoleService) GetEffectList() ([]*models.RoleModel, *errors.ErrMsg){

	list := make([]*models.RoleModel, constants.DefaultZero)

	err := redisclient.GetCache(constants.RoleCacheEffectList, list)
	if err == nil && len(list) > constants.DefaultZero{
		return list,nil
	}

	filter := map[string]interface{}{
		"status": constants.RoleStatusSuccess,
	}

	list,_,err = s.model.List(filter, constants.DefaultZero,constants.DefaultZero)
	if err != nil{
		return nil,errors.ErrQueryError
	}

	redisclient.SetCache(constants.RoleCacheEffectList, list)

	return list,nil
}
