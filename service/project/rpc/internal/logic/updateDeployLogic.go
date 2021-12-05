package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &project.UpdateDeployRsp{}, nil
}
