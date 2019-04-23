package controllers

import (
	"Service-Content/constants"
	"Service-Content/errors"
	"crypto/sha512"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
	"time"
)

var (
	allowController = map[string][]funcFilter{
		"UserController/*": {AppAuth},
	}
	platformConf, accessConf, permissionConf config.Configer
)

type funcFilter func(c *BaseController)

func init() {
	// 是否可以采用分发配置的方式
	var err error
	platformConf, err = config.NewConfig("json", "conf/platform.json")
	if err != nil {
		logs.Error(constants.DefaultErrorTemplate, "filter.init", "config.NewConfig", err)
		panic(err)
	}

	accessConf, err = config.NewConfig("json", "conf/access.json")
	if err != nil {
		logs.Error(constants.DefaultErrorTemplate, "filter.init", "config.NewConfig", err)
		panic(err)
	}

	permissionConf, err = config.NewConfig("json", "conf/permission.json")
	if err != nil {
		logs.Error(constants.DefaultErrorTemplate, "filter.init", "config.NewConfig", err)
		panic(err)
	}
}

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

	for _, filter := range filters {
		filter(c)
	}
}

func AppAuth(c *BaseController) {
	request := c.Ctx.Input.GetData(ApiRequestBody).(*ApiRequestParam)

	clientLocalTimestamps := time.Unix(request.Timestamps, constants.DefaultZero)
	serviceLocalTimestamps := time.Now()
	if serviceLocalTimestamps.Sub(clientLocalTimestamps) > constants.DefaultRequestMaxTimestamps ||
		serviceLocalTimestamps.Sub(clientLocalTimestamps) < constants.DefaultRequestMinTimestamps {
		c.ApiErrorReturn(errors.ErrRequestTimeout)
		return
	}

	if request.AppId == constants.DefaultEmptyString ||
		request.AppToken == constants.DefaultEmptyString {
		c.ApiErrorReturn(errors.ErrParam)
		return
	}

	key := fmt.Sprintf("%s::%s::appSecret", beego.AppConfig.String("runmode"), request.AppId)
	appSecret := platformConf.String(key)
	if appSecret == constants.DefaultEmptyString {
		c.ApiErrorReturn(errors.ErrInvalidRequest)
		return
	}

	signStr := fmt.Sprintf("app_id=%s&app_secret=%s&timestamps=%d&nonce=%s", request.AppId, appSecret, request.Timestamps, request.Nonce)

	if sign := fmt.Sprintf("%x", sha512.Sum512([]byte(signStr))); sign != request.AppToken {
		c.ApiErrorReturn(errors.ErrSignature)
		return
	}

	// 检验接口是否有访问的权限
	checkPermission(c)
}

func checkPermission(c *BaseController) {
	request := c.Ctx.Input.GetData(ApiRequestBody).(*ApiRequestParam)

	key := fmt.Sprintf("%s::%s::id", beego.AppConfig.String("runmode"), request.AppId)
	id, err := platformConf.Int(key)
	if err != nil {
		logs.Error(constants.DefaultEmptyString, "filter.checkPermission", "platformConf.Int", err)
		c.ApiErrorReturn(errors.ErrSysBusy)
		return
	}

	// 获取平台所有可以访问的接口
	api := permissionConf.String(fmt.Sprintf("%d::api", id))
	if api == constants.DefaultEmptyString {
		logs.Error(constants.DefaultEmptyString, "filter.checkPermission", "permissionConf.String", errors.ErrAccessUrl)
		c.ApiErrorReturn(errors.ErrSysBusy)
		return
	}

	controllerName, actionName := c.GetControllerAndAction()
	completeUrl := fmt.Sprintf("%s/%s", controllerName, actionName)

	access, err := accessConf.Int(completeUrl)
	if err != nil {
		logs.Error(constants.DefaultEmptyString, "filter.checkPermission", "accessConf.Int", err)
		c.ApiErrorReturn(errors.ErrSysBusy)
		return
	}

	// 验证当前的appId是否有访问当前接口的权限
	exists := strings.Contains(api, strconv.Itoa(access))
	if !exists {
		c.ApiErrorReturn(errors.ErrInvalidRequest)
		return
	}
	return
}
