package svc

import (
	"github.com/cjstarcc/zeroframework/user/rpc/user/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"greet/internal/config"
	"greet/internal/middleware"
	"greet/model"
)

type ServiceContext struct {
	Config    config.Config
	I18n      rest.Middleware
	UserModel model.UsersModel
	UserRpc   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		UserModel: model.NewUsersModel(sqlConn, c.CacheRedis),
		Config:    c,
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		I18n:      middleware.NewI18nMiddleware(c).Handle,
	}
}
