package svc

import (
	"e5Code-Service/service/project/api/internal/config"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProjectRpc project.Project
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProjectRpc: project.NewProject(zrpc.MustNewClient(c.ProjectRpc)),
	}
}
