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

type MoveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) MoveFileLogic {
	return MoveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFileLogic) MoveFile(req types.MoveFileReq) (resp *types.MoveFileReply, err error) {
	if _, err := l.svcCtx.ProjectRpc.MoveFile(l.ctx, &project.MoveFileReq{
		Id:      req.ID,
		Oldpath: req.Oldpath,
		Newpath: req.Newpath,
	}); err != nil {
		logx.Error("Fail to MoveFile:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.MoveFileReply{Result: true}, nil
}
