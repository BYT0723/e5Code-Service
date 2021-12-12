package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"e5Code-Service/common"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddDeployLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeployLogic {
	return &AddDeployLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddDeployLogic) AddDeploy(in *project.AddDeployReq) (*project.AddDeployRsp, error) {
	payload := model.Deploy{
		Id:        common.GetUUID(),
		Name:      in.Name,
		ProjectId: in.ProjectID,
	}
	if in.SshConfig != nil {
		sshconfig, err := json.Marshal(in.SshConfig)
		if err != nil {
			logx.Error("Fail to Marshal SSHConfig, err: ", err.Error())
			return nil, err
		}
		payload.SshConfig = sql.NullString{
			String: string(sshconfig),
			Valid:  true,
		}
	}
	if in.ContainerConfig != nil {
		containerConfig, err := json.Marshal(in.ContainerConfig)
		if err != nil {
			logx.Error("Fail to Marshal ContainerConfig, err: ", err.Error())
			return nil, err
		}
		payload.ContainerConfig = sql.NullString{
			String: string(containerConfig),
			Valid:  true,
		}
	}
	if _, err := l.svcCtx.DeployModel.Insert(&payload); err != nil {
		logx.Error("Fail to insert deploy, err: ", err.Error())
		return nil, err
	}
	return &project.AddDeployRsp{
		Result: &project.Deploy{
			Id:              payload.Id,
			Name:            in.Name,
			ProjectID:       in.ProjectID,
			SshConfig:       in.SshConfig,
			ContainerConfig: in.ContainerConfig,
		},
	}, nil
}
