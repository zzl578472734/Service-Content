package vars

type DefaultListParam struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type DefaultIdQueryParam struct {
	Id int64 `json:"id"`
}

type DefaultStatusQueryParam struct {
	Status int8 `json:"status"`
}
