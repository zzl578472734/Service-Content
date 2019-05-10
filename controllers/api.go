package controllers

import (
	"Service-Content/vars"
	"fmt"
)

type funcAdapter func() interface{}

var (
	apiRequestAdapter = map[string]funcAdapter{
		"UserController/Detail": DefaultQueryParamAdapter,
		"UserController/Login":  UserLoginAdapter,
		"UserController/Insert": UserInsertAdapter,
		"UserController/Search": UserSearchAdapter,

		"AdminController/Login": AdminLoginAdapter,
		"AdminController/Insert": AdminParamAdapter,
	}
)

func getApiRequestAdapter(c *BaseController) interface{} {
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
	return new(vars.DefaultIdQueryParam)
}


func UserSearchAdapter() interface{} {
	return new(vars.UserSearchParam)
}

func UserInsertAdapter() interface{} {
	return new(vars.UserInsertParam)
}

func UserLoginAdapter() interface{} {
	return new(vars.UserLoginParam)
}

func AdminLoginAdapter() interface{} {
	return new(vars.AdminLoginParam)
}

func AdminParamAdapter() interface{} {
	return new(vars.AdminParam)
}