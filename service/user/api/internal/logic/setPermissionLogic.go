package logic

import (
	"context"

	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetPermissionLogic {
	return SetPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetPermissionLogic) SetPermission(req types.SetPermissionReq) (resp *types.SetPermissionReply, err error) {
	if _, err := l.svcCtx.UserRpc.SetPermission(l.ctx, &pb.SetPermissionReq{
		UserID:     req.UserID,
		ProjectID:  req.ProjectID,
		Permission: req.Permission,
	}); err != nil {
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.SetPermissionReply{
		Result: true,
	}, nil
}
