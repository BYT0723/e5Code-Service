package logic

import (
	"context"

	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeployLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeployLogic {
	return DeployLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeployLogic) Deploy(req types.DeployReq) (resp *types.DeployRsp, err error) {
	if _, err := l.svcCtx.DeployServer.Deploy(l.ctx, &project.DeployReq{Id: req.ID}); err != nil {
		logx.Error("Fail to deploy project, err: ", err.Error())
		return &types.DeployRsp{Result: false}, err
	}
	return &types.DeployRsp{Result: true}, nil
}
