package svc

import (
	"e5Code-Service/service/project/api/internal/config"
	"e5Code-Service/service/project/api/internal/middleware"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	LoadValue  rest.Middleware
	ProjectRpc project.Project
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		LoadValue:  middleware.NewLoadValueMiddleware().Handle,
		ProjectRpc: project.NewProject(zrpc.MustNewClient(c.ProjectRpc)),
	}
}
