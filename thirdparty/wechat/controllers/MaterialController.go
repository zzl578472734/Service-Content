package controllers

import "Service-Content/thirdparty/wechat/services"

type MaterialController struct {
	BaseController
}

func (c *MaterialController) List()  {

	service := services.NewMaterialService(c.Ctx)

}
