package services

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"Service-Content/models"
	"Service-Content/utils"
	"Service-Content/utils/redisclient"
	"Service-Content/vars"
	"crypto/sha512"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserService struct {
	BaseService
	ctx   *context.Context
	model *models.UserModel
}

var (
//UserStatusMap = map[int8]interface{}{
//	constants.UserStatusActive: struct{}{},
//	constants.UserStatusError:  struct{}{},
//}
)

func NewUserService(ctx *context.Context) *UserService {
	return &UserService{ctx: ctx, model: models.NewUserModel()}
}

/**
 * 用户登录
 */
func (s *UserService) Login(param *vars.UserLoginParam) (*models.UserModel, *errors.ErrMsg) {
	errMsg := s.filterUserLogin(param)
	if errMsg != nil {
		return nil, errMsg
	}

	user := new(models.UserModel)
	// 查询当前当前的用户信息
	user, err := s.model.GetByUsername(param.Username)
	if err != nil {
		return nil, errors.ErrQueryError
	}
	if user == nil || user.Status != constants.UserStatusActive {
		return nil, errors.ErrInputUsernameOrPassword
	}
	if password := fmt.Sprintf("%x", sha512.Sum512([]byte(param.Password))); password != user.Password {
		return nil, errors.ErrInputUsernameOrPassword
	}

	// 更新用户的登录信息
	timeFormat := utils.TimeFormat(time.Now())
	updateParam := orm.Params{
		"last_login_time": timeFormat,
		"last_login_ip" : s.ctx.Input.IP(),
		"update_time":     timeFormat,
	}

	filter := map[string]interface{}{
		"id": user.Id,
	}

	_, err = s.model.Update(filter, updateParam)
	if err != nil {
		return nil, errors.ErrSysBusy
	}

	return user, nil
}

/**
 * 用户搜索功能
 */
func (s *UserService) Search(param *vars.UserSearchParam) ([]*models.UserModel, *errors.ErrMsg) {
	if param == nil {
		return nil, errors.ErrParam
	}

	errMsg := s.filterUserSearch(param)
	if errMsg != nil {
		return nil, errMsg
	}

	//host := beego.AppConfig.String("elastic.host")
	//port, err := beego.AppConfig.Int("elastic.port")
	//if err != nil {
	//	return nil, errors.ErrParam
	//}
	//
	//_ := utils.NewElasticClient(&utils.ElasticConnect{
	//	Host: host,
	//	Port: port,
	//})

	return nil, nil
}

/**
 * 根据id获取详情
 */
func (s *UserService) Detail(id int64) (*models.UserModel, *errors.ErrMsg) {
	if id <= constants.DefaultZero {
		return nil, errors.ErrParam
	}

	user := new(models.UserModel)
	key := fmt.Sprintf("%s%d", constants.UserDetailCacheKey, id)
	err := redisclient.GetCache(key, user)
	if err == nil && user.Id > constants.DefaultZero {
		return user, nil
	}

	user, err = s.model.GetById(id)
	if err != nil {
		return nil, errors.ErrQueryError
	}

	if user == nil || user.Status != constants.UserStatusActive {
		return nil, errors.ErrQueryRecordNotExists
	}

	redisclient.SetCache(key, user)

	return user, nil
}

/**
 * 数据插入
 */
func (s *UserService) Insert(param *vars.UserInsertParam) (int64, *errors.ErrMsg) {
	if param == nil {
		return constants.DefaultZero, errors.ErrParam
	}
	errMsg := s.filterUser(param)
	if errMsg != nil {
		return constants.DefaultZero, errMsg
	}

	// 查询用户是否已经存在
	user, err := s.model.GetByUsername(param.Username)
	if err != nil {
		return constants.DefaultZero, errors.ErrQueryError
	}
	if user != nil && user.Id > constants.DefaultZero {
		return constants.DefaultZero, errors.ErrUserIsExist
	}

	// 生成用户的密码
	password := fmt.Sprintf("%x", sha512.Sum512([]byte(param.Password)))

	data := &models.UserModel{
		Status:   param.Status,
		Username: param.Username,
		Password: password,
	}

	// 添加默认字段
	s.defaultField(data)

	id, err := s.model.Insert(data)
	if err != nil {
		return constants.DefaultZero, errors.ErrInsertError
	}

	// 删除缓存
	s.model.ReleaseCache()

	return id, nil
}

/**
 * 处理用户搜索请求的过滤
 */
func (s *UserService) filterUserSearch(param *vars.UserSearchParam) *errors.ErrMsg {
	if param == nil {
		return errors.ErrParam
	}
	if !utils.IntRange(len([]rune(param.Username)), constants.UserUsernameMinLength, constants.UserUsernameMaxLength) {
		return errors.ErrUsernameLength
	}
	return nil
}

/**
 * 过滤用户的信息
 */
func (s *UserService) filterUser(param *vars.UserInsertParam) *errors.ErrMsg {
	if param == nil {
		return errors.ErrParam
	}
	if !utils.IntRange(len([]rune(param.Username)), constants.UserUsernameMinLength, constants.UserUsernameMaxLength) {
		return errors.ErrUsernameLength
	}
	if !utils.IntRange(len([]rune(param.Password)), constants.UserPasswordMinLength, constants.UserPasswordMaxLength) {
		return errors.ErrPasswordLength
	}
	return nil
}

/**
 * 登录验证
 */
func (s *UserService) filterUserLogin(param *vars.UserLoginParam) *errors.ErrMsg {
	if param == nil {
		return errors.ErrParam
	}
	if !utils.IntRange(len([]rune(param.Username)), constants.UserUsernameMinLength, constants.UserUsernameMaxLength) {
		return errors.ErrUsernameLength
	}
	if !utils.IntRange(len([]rune(param.Password)), constants.UserPasswordMinLength, constants.UserPasswordMaxLength) {
		return errors.ErrPasswordLength
	}
	return nil
}

/**
 * 新增默认的字段
 */
func (s *UserService) defaultField(user *models.UserModel) {
	timeFormat := utils.TimeFormat(time.Now())
	user.CreateTime = timeFormat
	user.UpdateTime = timeFormat
	user.Status = constants.UserStatusActive
}
