package svc

import (
	"e5Code-Service/common/dockerx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/config"
	"e5Code-Service/service/project/rpc/project"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config       config.Config
	UserRpc      user.User
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
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logx.Error("Fail to New Gorm DB:", err.Error())
		panic("DB Uninitialized")
	}
	db.AutoMigrate(
		&model.BuildPlan{},
		&model.Image{},
	)
	return &ServiceContext{
		Config:       c,
		UserRpc:      user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProjectRpc:   project.NewProject(zrpc.MustNewClient(c.ProjectRpc)),
		DockerClient: dockerCli,
		DB:           db,
	}
}
