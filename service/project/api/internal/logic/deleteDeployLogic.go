package logic

import (
	"context"

	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteDeployLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteDeployLogic {
	return DeleteDeployLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDeployLogic) DeleteDeploy(req types.DeleteDeployReq) (resp *types.DeleteDeployRsp, err error) {
	if rsp, _ := l.svcCtx.DeployServer.DeleteDeploy(l.ctx, &project.DeleteDeployReq{Id: req.ID}); rsp.Result {
		logx.Error("Fail to delete deploy, err: ", err.Error())
		return &types.DeleteDeployRsp{Result: false}, nil
	}
	return &types.DeleteDeployRsp{Result: true}, nil
}
