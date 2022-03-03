package logic

import (
	"context"

	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeployLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDeployLogic {
	return GetDeployLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeployLogic) GetDeploy(req types.GetDeployReq) (resp *types.GetDeployRsp, err error) {
	rsp, err := l.svcCtx.DeployServer.GetDepoly(l.ctx, &project.GetDeployReq{Id: req.ID})
	if err != nil {
		logx.Error("Fail to get deploy, err: ", err.Error())
		return nil, err
	}
	return &types.GetDeployRsp{
		Result: types.Deploy{
			ID:        rsp.Result.Id,
			Name:      rsp.Result.Name,
			ProjectID: rsp.Result.ProjectID,
			SSHConfig: types.SSHConfig{
				Host:     rsp.Result.SshConfig.Host,
				User:     rsp.Result.SshConfig.User,
				SSHType:  rsp.Result.SshConfig.SshType,
				Password: rsp.Result.SshConfig.Password,
				SSHKey:   rsp.Result.SshConfig.SshKey,
			},
			ContainerConfig: types.ContainerConfig{
				Name:         rsp.Result.ContainerConfig.Name,
				NetworkType:  rsp.Result.ContainerConfig.NetworkType,
				IP:           rsp.Result.ContainerConfig.Ip,
				Ports:        rsp.Result.ContainerConfig.Ports,
				Environments: rsp.Result.ContainerConfig.Environments,
			},
		},
	}, nil
}
