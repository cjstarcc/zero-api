package errorx

import (
	"fmt"
	"github.com/beego/i18n"
	"greet/internal/middleware"
)

type CodeError struct {
	errCode uint32
	errMsg  string
}

// errorCode
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// errMsg
func (e *CodeError) GetErrMsg() string {
	return i18n.Tr(middleware.GetLang(), MapErrMsg(e.errCode))
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, i18n.Tr(middleware.GetLang(), MapErrMsg(e.errCode)))
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: i18n.Tr(middleware.GetLang(), MapErrMsg(errCode))}
}
