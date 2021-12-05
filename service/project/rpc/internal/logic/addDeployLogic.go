package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &project.AddDeployRsp{}, nil
}
