package svc

import (
	"e5Code-Service/common/dockerx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/config"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config       config.Config
	ProjectRpc   project.Project
	DockerClient *dockerx.DockerClient
	DB           *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dockerCli, err := dockerx.NewDockerClient()
	if err != nil {
		logx.Error("Fail to New DockerClient: ", err.Error())
		panic("DockerClient Uninitialized")
	}
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		logx.Error("Fail to New Gorm DB:", err.Error())
		panic("DB Uninitialized")
	}
	db.AutoMigrate(
		&model.BuildPlan{},
	)
	return &ServiceContext{
		Config:       c,
		ProjectRpc:   project.NewProject(zrpc.MustNewClient(c.ProjectRpc)),
		DockerClient: dockerCli,
		DB:           db,
	}
}
