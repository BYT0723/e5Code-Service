package logic

import (
	"context"

	"e5Code-Service/api/pb/user"
	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeletePermissionLogic {
	return DeletePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePermissionLogic) DeletePermission(req types.DeletePermissionReq) (resp *types.DeletePermissionReply, err error) {
	if _, err := l.svcCtx.UserRpc.DeletePermission(l.ctx, &user.DeletePermissionReq{
		UserID:    req.UserID,
		ProjectID: req.ProjectID,
	}); err != nil {
		logx.Error("Fail to DeletePermission: ", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.DeletePermissionReply{Result: true}, nil
}
