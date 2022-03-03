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

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *pb.AddUserReq) (*pb.AddUserRsp, error) {
	if _, err := l.svcCtx.UserRpc.SetPermission(l.ctx, &user.SetPermissionReq{
		UserID:     in.UserID,
		ProjectID:  in.ProjectID,
		Permission: permission.Read,
	}); err != nil {
		logx.Error("Fail to AddPermission on AddUserToProject: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}
	return &pb.AddUserRsp{}, nil
}
