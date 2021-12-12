package logic

import (
	"context"
	"encoding/json"

	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDepolyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepolyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepolyLogic {
	return &GetDepolyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepolyLogic) GetDepoly(in *project.GetDeployReq) (*project.GetDeployRsp, error) {
	deploy, err := l.svcCtx.DeployModel.FindOne(in.Id)
	if err != nil {
		logx.Error("Fail to findone deploy, err: ", err.Error())
		return nil, err
	}
	sshConfig := project.SSHConfig{}
	if deploy.SshConfig.Valid {
		json.Unmarshal([]byte(deploy.SshConfig.String), sshConfig)
	}
	containerConfig := project.ContainerConfig{}
	if deploy.ContainerConfig.Valid {
		json.Unmarshal([]byte(deploy.ContainerConfig.String), containerConfig)
	}
	return &project.GetDeployRsp{
		Result: &project.Deploy{
			Id:              deploy.Id,
			Name:            deploy.Name,
			ProjectID:       deploy.ProjectId,
			SshConfig:       &sshConfig,
			ContainerConfig: &containerConfig,
		},
	}, nil
}
