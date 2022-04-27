package main

import (
	"flag"
	"fmt"
	"net/http"

	"e5Code-Service/service/ci/api/internal/config"
	"e5Code-Service/service/ci/api/internal/handler"
	"e5Code-Service/service/ci/api/internal/hub"
	"e5Code-Service/service/ci/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/ci-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ci/build",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			hub.ServeWs(ctx.Hub, w, r)
		},
	})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
