package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	RepositoryConf struct {
		Repositories string
		Tars         string
		BuildLogs    string
	}
	ProjectRpc zrpc.RpcClientConf
}
