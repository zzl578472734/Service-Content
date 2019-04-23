package errors

type ErrMsg struct{
	Code int64	`json:"code"`
	Msg  string `json:"msg"`
}

var (
	ErrSysBusy = &ErrMsg{Code:1000001, Msg:"系统繁忙,请稍后再试~"}
	ErrUrlError = &ErrMsg{Code:1000001, Msg:"您访问的链接未找到~"}

	ErrParam = &ErrMsg{Code:2000002, Msg:"参数错误~"}
	ErrQueryError = &ErrMsg{Code:2000003, Msg:"查询出错了~"}
	ErrQueryRecordNotExists = &ErrMsg{Code:2000003, Msg:"查询记录不存在~"}
	ErrInsertError = &ErrMsg{Code:2000003, Msg:"數據添加出错了~"}

	ErrStatusMap = &ErrMsg{Code:2000001, Msg:"状态不合法!~"}
)