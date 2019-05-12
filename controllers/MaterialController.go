package controllers

import (
	"github.com/astaxie/beego/logs"
	"Service-Content/thirdparty/wechat/services"
	vars2 "Service-Content/thirdparty/wechat/vars"
	"Service-Content/errors"
)

type MaterialController struct {
	BaseController
}

func (c *MaterialController) List()  {
	param,err := c.GetRequestParam().(*vars2.BatchGetMaterialParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
	}

	service := services.NewMaterialService(c.Ctx)
	data := service.BatchGetMaterial(param)
	logs.Info(data)
}


