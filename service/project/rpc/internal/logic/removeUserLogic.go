package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/common/permission"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/pb"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserLogic {
	return &RemoveUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveUserLogic) RemoveUser(in *pb.RemoveUserReq) (*pb.RemoveUserRsp, error) {
	if _, err := l.svcCtx.UserRpc.SetPermission(l.ctx, &user.SetPermissionReq{
		UserID:     in.UserID,
		ProjectID:  in.ProjectID,
		Permission: permission.None,
	}); err != nil {
		logx.Error("Fail to SetPermission on RemoveUser: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}
	return &pb.RemoveUserRsp{}, nil
}
