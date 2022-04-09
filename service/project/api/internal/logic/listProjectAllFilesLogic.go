package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListProjectAllFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProjectAllFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListProjectAllFilesLogic {
	return ListProjectAllFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProjectAllFilesLogic) ListProjectAllFiles(req types.ListProjectAllFilesReq) (resp *types.ListProjectAllFilesReply, err error) {
	rsp, err := l.svcCtx.ProjectRpc.ListProjectAllFiles(l.ctx, &project.ListProjectAllFilesReq{Id: req.ID, IsWork: req.IsWork})
	if err != nil {
		logx.Error("Fail to ListProjectAllFiles:", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	res := make([]types.File, rsp.Count)

	copier.Copy(&res, &rsp.Result)

	return &types.ListProjectAllFilesReply{
		Count:  rsp.Count,
		Result: res,
	}, nil
}
