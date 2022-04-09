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

type CommitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommitLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommitLogic {
	return CommitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommitLogic) Commit(req types.CommitReq) (resp *types.CommitReply, err error) {
	if _, err := l.svcCtx.ProjectRpc.Commit(l.ctx, &project.CommitReq{
		Id:        req.ID,
		Msg:       req.Msg,
		FilePaths: req.FilePaths,
	}); err != nil {
		logx.Error("Fail to Commit: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}
	return &types.CommitReply{
		Result: true,
	}, nil
}
