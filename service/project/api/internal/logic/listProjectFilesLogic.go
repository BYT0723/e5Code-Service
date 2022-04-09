package logic

import (
	"context"

	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListProjectFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProjectFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListProjectFilesLogic {
	return ListProjectFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProjectFilesLogic) ListProjectFiles(req types.ListProjectFilesReq) (resp *types.ListProjectFilesReply, err error) {
	rsp, err := l.svcCtx.ProjectRpc.ListProjectFiles(l.ctx, &project.ListProjectFilesReq{
		Id:     req.ID,
		Path:   req.Path,
		IsWork: req.IsWork,
	})
	if err != nil {
		logx.Error("Fail to ListProjectFiles:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	res := make([]types.File, rsp.Count)

	copier.Copy(&res, &rsp.Result)

	return &types.ListProjectFilesReply{
		Count:  rsp.Count,
		Result: res,
	}, nil
}
