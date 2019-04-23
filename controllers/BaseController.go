package controllers

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ApiResponse struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ApiRequestParam struct {
	UserId     int64       `json:"user_id"`
	AppId      string      `json:"app_id"`
	AppToken   string      `json:"app_token"`
	Nonce      string      `json:"nonce"`
	Timestamps int64       `json:"timestamps"`
	Param      interface{} `json:"param"`
}

const ApiRequestBody = "ApiRequestBody"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) ApiSuccessReturn(data interface{}) {
	resp := new(ApiResponse)

	resp.Data = data
	resp.Code = constants.DefaultApiSuccessCode
	resp.Msg = constants.DefaultApiSuccessMsg

	c.Data["json"] = resp

	c.ServeJSON()
	c.StopRun()
	return
}

func (c *BaseController) ApiErrorReturn(errMsg *errors.ErrMsg) {
	c.Data["json"] = errMsg

	c.ServeJSON()
	c.StopRun()
	return
}

func (c *BaseController) Prepare() {
	filterPrepare(c)
}

func (c *BaseController) getRequestBodyParam() {
	request := new(ApiRequestParam)
	request.Param = getApiRequestAdapter(c)

	body := c.Ctx.Input.RequestBody

	if len(body) <= constants.DefaultZero {
		c.ApiErrorReturn(errors.ErrParam)
		return
	}

	if beego.AppConfig.String("runmode") != beego.PROD {
		logs.Error("Api request url:%s request body:%s", c.Ctx.Input.URL(), string(body))
	}

	if err := json.Unmarshal(body, request); err != nil {
		logs.Error(constants.DefaultErrorTemplate, "BaseController。decodeRequestBody", "json.Unmarshal", err)
		c.ApiErrorReturn(errors.ErrSysBusy)
		return
	}

	c.Ctx.Input.SetData(ApiRequestBody, request)
}

func (c *BaseController) GetRequestParam() interface{} {
	request, err := c.Ctx.Input.GetData(ApiRequestBody).(*ApiRequestParam)

	if !err {
		logs.Error(constants.DefaultErrorTemplate, "BaseController.GetRequestParam", "GetData", errors.ErrInterfaceAssert)
		return nil
	}
	return request.Param
}
