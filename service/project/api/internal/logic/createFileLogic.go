package logic

import (
	"context"

	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateFileLogic {
	return CreateFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFileLogic) CreateFile(req types.CreateFileReq) (resp *types.CreateFileReply, err error) {
	if _, err := l.svcCtx.ProjectRpc.CreateFile(l.ctx, &project.CreateFileReq{Id: req.ID, Path: req.Path}); err != nil {
		logx.Error("Fail to CreateFile:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.CreateFileReply{Result: true}, nil
}
