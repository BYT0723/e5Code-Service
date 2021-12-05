package main

import (
	"flag"
	"fmt"

	"e5Code-Service/service/project/rpc/internal/config"
	"e5Code-Service/service/project/rpc/internal/server"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/project.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewProjectServer(ctx)
	deploy := server.NewDepolyServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		project.RegisterProjectServer(grpcServer, srv)
		project.RegisterDepolyServer(grpcServer, deploy)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
