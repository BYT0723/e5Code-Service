package svc

import (
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/config"

	"github.com/go-redis/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	UserModel   model.UserModel
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
			DB:       c.Redis.DB,
		}),
	}
}
