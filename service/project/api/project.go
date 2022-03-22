package main

import (
	"flag"
	"fmt"
	"net/http"

	"e5Code-Service/common/errorx"
	"e5Code-Service/service/project/api/internal/config"
	"e5Code-Service/service/project/api/internal/handler"
	"e5Code-Service/service/project/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/project-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(e error) (int, interface{}) {
		switch err := e.(type) {
		case *errorx.CodeError:
			return http.StatusOK, err.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
