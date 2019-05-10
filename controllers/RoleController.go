package controllers

import "Service-Content/services"

type RoleController struct {
	BaseController
}

/**
 * 获取角色列表
 */
func (c *RoleController)GetEffectList()  {

	service := services.NewRoleService(c.Ctx)
	list,errMsg := service.GetEffectList()
	if errMsg != nil{
		c.ApiErrorReturn(errMsg)
	}

	data := map[string]interface{}{
		"list": list,
	}
	c.ApiSuccessReturn(data)
}