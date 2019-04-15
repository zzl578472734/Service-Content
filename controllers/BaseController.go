package controllers

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"github.com/astaxie/beego"
)

type ApiResponse struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type BaseController struct {
	beego.Controller
}

func (c *BaseController)ApiSuccessReturn(data interface{})  {
	resp := new(ApiResponse)

	resp.Data = data
	resp.Code = constants.DefaultApiSuccessCode
	resp.Msg = constants.DefaultEmptyString

	c.Data["json"] = resp

	c.ServeJSON()
	c.StopRun()
	return
}

func (c *BaseController)ApiErrorReturn(errMsg *errors.ErrMsg){
	c.Data["json"] = errMsg
	c.ServeJSON()
	c.StopRun()
	return
}

func (c *BaseController)GetRequestBody(param interface{})  {
	//c.Ctx.Input.RequestBody
}