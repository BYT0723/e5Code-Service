package svc

import (
	"e5Code-Service/service/user/api/internal/config"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/userClient"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	UserRpc   userClient.UserServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		UserRpc:   userClient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
