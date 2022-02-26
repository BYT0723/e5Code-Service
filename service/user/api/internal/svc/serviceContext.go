package svc

import (
	"e5Code-Service/service/user/api/internal/config"
	"e5Code-Service/service/user/rpc/userclient"

	"github.com/go-redis/redis"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.UserServer
	Redis   *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
			DB:       int(c.Redis.DB),
		}),
	}
}
