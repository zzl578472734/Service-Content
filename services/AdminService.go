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
	"strings"
	"time"
)

type AdminService struct {
	BaseService
	ctx   *context.Context
	model *models.AdminModel
}

func NewAdminService(ctx *context.Context) *AdminService {
	return &AdminService{ctx: ctx, model: models.NewAdminModel()}
}

/**
 * 获取管理员列表
 */
func (s *AdminService) List(param *vars.AdminListParam) ([]*models.AdminModel, int64, int, int, *errors.ErrMsg) {
	if param == nil {
		return nil, constants.DefaultZero, constants.DefaultZero, constants.DefaultZero, errors.ErrParam
	}

	page := utils.FilterPage(param.Page)
	pageSize := utils.FilterPageSize(param.PageSize)

	filter := s.filterList(param)
	list, total, err := s.model.List(filter, page, pageSize)
	if err != nil {
		return nil, constants.DefaultZero, constants.DefaultZero, constants.DefaultZero, errors.ErrQueryError
	}

	return list, total, page, pageSize, nil
}

/**
 * 拼接请求参数
 */
func (s *AdminService) filterList(param *vars.AdminListParam) map[string]interface{} {
	filter := make(map[string]interface{}, constants.DefaultZero)

	if param.Name != constants.DefaultEmptyString {
		filter["name__contains"] = param.Name
	}
	if param.Phone != constants.DefaultEmptyString {
		filter["phone__contains"] = param.Phone
	}

	return filter
}

/**
 * 绑定列表返回的参数
 */
func (s *AdminService) BuildListResp(list []*models.AdminModel, total int64, page, pageSize int) map[string]interface{} {

	return map[string]interface{}{
		"list":      list,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}
}

/**
 * 根据管理员id,获取管理员信息
 */
func (s *AdminService) GetAdmin(id int64) (*models.AdminModel, *errors.ErrMsg) {
	if id <= constants.DefaultZero {
		return nil, errors.ErrParam
	}

	admin := new(models.AdminModel)
	cacheKey := fmt.Sprintf("%s%d", constants.AdminCacheGetAdmin, id)
	err := redisclient.GetCache(cacheKey, admin)
	if err == nil && admin.Id > constants.DefaultZero {
		return admin, nil
	}

	admin, err = s.model.GetById(id)
	if err != nil {
		return nil, errors.ErrQueryError
	}

	redisclient.SetCache(cacheKey, admin)

	return admin, nil
}

/**
 * 根据角色id获取管理员的信息
 */
func (s *AdminService) GetMenuByRoleId(roleId int64) ([]*models.MenuModel, *errors.ErrMsg) {
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

/**
 * 后台登录
 */
func (s *AdminService) Login(param *vars.AdminLoginParam) (map[string]interface{}, *errors.ErrMsg) {
	errMsg := s.filterAdminLogin(param)
	if errMsg != nil {
		return nil, errMsg
	}

	// 查询当前当前的用户信息
	admin, err := s.model.GetByAccount(param.Account)
	if err != nil {
		return nil, errors.ErrQueryError
	}
	if admin == nil ||
		admin.Id <= constants.DefaultZero {
		return nil, errors.ErrInputAccountOrPassword
	}
	if admin.Status != constants.AdminStatusSuccess{
		return nil,errors.ErrAdminStatusErr
	}

	summary := fmt.Sprintf("%s%s", param.Password, admin.Salt)
	if password := s.getEncryptPassword(summary); password != admin.Password {
		return nil, errors.ErrInputAccountOrPassword
	}

	// 更新用户的登录信息
	timeFormat := utils.TimeFormat(time.Now())
	updateParam := orm.Params{
		"last_login_time": timeFormat,
		"last_login_ip":   s.ctx.Input.IP(),
		"update_time":     timeFormat,
	}

	filter := map[string]interface{}{
		"id": admin.Id,
	}

	_, err = s.model.Update(filter, updateParam)
	if err != nil {
		return nil, errors.ErrSysBusy
	}

	data := map[string]interface{}{
		"id": admin.Id,
		"account": admin.Account,
		"name": admin.Name,
		"phone": admin.Phone,
		"email": admin.Email,
		"role_id": admin.RoleId,
		"status": admin.Status,
		"last_login_ip": admin.LastLoginIp,
		"last_login_time": admin.LastLoginTime,
		"create_time": admin.CreateTime,
		"update_time": admin.UpdateTime,
	}

	return data, nil
}

/**
 * 生成管理员密码
 */
func (s *AdminService) getEncryptPassword(password string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(password)))
}

/**
 * 登录验证
 */
func (s *AdminService) filterAdminLogin(param *vars.AdminLoginParam) *errors.ErrMsg {
	if param == nil {
		return errors.ErrParam
	}
	if !utils.IntRange(len([]rune(param.Account)), constants.AdminAccountMinLength, constants.AdminAccountMaxLength) {
		return errors.ErrUsernameLength
	}
	if !utils.IntRange(len([]rune(param.Password)), constants.AdminPasswordMinLength, constants.AdminPasswordMaxLength) {
		return errors.ErrPasswordLength
	}
	return nil
}

/**
 * 添加管理员
 */
func (s *AdminService) Insert(param *vars.AdminParam) (int64, *errors.ErrMsg) {
	if param == nil {
		return constants.DefaultZero, errors.ErrParam
	}

	errMsg := s.filterAdmin(param)
	if errMsg != nil {
		return constants.DefaultZero, errMsg
	}

	admin, err := s.model.GetByAccount(param.Account)
	if err != nil {
		return constants.DefaultZero, errors.ErrQueryError
	}
	if admin != nil && admin.Id > constants.DefaultZero {
		return constants.DefaultZero, errors.ErrAdminIsExist
	}

	// 验证当前的角色是否存在
	roleService := NewRoleService(s.ctx)
	role, errMsg := roleService.GetById(param.RoleId)
	if errMsg != nil {
		return constants.DefaultZero, errMsg
	}
	if role == nil ||
		role.Id <= constants.DefaultZero ||
		role.Status != constants.RoleStatusSuccess {
		return constants.DefaultZero, errors.ErrRoleNotExists
	}

	salt := utils.GetRandomString(constants.AdminSaltLength)
	summary := fmt.Sprintf("%s%s", param.Password, salt)
	password := s.getEncryptPassword(summary)

	data := &models.AdminModel{
		Account:  param.Account,
		Password: password,
		RoleId:   param.RoleId,
		Salt: salt,
		Email: param.Email,
		Phone: param.Phone,
		Name: param.Name,
		LastLoginIp: s.ctx.Input.IP(),
	}

	s.defaultField(data)

	id, err := s.model.Insert(data)
	if err != nil {
		return constants.DefaultZero, errors.ErrInsertError
	}
	return id, nil
}

/**
 * 验证管理员信息
 */
func (s *AdminService) filterAdmin(param *vars.AdminParam) *errors.ErrMsg {
	if param == nil {
		return errors.ErrParam
	}
	if !utils.IntRange(len(param.Account), constants.AdminAccountMinLength, constants.AdminAccountMaxLength) {
		return errors.ErrAdminAccountErr
	}
	if !utils.IntRange(len(param.Password), constants.AdminPasswordMinLength, constants.AdminPasswordMaxLength) {
		return errors.ErrAdminPasswordErr
	}

	// TODO
	// 验证用户的姓名
	// 邮箱
	// 手机号码

	return nil
}

/**
 * 添加默认的字段
 */
func (s *AdminService) defaultField(data *models.AdminModel) {
	timeNow := utils.TimeFormat(time.Now())
	data.Status = constants.AdminStatusDefault
	data.CreateTime = timeNow
}

/**
 * 激活管理员
 */
func (s *AdminService) Active(id int64) (int64, *errors.ErrMsg) {
	if id <= constants.DefaultZero {
		return constants.DefaultZero, errors.ErrParam
	}

	// 查询当前管理员是否存在
	admin, errMsg := s.GetAdmin(id)
	if errMsg != nil {
		return constants.DefaultZero, errMsg
	}

	if admin == nil ||
		admin.Id <= constants.DefaultZero {
		return constants.DefaultZero, errors.ErrAdminNotExist
	}

	switch admin.Status {
	case constants.AdminStatusSuccess:
		return constants.DefaultZero, errors.ErrAdminIsActive
	}
	
	filter := map[string]interface{}{
		"id": id,
	}

	timeNow := utils.TimeFormat(time.Now())
	update := orm.Params{
		"status":      constants.AdminStatusSuccess,
		"update_time": timeNow,
	}

	effectId, err := s.model.Update(filter, update)
	if err != nil {
		return constants.DefaultZero, errors.ErrSysBusy
	}

	// 删除cache
	admin.DeleteCache()

	return effectId, nil
}

/**
 * 禁用管理员
 */
func (s *AdminService) Disable(id int64) (int64, *errors.ErrMsg) {
	if id <= constants.DefaultZero {
		return constants.DefaultZero, errors.ErrParam
	}

	admin, err := s.model.GetById(id)
	if err != nil {
		return constants.DefaultZero, errors.ErrQueryError
	}

	if admin == nil ||
		admin.Id <= constants.DefaultZero ||
		admin.Status != constants.AdminStatusSuccess {
		return constants.DefaultZero, errors.ErrAdminNotExist
	}

	filter := map[string]interface{}{
		"id": admin.Id,
	}

	timeNow := utils.TimeFormat(time.Now())
	update := orm.Params{
		"status":      constants.AdminStatusDefault,
		"update_time": timeNow,
	}

	effectId, err := s.model.Update(filter, update)
	if err != nil {
		return constants.DefaultZero, errors.ErrSysBusy
	}

	admin.DeleteCache()

	return effectId, nil
}

/**
 * 修改用户的密码
 */
func (s *AdminService) ModifyPassword(param *vars.AdminModifyPasswordParam) (int64, *errors.ErrMsg) {
	errMsg := s.filterModifyPassword(param)
	if errMsg != nil {
		return constants.DefaultZero, errMsg
	}

	admin, err := s.model.GetById(param.Id)
	if err != nil {
		return constants.DefaultZero, errors.ErrQueryError
	}
	if admin == nil || admin.Id <= constants.DefaultZero || admin.Status != constants.AdminStatusSuccess {
		return constants.DefaultZero, errors.ErrQueryRecordNotExists
	}

	filter := map[string]interface{}{
		"id": admin.Id,
	}

	summary := fmt.Sprintf("%s%s", param.Password, admin.Salt)
	password := s.getEncryptPassword(summary)

	timeNow := utils.TimeFormat(time.Now())

	update := orm.Params{
		"password":    password,
		"update_time": timeNow,
	}

	effectId, err := s.model.Update(filter, update)
	if err != nil {
		return constants.DefaultZero, errors.ErrSysBusy
	}

	admin.DeleteCache()
	return effectId, nil
}

/**
 * 处理用户修改密码的参数
 */
func (s *AdminService) filterModifyPassword(param *vars.AdminModifyPasswordParam) *errors.ErrMsg {
	if param == nil {
		return errors.ErrParam
	}

	if param.Id <= constants.DefaultZero {
		return errors.ErrParam
	}

	if !utils.IntRange(len(param.Password), constants.AdminPasswordMinLength, constants.AdminPasswordMaxLength) {
		return errors.ErrPasswordLength
	}
	return nil
}
