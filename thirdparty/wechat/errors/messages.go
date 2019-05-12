package errors

type ErrMsg struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

var (
	ErrSysBusy    = &ErrMsg{Code: 1000001, Msg: "系统繁忙,请稍后再试~"}
	ErrWechatQueryParam = &ErrMsg{Code: 1000001, Msg: "参数错误~"}

	ErrMaterialType   = &ErrMsg{Code: 1000001, Msg: "素材管理类型不合法"}
	ErrMaterialOffset = &ErrMsg{Code: 1000001, Msg: "请输入大于等于0的下标位置"}
	ErrMaterialCount  = &ErrMsg{Code: 1000001, Msg: "请输入1~25之间的数字~"}
)
