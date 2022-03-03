package logic

import (
	"context"

	"e5Code-Service/common"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
)

type SetPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPermissionLogic {
	return &SetPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetPermissionLogic) SetPermission(in *pb.SetPermissionReq) (*pb.SetPermissionRsp, error) {
	if pm, err := l.svcCtx.PermissionModel.FindOneByUserIdProjectId(in.UserID, in.ProjectID); err == sqlx.ErrNotFound {
		if _, err := l.svcCtx.PermissionModel.Insert(&model.Permission{
			Id:         common.GenUUID(),
			UserId:     in.UserID,
			ProjectId:  in.ProjectID,
			Permission: in.Permission,
		}); err != nil {
			logx.Error("Fail to insert permission on SetPermission :", err.Error())
			return nil, status.Error(codesx.SQLError, err.Error())
		}
	} else if err == nil {
		if err := l.svcCtx.PermissionModel.Update(&model.Permission{
			Id:         pm.Id,
			UserId:     pm.UserId,
			ProjectId:  pm.ProjectId,
			Permission: in.Permission,
		}); err != nil {
			logx.Error("Fail to update permission on SetPermission :", err.Error())
			return nil, status.Error(codesx.SQLError, err.Error())
		}
	}
	return &pb.SetPermissionRsp{}, nil
}
