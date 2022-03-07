package svc

import (
	"e5Code-Service/common/gitx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/config"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	ProjectModel model.ProjectModel
	UserRpc      user.User
	GitCli       *gitx.Cli
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProjectModel: model.NewProjectModel(conn, c.CacheRedis),
		UserRpc:      user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		GitCli:       gitx.NewCli(),
	}
}
