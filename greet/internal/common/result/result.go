package result

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/internal/common/errorx"
	"net/http"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//return error
		errcode := errorx.ServerCommonError

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*errorx.CodeError); ok {
			errcode = e.GetErrCode()
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode))
	}
}

// http 参数错误
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	httpx.WriteJson(w, http.StatusBadRequest, Error(errorx.ReuqestParamError))
}
