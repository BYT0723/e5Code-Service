package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	p := &model.Permission{}
	if err := l.svcCtx.Db.Where("user_id = ? and project_id = ?", in.UserID, in.ProjectID).First(&p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "PermissionNotFound")
		}
		logx.Errorf("Fail to get permission(%s-%s) : %s", in.UserID, in.ProjectID, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &pb.GetPermissionRsp{Permission: int64(p.Permission)}, nil
}
