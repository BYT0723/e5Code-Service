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

type GetPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPermissionLogic {
	return GetPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionLogic) GetPermission(req types.GetPermissionReq) (resp *types.GetPermissionReply, err error) {
	pm, err := l.svcCtx.UserRpc.GetPermission(l.ctx, &pb.GetPermissionReq{
		UserID:    req.UserID,
		ProjectID: req.ProjectID,
	})
	if err != nil {
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	resp.Permission = pm.Permission
	return
}
