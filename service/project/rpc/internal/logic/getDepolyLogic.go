package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &project.GetDeployRsp{}, nil
}
