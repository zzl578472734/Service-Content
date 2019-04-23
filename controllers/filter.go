package controllers

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"fmt"
	"github.com/astaxie/beego/logs"
)

var (
	allowController = map[string][]funcFilter{
		"UserController/*": {UserAuth},
	}
)

type funcFilter func(c *BaseController)


func filterPrepare(c *BaseController) {
	controllerName, actionName := c.GetControllerAndAction()

	completePath := fmt.Sprintf("%s/%s", controllerName, actionName)

	filters, exists := allowController[completePath]
	switch exists {
	case false:
		completePath = fmt.Sprintf("%s/*", controllerName)
		filters, exists = allowController[completePath]
		if !exists {
			logs.Error(constants.DefaultErrorTemplate, "filter.filterPrepare", "allowController", errors.ErrAllowController)
			return
		}
	}

	if len(filters) <= constants.DefaultZero {
		logs.Error(constants.DefaultErrorTemplate, "filter.filterPrepare", "allowController", errors.ErrAllowController)
		return
	}

	c.getRequestBodyParam()

	for _,filter := range filters{
		filter(c)
	}
}

func UserAuth(c *BaseController)  {

}
