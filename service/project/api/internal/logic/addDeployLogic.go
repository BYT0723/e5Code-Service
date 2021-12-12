package logic

import (
	"context"
	"fmt"

	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddDeployLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddDeployLogic {
	return AddDeployLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDeployLogic) AddDeploy(req types.AddDeployReq) (resp *types.AddDeployRsp, err error) {
	fmt.Println("add deploy ....")
	payload := project.AddDeployReq{
		Name:      req.Name,
		ProjectID: req.ProjectID,
		SshConfig: &project.SSHConfig{
			Host:     req.SSHConfig.Host,
			User:     req.SSHConfig.User,
			SshType:  req.SSHConfig.SSHType,
			Password: req.SSHConfig.Password,
			SshKey:   req.SSHConfig.SSHKey,
		},
		ContainerConfig: &project.ContainerConfig{
			Name:         req.ContainerConfig.Name,
			NetworkType:  req.ContainerConfig.NetworkType,
			Ip:           req.ContainerConfig.IP,
			Ports:        req.ContainerConfig.Ports,
			Environments: req.ContainerConfig.Environments,
		},
	}
	rsp, err := l.svcCtx.DeployServer.AddDeploy(l.ctx, &payload)
	if err != nil {
		logx.Error("Fail to add deploy, err: ", err.Error())
		return nil, err
	}
	return &types.AddDeployRsp{
		Result: types.Deploy{
			ID:              rsp.Result.Id,
			Name:            rsp.Result.Name,
			ProjectID:       rsp.Result.ProjectID,
			SSHConfig:       req.SSHConfig,
			ContainerConfig: req.ContainerConfig,
		},
	}, nil
}
