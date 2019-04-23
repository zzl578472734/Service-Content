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

func getApiRequestAdapter(c *BaseController) {
	controllerName, actionName := c.GetControllerAndAction()

	completeAdapter := fmt.Sprintf("%s/%s", controllerName, actionName)

	adapter, exists := apiRequestAdapter[completeAdapter]
	if exists {
		adapter()
	}
}

func DefaultQueryParamAdapter() interface{} {
	return new(vars.DefaultQueryParam)
}
