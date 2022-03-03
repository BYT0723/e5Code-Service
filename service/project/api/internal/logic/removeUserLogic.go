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

type RemoveUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveUserLogic {
	return RemoveUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveUserLogic) RemoveUser(req types.RemoveUserReq) (resp *types.RemoveUserReply, err error) {
	if _, err := l.svcCtx.ProjectRpc.RemoveUser(l.ctx, &project.RemoveUserReq{
		UserID:    req.UserID,
		ProjectID: req.ProjectID,
	}); err != nil {
		logx.Error("Fail to RemoveUser: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	resp.Result = true
	return
}
