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
}

func (c *BaseController)ApiErrorReturn(errMsg *errors.ErrMsg){
	resp := new(ApiResponse)
	resp.Code = errMsg.Code
	resp.Msg = errMsg.Msg
	resp.Data = constants.DefaultEmptyString
	c.Data["json"] = resp
	c.ServeJSON()
	c.StopRun()
}