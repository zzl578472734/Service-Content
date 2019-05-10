package controllers

import (
	"Service-Content/errors"
	"Service-Content/models"
	"Service-Content/services"
	"Service-Content/vars"
)

type AdminController struct {
	BaseController
}

/**
 * 获取管理员列表
 */
func (c *AdminController) List() {
	param, err := c.GetRequestParam().(*vars.AdminListParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
	}

	service := services.NewAdminService(c.Ctx)
	list, total, page, pageSize, errMsg := service.List(param)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
	}

	c.ApiSuccessReturn(service.BuildListResp(list, total, page, pageSize))
}

/**
 * 获取管理员的详情信息
 */
func (c *AdminController) Detail() {
	param, err := c.GetRequestParam().(*vars.DefaultIdQueryParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
	}

	service := services.NewAdminService(c.Ctx)
	admin, errMsg := service.GetAdmin(param.Id)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
	}

	c.ApiSuccessReturn(admin)
}

/**
 * 修改管理员的密码
 */
func (c *AdminController) ModifyPassword() {
	param, err := c.GetRequestParam().(*vars.AdminModifyPasswordParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
	}

	service := services.NewAdminService(c.Ctx)
	effectId, errMsg := service.ModifyPassword(param)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
	}
	data := map[string]int64{
		"id": effectId,
	}
	c.ApiSuccessReturn(data)
}

/**
 * 获取管理员权限
 */
func (c *AdminController) GetMenu() {
	admin, err := c.Ctx.Input.GetData("admin").(*models.AdminModel)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
	}

	service := services.NewAdminService(c.Ctx)
	list, errMsg := service.GetMenuByRoleId(admin.RoleId)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
	}
	c.ApiSuccessReturn(list)
}

/**
 * 后台用户登录
 */
func (c *AdminController) Login() {
	param, err := c.GetRequestParam().(*vars.AdminLoginParam)

	if !err {
		c.ApiErrorReturn(errors.ErrParam)
	}

	service := services.NewAdminService(c.Ctx)
	admin, errMsg := service.Login(param)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
	}
	c.ApiSuccessReturn(admin)
}

/**
 * 添加管理员
 */
func (c *AdminController) Insert() {
	param, err := c.GetRequestParam().(*vars.AdminParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
	}

	service := services.NewAdminService(c.Ctx)
	id, errMsg := service.Insert(param)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
	}
	data := map[string]int64{
		"id": id,
	}
	c.ApiSuccessReturn(data)
}

/**
 * 激活
 */
func (c *AdminController) Active() {
	param, err := c.GetRequestParam().(*vars.DefaultIdQueryParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
	}

	service := services.NewAdminService(c.Ctx)
	effectId, errMsg := service.Active(param.Id)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
	}
	data := map[string]int64{
		"id": effectId,
	}
	c.ApiSuccessReturn(data)
}

/**
 * 禁用
 */
func (c *AdminController) Disable() {
	param, err := c.GetRequestParam().(*vars.DefaultIdQueryParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
	}

	service := services.NewAdminService(c.Ctx)
	effectId, errMsg := service.Disable(param.Id)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
	}
	data := map[string]int64{
		"id": effectId,
	}
	c.ApiSuccessReturn(data)
}
