package logic

import (
	"context"

	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteDeployLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeployLogic {
	return &DeleteDeployLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDeployLogic) DeleteDeploy(in *project.DeleteDeployReq) (*project.DeleteDeployRsp, error) {
	// todo: add your logic here and delete this line

	return &project.DeleteDeployRsp{}, nil
}
