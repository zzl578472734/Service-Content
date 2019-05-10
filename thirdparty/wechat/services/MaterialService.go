package services

import (
	"Service-Content/errors"
	"Service-Content/models"
	"github.com/astaxie/beego/context"
)

type MaterialService struct {
	BaseService
	ctx   *context.Context
	model *models.WxMaterialModel
}

func NewMaterialService(ctx *context.Context) *MaterialService {
	return &MaterialService{ctx: ctx, model: models.NewWxMaterialModel()}
}

func (s *MaterialService) List() ([]*models.WxMaterialModel, *errors.ErrMsg) {

}

func (s *MaterialService) Insert() (*models.WxMaterialModel, *errors.ErrMsg) {

}
