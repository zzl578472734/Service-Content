package services

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"Service-Content/models"
	"github.com/astaxie/beego/context"
)

type MenuService struct {
	BaseService
	ctx   *context.Context
	model *models.MenuModel
}

func NewMenuService(ctx *context.Context) *MenuService {
	return &MenuService{ctx: ctx, model: models.NewMenuModel()}
}

//func (s *MenuService) List(filter map[string]interface{}, page, pageSize int) ([]*models.MenuModel, int64, *errors.ErrMsg) {
//
//}

func (s *MenuService) GetByIds(ids []string) ([]*models.MenuModel, *errors.ErrMsg) {
	if len(ids) <= constants.DefaultZero {
		return nil, errors.ErrParam
	}

	filter := map[string]interface{}{
		"id__in": ids,
		"status": constants.MenuStatusSuccess,
	}

	list, _, err := s.model.List(filter, constants.DefaultZero, constants.DefaultZero)
	if err != nil {
		return nil, errors.ErrQueryRecordNotExists
	}
	return list, nil
}
