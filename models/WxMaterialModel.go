package models

import (
	constants2 "Service-Content/constants"
	"Service-Content/thirdparty/wechat/constants"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type WxMaterialModel struct {
	Id                 int64  `json:"id" orm:"column(id)"`
	MediaId            string `json:"media_id" orm:"column(media_id)"`
	Title              string `json:"title" orm:"column(title)"`
	ThumbMediaId       string `json:"thumb_media_id" orm:"column(thumb_media_id)"`
	ShowCoverPic       int8   `json:"show_cover_pic" orm:"column(show_cover_pic)"`
	Author             string `json:"author" orm:"column(author)"`
	Digest             string `json:"digest" orm:"column(digest)"`
	Content            string `json:"content" orm:"column(content)"`
	Url                string `json:"url" orm:"column(url)"`
	ThumbUrl           string `json:"thumb_url" orm:"column(thumb_url)"`
	ContentSourceUrl   string `json:"content_source_url" orm:"column(content_source_url)"`
	NeedOpenComment    int8   `json:"need_open_comment" orm:"column(need_open_comment)"`
	OnlyFansCanComment int8   `json:"only_fans_can_comment" orm:"column(only_fans_can_comment)"`
	SingleCreateTime   int    `json:"single_create_time" orm:"column(single_create_time)"`
	SingleUpdateTime   int    `json:"single_update_time" orm:"column(single_update_time)"`
	CreateTime         int    `json:"create_time" orm:"column(create_time)"`
	UpdateTime         int    `json:"update_time" orm:"column(update_time)"`
	BaseModel
	O orm.Ormer `json:"-"`
}

func NewWxMaterialModel() *WxMaterialModel {
	return &WxMaterialModel{O: orm.NewOrm()}
}

func (m *WxMaterialModel) TableName() string {
	return TableName(constants.MaterialTableNmae)
}

func (m *WxMaterialModel) Insert(param *WxMaterialModel) (int64, error) {
	id, err := m.O.Insert(param)
	if ormErr(err) != nil {
		logs.Error(constants2.DefaultErrorTemplate, "WxMaterialModel.Insert", "Insert", err)
		return constants2.DefaultZero, err
	}
	return id, nil
}

func (m *WxMaterialModel) List(param map[string]interface{}, page, pageSize int) ([]*WxMaterialModel, int64, error) {
	query := m.O.QueryTable(m.TableName())

	if len(param) > constants2.DefaultZero {
		for key, value := range param {
			query = query.Filter(key, value)
		}
	}

	count, err := query.Count()
	if err != nil {
		logs.Error(constants2.DefaultErrorTemplate, "WxMaterialModel.List", "query.Count", err)
		return nil, constants2.DefaultZero, err
	}

	list := make([]*WxMaterialModel, constants2.DefaultZero)

	query = query.Limit(pageSize).Offset((page - 1) * pageSize)
	if _, err = query.All(&list); ormErr(err) != nil {
		logs.Error(constants2.DefaultErrorTemplate, "WxMaterialModel.List", "query.All", err)
		return nil, constants2.DefaultZero, err
	}
	return list, count, nil
}
