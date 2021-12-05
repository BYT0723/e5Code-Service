package logic

import (
	"context"

	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteProjectLogic {
	return DeleteProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProjectLogic) DeleteProject(req types.DeleteProjectReq) (*types.DeleteProjectReply, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteProjectReply{}, nil
}
