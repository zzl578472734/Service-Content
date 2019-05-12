package services

import (
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"Service-Content/thirdparty/wechat/vars"
	"github.com/astaxie/beego/logs"
	"Service-Content/utils"
	"Service-Content/thirdparty/wechat/constants"
	"Service-Content/utils/redisclient"
)

type BaseService struct {
}

var (
	appId     = beego.AppConfig.String("wechat.appId")
	appSecret = beego.AppConfig.String("wechat.appSecret")
)

func getAccessToken() string{

	var accessToken string
	err := redisclient.GetCache(constants.WechatAccessToken, accessToken)
	if err == nil && accessToken != constants.DefaultEmptyString {
		return accessToken
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	option := &utils.HttpRequest{
		Url: url,
	}
	resp, err := utils.HttpGet(option)
	if err != nil{
		return constants.DefaultEmptyString
	}
	WXAccessToken := new(vars.WXAccessTokenReturn)
	err = json.Unmarshal(resp.Body, WXAccessToken)
	if err != nil{
		logs.Error(constants.DefaultErrorTemplate, "BaseService.getAccessToken", "json.Unmarshal", err)
		return constants.DefaultEmptyString
	}
	if WXAccessToken != nil && WXAccessToken.Errcode > constants.DefaultEmptyZero {
		// 出现错误
		logs.Error(constants.DefaultErrorTemplate, "BaseService.getAccessToken", "json.Unmarshal", string(resp.Body))
		return constants.DefaultEmptyString
	}

	redisclient.SetCache(constants.WechatAccessToken, WXAccessToken.ExpireIn)

	return WXAccessToken.AccessToken
}

