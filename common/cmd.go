package common

import "e5Code-Service/service/project/rpc/project"

/**
在新建仓库时，如果用户选择将代码托管到本平台，
则新建仓库，并返回url
*/
func InitProject() (string, error) {
	return "", nil
}

// git拉取仓库、docker编译Dockfile
func CompileProject() error {
	return nil
}

// 将编译完成的镜像拒，run一个容器
func DeployProject() error {
	return nil
}

// 利用crypto/ssh 执行远程命令
func ExecSSH(project.SSHConfig) (string, error) {
	return "", nil
}
