package logic

import (
	"context"

	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPermissionsLogic {
	return GetPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionsLogic) GetPermissions(req types.GetPermissionsReq) (resp *types.GetPermissionsReply, err error) {
	rsp, err := l.svcCtx.UserRpc.GetPermissions(l.ctx, &user.GetPermissionsReq{ProjectID: req.ProjectID, Permission: req.Permission})
	if err != nil {
		logx.Error("Fail to GetPermissions:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	res := []types.PermissionInfo{}
	copier.Copy(&res, &rsp.Result)

	return &types.GetPermissionsReply{Count: rsp.Count, Result: res}, nil
}
