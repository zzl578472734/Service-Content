package controllers

import "Service-Content/services"

type AdminController struct {
	BaseController
}

/**
 * 獲取用戶的權限
 */
func (c *AdminController)GetMenu()  {

	userId,_ := c.GetInt64("user_id")

	service := services.NewAdminService(c.Ctx)
	list,errMsg := service.GetMenu(userId)
	if errMsg != nil{
		c.ApiErrorReturn(errMsg)
	}
	c.ApiSuccessReturn(list)
}
