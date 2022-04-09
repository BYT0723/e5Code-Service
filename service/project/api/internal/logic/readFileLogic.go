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

type ReadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReadFileLogic {
	return ReadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadFileLogic) ReadFile(req types.ReadFileReq) (resp *types.ReadFileReply, err error) {
	rsp, err := l.svcCtx.ProjectRpc.ReadFile(l.ctx, &project.ReadFileReq{
		Id:     req.ID,
		Path:   req.Path,
		IsWork: req.IsWork,
	})
	if err != nil {
		logx.Error("Fail to ReadFile:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}

	return &types.ReadFileReply{
		Body: rsp.Body,
	}, nil
}
