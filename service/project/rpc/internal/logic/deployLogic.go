package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"e5Code-Service/common"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
	"golang.org/x/crypto/ssh"
)

type DeployLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeployLogic {
	return &DeployLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeployLogic) Deploy(in *project.DeployReq) (*project.DeployRsp, error) {
	deploy, err := l.svcCtx.DeployModel.FindOne(in.Id)
	if err != nil {
		logx.Error("Fail to find one Deploy, err: ", err.Error())
		return nil, err
	}
	var sshConfig project.SSHConfig
	if err := json.Unmarshal([]byte(deploy.SshConfig.String), &sshConfig); err != nil {
		logx.Error("Fail to unmarshal deploy, err: ", err.Error())
		return nil, err
	}
	client, err := l.getSSHClient(&sshConfig)
	if err != nil {
		logx.Error("Fail to get sshClient, err: ", err.Error())
		return nil, err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		logx.Error("Fail to get sshSession, err: ", err.Error())
		return nil, err
	}
	defer session.Close()
	pro, err := l.svcCtx.ProjectModel.FindOne(deploy.ProjectId)
	if err != nil {
		logx.Error("Fail to find one project, err: ", err.Error())
		return nil, err
	}
	cmd := fmt.Sprintf("cd; git clone %s&", pro.Url)
	if output, err := l.exec(session, cmd); err != nil {
		logx.Errorf("Fail to exec cmd( %s ), err: ", cmd, err.Error())
		return nil, err
	}
	return &project.DeployRsp{Result: true}, nil
}

func (l *DeployLogic) exec(session *ssh.Session, cmd string) (string, error) {
	result, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}
	return string(result), nil

}

func (l *DeployLogic) getSSHClient(config *project.SSHConfig) (*ssh.Client, error) {
	c := &ssh.ClientConfig{
		Timeout:         time.Second * 3,
		User:            config.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if config.SshType == common.PassWord {
		c.Auth = []ssh.AuthMethod{ssh.Password(config.Password)}
	} else if config.SshType == common.SSHKey {
		publicKey, err := ssh.ParsePrivateKey([]byte(config.SshKey))
		if err != nil {
			logx.Error("Fail to parse ssh public key, err: ", err.Error())
			return nil, err
		}
		c.Auth = []ssh.AuthMethod{ssh.PublicKeys(publicKey)}
	}
	return ssh.Dial("tcp", config.Host, c)
}
