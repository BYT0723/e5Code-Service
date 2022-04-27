package svc

import (
	"e5Code-Service/common/gitx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	GitCli *gitx.Cli
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logx.Error("Fail to Connect DB:", err.Error())
	}
	db.AutoMigrate(
		&model.User{},
		&model.Permission{},
	)
	return &ServiceContext{
		Config: c,
		Db:     db,
		GitCli: gitx.NewCli(),
	}
}
