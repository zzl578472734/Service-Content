package vars

type DefaultListParam struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type DefaultQueryParam struct {
	Id int64 `json:"id"`
}
