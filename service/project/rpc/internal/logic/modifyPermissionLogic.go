package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/pb"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ModifyPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyPermissionLogic {
	return &ModifyPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyPermissionLogic) ModifyPermission(in *pb.ModifyPermissionReq) (*pb.ModifyPermissionRsp, error) {
	if _, err := l.svcCtx.UserRpc.SetPermission(l.ctx, &user.SetPermissionReq{
		UserID:     in.UserID,
		ProjectID:  in.ProjectID,
		Permission: in.Value,
	}); err != nil {
		logx.Error("Fail to SetPermission on ModifyPermission: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}
	return &pb.ModifyPermissionRsp{}, nil
}
