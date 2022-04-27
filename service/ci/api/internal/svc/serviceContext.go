package svc

import (
	"e5Code-Service/service/ci/api/internal/config"
	"e5Code-Service/service/ci/api/internal/hub"
	"e5Code-Service/service/ci/api/internal/middleware"
	"e5Code-Service/service/ci/rpc/ci"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	LoadValue rest.Middleware
	CIRpc     ci.Ci
	Hub       *hub.Hub
}

func NewServiceContext(c config.Config) *ServiceContext {
	ciRpc := ci.NewCi(zrpc.MustNewClient(c.CIRpc))
	hub := hub.NewHub(ciRpc)
	go hub.Run()
	return &ServiceContext{
		Config:    c,
		LoadValue: middleware.NewLoadValueMiddleware().Handle,
		CIRpc:     ciRpc,
		Hub:       hub,
	}
}
