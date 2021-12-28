package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Host string
		Pass string
		DB   int
	}
	CacheRedis cache.CacheConf
	Token      struct {
		JwtKey string
		Expire int64
	}
}
