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

type DeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteFileLogic {
	return DeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFileLogic) DeleteFile(req types.DeleteFileReq) (resp *types.DeleteFileReply, err error) {
	if _, err := l.svcCtx.ProjectRpc.DeleteFile(l.ctx, &project.DeleteFileReq{
		Id:   req.ID,
		Path: req.Path,
	}); err != nil {
		logx.Error("Fail to DeleteFile:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.DeleteFileReply{Result: true}, nil
}
