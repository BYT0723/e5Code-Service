package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	ProjectRpc   zrpc.RpcClientConf
	RegistryConf struct {
		Local    string
		Tar      string
		BuildLog string
	}
}
