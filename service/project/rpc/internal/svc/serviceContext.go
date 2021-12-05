package svc

import (
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/config"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	ProjectModel   model.ProjectModel
	DeployModel    model.DeployModel
	ContainerModel model.ContainerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		ProjectModel:   model.NewProjectModel(conn, c.CacheRedis),
		DeployModel:    model.NewDeployModel(conn, c.CacheRedis),
		ContainerModel: model.NewContainerModel(conn, c.CacheRedis),
	}
}
