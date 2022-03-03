package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
)

type GetPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionLogic {
	return &GetPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPermissionLogic) GetPermission(in *pb.GetPermissionReq) (*pb.GetPermissionRsp, error) {
	ps, err := l.svcCtx.PermissionModel.FindOneByUserIdProjectId(in.UserID, in.ProjectID)
	if err != nil {
		logx.Errorf("Fail to get permission(%s-%s) : %s", in.UserID, in.ProjectID, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, status.Error(codesx.NotFound, "PermissionNotFound")
		}
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &pb.GetPermissionRsp{Permission: ps.Permission}, nil
}
