package svc

import (
	"e5Code-Service/service/project/api/internal/config"
	"e5Code-Service/service/project/rpc/deploy"
	"e5Code-Service/service/project/rpc/projectClient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	ProjectServer projectClient.ProjectServer
	DeployServer  deploy.DeployServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ProjectServer: projectClient.NewProject(zrpc.MustNewClient(c.ProjectRpc)),
		DeployServer:  deploy.NewDeploy(zrpc.MustNewClient(c.ProjectRpc)),
	}
}
