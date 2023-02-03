package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CacheRedis cache.CacheConf
	Mysql      struct {
		DataSource string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Lang struct {
		Names string
		Types string
		Path  string
	}
	UserRpc zrpc.RpcClientConf
}
