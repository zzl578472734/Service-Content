package services

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"Service-Content/models"
	"github.com/astaxie/beego/context"
	"Service-Content/utils"
	"time"
	"fmt"
)

type UserService struct {
	BaseService
	ctx *context.Context
	m   *models.UserModel
}

var (
	UserStatusMap = map[int8]interface{}{
		constants.UserStatusActive: struct{}{},
		constants.UserStatusError:  struct{}{},
	}
)

func NewUserService(ctx *context.Context) *UserService {
	return &UserService{ctx: ctx, m: models.NewUserModel()}
}

func (s *UserService) Detail(id int64) (*models.UserModel, *errors.ErrMsg) {
	if id <= constants.DefaultZero {
		return nil, errors.ErrParam
	}

	user := new(models.UserModel)
	key := fmt.Sprintf("%s%d",constants.UserDetailCacheKey, id)
	err := s.m.GetCache(key, user)
	if err == nil && user.Id > constants.DefaultZero{
		return user,nil
	}

	user, err = s.m.GetById(id)
	if err != nil {
		return nil, errors.ErrQueryError
	}

	if user == nil || user.Status != constants.UserStatusActive {
		return nil, errors.ErrQueryRecordNotExists
	}

	s.m.SetCache(key, user)

	return user, nil
}

func (s *UserService) Insert(param *models.UserModel) (int64, *errors.ErrMsg) {
	if param == nil {
		return constants.DefaultZero, errors.ErrParam
	}
	data, errMsg := s.filterUser(param)
	if errMsg != nil {
		return constants.DefaultZero, errMsg
	}

	user := &models.UserModel{
		Status: data.Status,
	}

	id, errMsg := s.defaultInsert(user)
	if errMsg !=nil{
		return constants.DefaultZero, errMsg
	}
	return id,nil
}

func (s *UserService) filterUser(param *models.UserModel) (*models.UserModel, *errors.ErrMsg) {
	if param == nil {
		return nil, errors.ErrParam
	}
	if _, ok := UserStatusMap[param.Status]; !ok {
		return nil, errors.ErrParam
	}
	return param, nil
}

func (s *UserService) defaultInsert(user *models.UserModel) (int64, *errors.ErrMsg) {
	if user == nil{
		return constants.DefaultZero, errors.ErrParam
	}

	// 添加默认字段
	s.defaultField(user)

	id, err := s.m.Insert(user)
	if err != nil{
		return constants.DefaultZero, errors.ErrInsertError
	}
	return id,nil
}

func (s *UserService) defaultField(user *models.UserModel)  {
	timeFormat := utils.FormatTime(time.Now())
	user.CreateTime = timeFormat
	user.UpdateTime = timeFormat
}
