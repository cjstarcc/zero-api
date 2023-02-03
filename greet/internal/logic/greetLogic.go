package logic

import (
	"context"
	"github.com/pkg/errors"
	"greet/internal/common/errorx"
	"net/http"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return nil, errors.Wrapf(errorx.NewErrCode(errorx.ReuqestParamError), "just for test !")
}
