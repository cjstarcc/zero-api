package logic

import (
	"context"
	"github.com/cjstarcc/zeroframework/user/rpc/user/userclient"
	"github.com/golang-jwt/jwt/v4"
	"net/http"

	"time"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	var loginReq types.LoginReq
	loginReq.Password = req.Password
	loginReq.Username = req.Username

	userResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.SCLoginReq{Password: req.Password, Username: req.Username})

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, userResp.GetId())

	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
