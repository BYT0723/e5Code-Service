package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateDeployLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeployLogic {
	return &UpdateDeployLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDeployLogic) UpdateDeploy(in *project.UpdateDeployReq) (*project.UpdateDeployRsp, error) {
	payload := model.Deploy{Id: in.Payload.Id}
	if in.Payload.Name != "" {
		payload.Name = in.Payload.Name
	}
	if in.Payload.ProjectID != "" {
		payload.ProjectId = in.Payload.ProjectID
	}
	if in.Payload.SshConfig != nil {
		sshConfig, err := json.Marshal(in.Payload.SshConfig)
		if err != nil {
			logx.Error("Fail to Marshal sshconfig, err: ", err.Error())
			return nil, err
		}
		payload.SshConfig = sql.NullString{
			String: string(sshConfig),
			Valid:  true,
		}
	}
	if in.Payload.ContainerConfig != nil {
		containerConfig, err := json.Marshal(in.Payload.ContainerConfig)
		if err != nil {
			logx.Error("Fail to Marshal ContainerConfig, err: ", err.Error())
			return nil, err
		}
		payload.ContainerConfig = sql.NullString{
			String: string(containerConfig),
			Valid:  true,
		}
	}
	if err := l.svcCtx.DeployModel.Update(&payload); err != nil {
		logx.Error("Fail to update deploy, err: ", err.Error())
		return &project.UpdateDeployRsp{Result: false}, err
	}

	return &project.UpdateDeployRsp{Result: true}, nil
}
