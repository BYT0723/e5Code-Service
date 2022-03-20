package svc

import (
	"e5Code-Service/common/gitx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/config"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	UserRpc user.User
	GitCli  *gitx.Cli
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	db.AutoMigrate(
		&model.Project{},
	)
	if err != nil {
		logx.Error("Fail to Open DB: ", err.Error())
		return nil
	}
	return &ServiceContext{
		Config:  c,
		DB:      db,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		GitCli:  gitx.NewCli(),
	}
}
