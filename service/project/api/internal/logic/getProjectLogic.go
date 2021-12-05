package logic

import (
	"context"

	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProjectLogic {
	return GetProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectLogic) GetProject(req types.GetProjectReq) (*types.GetProjectReply, error) {
	// todo: add your logic here and delete this line

	return &types.GetProjectReply{}, nil
}
