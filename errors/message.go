package errors

type ErrMsg struct{
	Code int64	`json:"code"`
	Msg  string `json:"msg"`
}

var (
	ErrSysBusy = &ErrMsg{Code:1000001, Msg:"系统繁忙,请稍后再试~"}
	ErrParam = &ErrMsg{Code:1000002, Msg:"参数错误~"}
	ErrQueryError = &ErrMsg{Code:1000003, Msg:"查询出错了~"}
	ErrQueryRecordNotExists = &ErrMsg{Code:1000003, Msg:"查询记录不存在~"}
)