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

type ModifyPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) ModifyPermissionLogic {
	return ModifyPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyPermissionLogic) ModifyPermission(req types.ModifyPermissionReq) (resp *types.ModifyPermissionReply, err error) {
	if _, err := l.svcCtx.ProjectRpc.ModifyPermission(l.ctx, &project.ModifyPermissionReq{
		UserID:       req.UserID,
		ProjectID:    req.ProjectID,
		ModifiedType: req.ModifiedType,
		Value:        req.Value,
	}); err != nil {
		logx.Error("Fail to ModifyPermission: ", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}

	resp.Result = true
	return
}
