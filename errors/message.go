package errors

type ErrMsg struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

var (
	ErrSysBusy  = &ErrMsg{Code: 1000001, Msg: "系统繁忙,请稍后再试~"}
	ErrUrlError = &ErrMsg{Code: 1000001, Msg: "您访问的链接未找到~"}

	ErrParam          = &ErrMsg{Code: 1000002, Msg: "参数错误~"}
	ErrRequestTimeout = &ErrMsg{Code: 1000003, Msg: "请求超时~"}
	ErrInvalidRequest = &ErrMsg{Code: 1000004, Msg: "无效的请求~"}
	ErrSignature      = &ErrMsg{Code: 1000005, Msg: "签名错误~"}

	ErrQueryError           = &ErrMsg{Code: 2000005, Msg: "查询出错了~"}
	ErrQueryRecordNotExists = &ErrMsg{Code: 2000006, Msg: "查询记录不存在~"}
	ErrInsertError          = &ErrMsg{Code: 2000007, Msg: "數據添加出错了~"}
	ErrStatusMap            = &ErrMsg{Code: 2000001, Msg: "状态不合法!~"}

	ErrUserNotExist            = &ErrMsg{Code: 3000001, Msg: "用户不存在~"}
	ErrUserIsExist             = &ErrMsg{Code: 3000001, Msg: "用户已经存在~"}
	ErrUsernameLength          = &ErrMsg{Code: 3000001, Msg: "请输入合法的用户姓名~"}
	ErrPasswordLength          = &ErrMsg{Code: 3000001, Msg: "请输入6~20位的用户密码~"}
	ErrInputUsernameOrPassword = &ErrMsg{Code: 3000001, Msg: "用户名或者密码错误~"}
	ErrPasswordErr             = &ErrMsg{Code: 3000001, Msg: "请输入6~用户名或密码错误~"}
)
