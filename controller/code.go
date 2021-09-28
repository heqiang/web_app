package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvaildParam
	CodeUserExist
	CodeUserNotExist
	CodeInvaildPasswordorUserName

	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:                   "success",
	CodeInvaildParam:              "请求参数错误",
	CodeUserExist:                 "用户已存在",
	CodeUserNotExist:              "用户不存在",
	CodeInvaildPasswordorUserName: "密码或用户名错误",
	CodeServerBusy:                "系统繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
