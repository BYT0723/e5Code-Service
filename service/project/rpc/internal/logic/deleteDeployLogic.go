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
	if err := l.svcCtx.DeployModel.Delete(in.Id); err != nil {
		logx.Error("Fail to delete deploy, err: ", err.Error())
		return &project.DeleteDeployRsp{Result: false}, err
	}
	return &project.DeleteDeployRsp{Result: true}, nil
}
