package vars

/**
 * 批量获取素材列表
 */
type BatchGetMaterialParam struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
}

/**
 * 永久素材列表接口返回
 */
type BatchGetMaterialResp struct {
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
	Item       []*GetMaterialItemResp
	WXDefaultErrReturn
}

type GetMaterialItemResp struct {
	MediaId    string `json:"media_id"`
	Content    *GetMaterialContentResp
	UpdateTime int `json:"update_time"`
}

type GetMaterialContentResp struct {
	NewsItem []*GetMaterialResp
}

type GetMaterialResp struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	ShowCoverPic     int8   `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	Url              string `json:"url"`
	ContentSourceUrl string `json:"content_source_url"`
	UpdateTime       int `json:"update_time"`
	Name             string `json:"name"`
}
