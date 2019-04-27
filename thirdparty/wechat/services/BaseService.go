package services

import (
	constants2 "Service-Content/constants"
	"Service-Content/thirdparty/wechat/constants"
	"Service-Content/utils/redisclient"
	"fmt"
	"github.com/astaxie/beego"
	"Content-Admin/utils"
	"encoding/json"
	"Service-Content/thirdparty/wechat/vars"
	"github.com/astaxie/beego/logs"
)

type BaseService struct {
}

var (
	appId     = beego.AppConfig.String("wechat.appId")
	appSecret = beego.AppConfig.String("wechat.appSecret")
)

func (s *BaseService) Request(url string, data interface{}) {

}

func (s *BaseService) getAccessToken() string{

	cacheKey := constants.WechatAccessToken
	var accessToken string
	err := redisclient.GetCache(cacheKey, accessToken)
	if err == nil && accessToken != constants2.DefaultEmptyString {
		return accessToken
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	option := &utils.HttpRequest{
		Url: url,
	}
	resp, err := utils.HttpGet(option)
	if err != nil{
		return constants2.DefaultEmptyString
	}
	WXAccessToken := new(vars.WXAccessTokenReturn)
	err = json.Unmarshal(resp.Body, WXAccessToken)
	if err != nil{
		logs.Info(constants2.DefaultErrorTemplate, "BaseService.getAccessToken", "json.Unmarshal", err)
		return constants2.DefaultEmptyString
	}

	redisclient.SetCache(cacheKey, WXAccessToken.ExpireIn)

	return WXAccessToken.AccessToken
}
