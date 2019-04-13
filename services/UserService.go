package services

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"Service-Content/models"
	"github.com/astaxie/beego/context"
)

type UserService struct{
	BaseService
	ctx *context.Context
	m  *models.UserModel
}

func NewUserService(ctx *context.Context)  *UserService{
	return &UserService{ctx:ctx, m: models.NewUserModel()}
}

func (s *UserService) Detail(id int64) (*models.UserModel, *errors.ErrMsg){
	if id <= constants.DefaultZero{
		return nil, errors.ErrParam
	}

	user,err := s.m.GetById(id)
	if err != nil{
		return nil,errors.ErrQueryError
	}

	if user == nil || user.Status != constants.UserStatusActive {
		return nil,errors.ErrQueryRecordNotExists
	}
	return user,nil
}
