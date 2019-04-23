package controllers

import (
	"Service-Content/vars"
	"fmt"
)

type funcAdapter func() interface{}

var (
	apiRequestAdapter = map[string]funcAdapter{
		"UserController/Detail": DefaultQueryParamAdapter,
	}
)

func getApiRequestAdapter(c *BaseController) interface{}{
	controllerName, actionName := c.GetControllerAndAction()

	completeAdapter := fmt.Sprintf("%s/%s", controllerName, actionName)

	adapter, exists := apiRequestAdapter[completeAdapter]
	if exists {
		return adapter()
	}

	// 使用map处理

	return adapter()
}

func DefaultQueryParamAdapter() interface{} {
	return new(vars.DefaultQueryParam)
}
