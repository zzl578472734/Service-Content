package controllers

import (
	vars2 "Service-Content/thirdparty/wechat/vars"
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

		"AdminController/Login":          AdminLoginAdapter,
		"AdminController/Insert":         AdminParamAdapter,
		"AdminController/Active":         AdminActiveAdapter,
		"AdminController/Disable":        AdminDisableAdapter,
		"AdminController/ModifyPassword": AdminModifyPasswordAdapter,

		"MaterialController/List": BatchGetMaterialAdapter,
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
	return map[string]interface{}{}
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

func AdminActiveAdapter() interface{} {
	return new(vars.DefaultIdQueryParam)
}

func AdminDisableAdapter() interface{} {
	return new(vars.DefaultIdQueryParam)
}

func AdminModifyPasswordAdapter() interface{} {
	return new(vars.AdminModifyPasswordParam)
}

func BatchGetMaterialAdapter() interface{} {
	return new(vars2.BatchGetMaterialParam)
}
