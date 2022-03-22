package svc

import (
	"e5Code-Service/common/dockerx"
	"e5Code-Service/service/ci/rpc/internal/config"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	ProjectRpc   project.Project
	DockerClient *dockerx.DockerClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	dockerCli, err := dockerx.NewDockerClient()
	if err != nil {
		logx.Error("Fail to New DockerClient: ", err.Error())
		panic("DockerClient Uninitialized")
	}
	return &ServiceContext{
		Config:       c,
		ProjectRpc:   project.NewProject(zrpc.MustNewClient(c.ProjectRpc)),
		DockerClient: dockerCli,
	}
}
