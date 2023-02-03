package result

import (
	"github.com/beego/i18n"
	"greet/internal/common/errorx"
	"greet/internal/middleware"
)

type ResponseSuccess struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *ResponseSuccess {
	return &ResponseSuccess{200, "OK", data}
}

type ResponseError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Error(errCode uint32) *ResponseError {
	return &ResponseError{errCode, i18n.Tr(middleware.GetLang(), errorx.MapErrMsg(errCode))}
}
