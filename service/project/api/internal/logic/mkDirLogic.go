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

type MkDirLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMkDirLogic(ctx context.Context, svcCtx *svc.ServiceContext) MkDirLogic {
	return MkDirLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MkDirLogic) MkDir(req types.MkDirReq) (resp *types.MkDirReply, err error) {
	if _, err := l.svcCtx.ProjectRpc.MkDir(l.ctx, &project.MkDirReq{Id: req.ID, Path: req.Path}); err != nil {
		logx.Error("Fail to mkDir: ", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.MkDirReply{Result: true}, nil
}
