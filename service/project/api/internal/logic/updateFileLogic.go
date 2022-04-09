package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateFileLogic {
	return UpdateFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFileLogic) UpdateFile(req types.UpdateFileReq) (resp *types.UpdateFileReply, err error) {
	if _, err := l.svcCtx.ProjectRpc.UpdateFile(l.ctx, &project.UpdateFileReq{
		Id:   req.ID,
		Path: req.Path,
		Body: req.Body,
	}); err != nil {
		logx.Error("Fail to UpdateFile:", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	return &types.UpdateFileReply{
		Result: true,
	}, nil
}
