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

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req types.AddUserReq) (resp *types.AddUserReply, err error) {
	if _, err := l.svcCtx.ProjectRpc.AddUser(l.ctx, &project.AddUserReq{
		UserID:    req.UserID,
		ProjectID: req.ProjectID,
	}); err != nil {
		logx.Error("Fail to AddUser: ", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.AddUserReply{
		Result: true,
	}, nil
}
