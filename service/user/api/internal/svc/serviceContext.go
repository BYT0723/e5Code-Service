package svc

import (
	"e5Code-Service/service/user/api/internal/config"
	"e5Code-Service/service/user/api/internal/middleware"
	"e5Code-Service/service/user/rpc/user"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
	Redis   *redis.Client
	LoadValue  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		LoadValue:  middleware.NewLoadValueMiddleware().Handle,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Redis: redis.NewClient(&redis.Options{
			Addr: c.Redis.Host,
			DB:   c.Redis.DB,
		}),
	}
}
