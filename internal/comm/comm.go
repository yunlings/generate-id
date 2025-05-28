package comm

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type ResCode uint

const (
	SuccessCode              = "200"
	FailCode                 = "600"
	CodeUserNotExist         = "10404"
	CodeUserExist            = "10405"
	CodeSuccess      ResCode = 1000 + iota
	CodeInvalidParam
	//CodeUserExist
	//CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeDateError
	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "请求成功",
	CodeInvalidParam: "请求参数错误",
	//CodeUserExist:    "用户名已存在",
	//CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeDateError:       "数据查询错误",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效Token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

type ResponseData struct {
	Code string      `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResFail(ctx *app.RequestContext, code string, msg string) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResSuccess(ctx *app.RequestContext, code string, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
