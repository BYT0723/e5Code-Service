package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	CacheRedis     cache.CacheConf
	UserRpc        zrpc.RpcClientConf
	GitRegistryUrl struct {
		Http string
		SSH  string
	}
	RepositoryConf struct {
		Repositories string
		Tars         string
		BuildLogs    string
	}
}
