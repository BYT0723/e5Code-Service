package logic

import (
	"context"

	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeployLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateDeployLogic {
	return UpdateDeployLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDeployLogic) UpdateDeploy(req types.UpdateDeployReq) (resp *types.UpdateDeployRsp, err error) {
	payload := project.Deploy{
		Id: req.Payload.ID,
		SshConfig: &project.SSHConfig{
			Host:     req.Payload.SSHConfig.Host,
			User:     req.Payload.SSHConfig.User,
			SshType:  req.Payload.SSHConfig.SSHType,
			SshKey:   req.Payload.SSHConfig.SSHKey,
			Password: req.Payload.SSHConfig.Password,
		},
		ContainerConfig: &project.ContainerConfig{
			Name:         req.Payload.ContainerConfig.Name,
			NetworkType:  req.Payload.ContainerConfig.NetworkType,
			Ip:           req.Payload.ContainerConfig.IP,
			Ports:        req.Payload.ContainerConfig.Ports,
			Environments: req.Payload.ContainerConfig.Environments,
		},
	}
	if req.Payload.Name != "" {
		payload.Name = req.Payload.Name
	}
	if req.Payload.ProjectID != "" {
		payload.ProjectID = req.Payload.ProjectID
	}
	rsp, err := l.svcCtx.DeployServer.UpdateDeploy(l.ctx, &project.UpdateDeployReq{Payload: &payload})
	if err != nil {
		logx.Error("Fail to update deploy, err: ", err.Error())
		return nil, err
	}
	return &types.UpdateDeployRsp{Result: rsp.Result}, nil
}
