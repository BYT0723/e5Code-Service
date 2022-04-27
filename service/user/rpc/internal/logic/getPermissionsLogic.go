package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetPermissionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionsLogic {
	return &GetPermissionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPermissionsLogic) GetPermissions(in *pb.GetPermissionsReq) (*pb.GetPermissionsRsp, error) {
	permissions := []*model.Permission{}
	if err := l.svcCtx.Db.Find(&permissions, "project_id = ? and permission >= ?", in.ProjectID, in.Permission).Error; err != nil {
		logx.Error("Fail to GetPermissions:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	upMap := make(map[string]int)
	count := len(permissions)
	userIDs := make([]string, count)
	for i, p := range permissions {
		upMap[p.UserID] = p.Permission
		userIDs[i] = p.UserID
	}

	users := []*model.User{}
	if err := l.svcCtx.Db.Find(&users, "id in ?", userIDs).Error; err != nil {
		logx.Error("Fail to Find User on GetPermissions:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	userModels := make([]*pb.UserModel, count)
	copier.Copy(&userModels, &users)

	res := make([]*pb.PermissionInfo, count)
	for i := 0; i < count; i++ {
		res[i] = &pb.PermissionInfo{
			User:       userModels[i],
			Permission: int64(upMap[userModels[i].ID]),
		}
	}
	return &pb.GetPermissionsRsp{Count: int64(count), Result: res}, nil
}
