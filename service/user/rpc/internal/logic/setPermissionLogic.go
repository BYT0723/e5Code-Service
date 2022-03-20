package logic

import (
	"context"

	"e5Code-Service/common"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	p := &model.Permission{}
	if err := l.svcCtx.Db.Where("user_id = ? and project_id = ?", in.UserID, in.ProjectID).First(&p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err2 := l.svcCtx.Db.Model(&model.Permission{}).Create(&model.Permission{
				ID:         common.GenUUID(),
				UserID:     in.UserID,
				ProjectID:  in.ProjectID,
				Permission: int(in.Permission),
			}).Error; err2 != nil {
				logx.Error("Fail to insert permission on SetPermission :", err2.Error())
				return nil, status.Error(codesx.SQLError, err2.Error())
			}
		} else {
			logx.Error("Fail to select permission on SetPermission :", err.Error())
			return nil, status.Error(codesx.SQLError, err.Error())
		}
	} else {
		if err2 := l.svcCtx.Db.Model(&p).Update("permission", in.Permission).Error; err2 != nil {
			logx.Error("Fail to update permission on SetPermission :", err2.Error())
			return nil, status.Error(codesx.SQLError, err2.Error())
		}
	}
	return &pb.SetPermissionRsp{}, nil
}
