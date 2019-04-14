package services

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"Service-Content/models"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
)

type UserService struct{
	BaseService
	ctx *context.Context
	m  *models.UserModel
}

var (
	UserStatusMap = map[int8]interface{}{
		constants.UserStatusActive : struct {}{},
		constants.UserStatusError : struct {}{},
	}
)

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

func (s *UserService)Insert(param *models.UserModel)  (int64, *errors.ErrMsg){
	if param == nil {
		return constants.DefaultZero, errors.ErrParam
	}
	data,errMsg := s.filterUser(param)
	if errMsg != nil{
		return constants.DefaultZero, errMsg
	}

	user := orm.Params{
		"id": data.Id,
		"status": data.Status,
	}
}

func (s *UserService)filterUser(param *models.UserModel) (*models.UserModel, *errors.ErrMsg){
	if param == nil {
		return nil, errors.ErrParam
	}
	if _,ok := UserStatusMap[param.Status]; !ok {
		return  nil,errors.ErrParam
	}
}