package services

import (
	"Service-Content/thirdparty/wechat/vars"
	"github.com/astaxie/beego/context"
	"Service-Content/thirdparty/wechat/constants"
	"Service-Content/thirdparty/wechat/errors"
	"Service-Content/utils"
	"github.com/astaxie/beego/logs"
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
)

type MaterialService struct {
	BaseService
	ctx *context.Context
}

var (
	NewMaterialTypeSet = map[string]interface{}{
		constants.NewMaterialTypeImage : struct {}{},
		constants.NewMaterialTypeVideo : struct {}{},
		constants.NewMaterialTypeNews : struct {}{},
	}
)

func NewMaterialService(ctx *context.Context) *MaterialService {
	return &MaterialService{ctx: ctx}
}

/**
 * 获取素材列表
 */
func (s *MaterialService) BatchGetMaterial(param *vars.BatchGetMaterialParam) *errors.ErrMsg {
	errMsg := s.filterBatchGetMaterialParam(param)
	if errMsg != nil{
		return errMsg
	}

	accessToken := getAccessToken()
	if accessToken == constants.DefaultEmptyString{
		return errors.ErrSysBusy
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s", accessToken)

	body,err := json.Marshal(param)
	if err != nil{
		logs.Error(constants.DefaultErrorTemplate, "MaterialService.BatchGetMaterial", "json.Marshal", err)
		return errors.ErrSysBusy
	}
	option := &utils.HttpRequest{
		Url: url,
		Body: body,
	}

	batchGetMaterialResp := new(vars.BatchGetMaterialResp)
	response,err := utils.HttpPost(option)
	if err != nil{
		logs.Error(constants.DefaultErrorTemplate, "MaterialService.BatchGetMaterial", "utils.HttpPost", err)
		return errors.ErrSysBusy
	}

	if beego.AppConfig.String("runmode") != beego.PROD{
		//logs.Info(constants.DefaultInfoTemplate, "MaterialService", "BatchGetMaterial", string(response.Body))
	}

	err = json.Unmarshal(response.Body, batchGetMaterialResp)
	if err != nil{
		logs.Error(constants.DefaultErrorTemplate, "MaterialService.BatchGetMaterial", "json.Unmarshal", err)
		return errors.ErrSysBusy
	}
	// 验证是否存在微信返回的错误
	if batchGetMaterialResp != nil && batchGetMaterialResp.Errcode > constants.DefaultEmptyZero {
		
	}

	return nil
}

/**
 * 素材管理列表的参数验证
 */
func (s *MaterialService) filterBatchGetMaterialParam(param *vars.BatchGetMaterialParam) *errors.ErrMsg {
	if param == nil{
		return errors.ErrWechatQueryParam
	}
	if _,ok := NewMaterialTypeSet[param.Type]; !ok {
		return errors.ErrMaterialType
	}
	if !utils.IntMin(param.Offset, constants.NewMaterialMinOffset) {
		return errors.ErrMaterialOffset
	}
	if !utils.IntRange(param.Count,constants.NewMaterialMinCount, constants.NewMaterialMaxCount) {
		return errors.ErrMaterialCount
	}
	return nil
}