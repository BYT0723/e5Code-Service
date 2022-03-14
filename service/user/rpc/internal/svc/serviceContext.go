package svc

import (
	"e5Code-Service/common/gitx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	UserModel       model.UserModel
	PermissionModel model.PermissionModel
	GitCli          *gitx.Cli
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	logx.MustSetup(c.Log)
	return &ServiceContext{
		Config:          c,
		UserModel:       model.NewUserModel(conn, c.CacheRedis),
		PermissionModel: model.NewPermissionModel(conn, c.CacheRedis),
		GitCli:          gitx.NewCli(),
	}
}
