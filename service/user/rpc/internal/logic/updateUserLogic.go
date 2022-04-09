package logic

import (
	"context"
	"e5Code-Service/common/cryptx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserRsp, error) {
	u := &model.User{}
	if err := l.svcCtx.Db.Where("id = ?", in.Id).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "UserNotFound")
		}
		logx.Error("Fail to get User on UpdateUser: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	if in.Name != "" {
		u.Name = in.Name
	}
	if in.Password != "" {
		u.Password = cryptx.EncryptPwd(in.Password, l.svcCtx.Config.Salt)
	}
	if in.Bio != "" {
		u.Bio = in.Bio
	}
	if err := l.svcCtx.Db.Save(u).Error; err != nil {
		l.Logger.Errorf("Fail to update user(id: %s): %s", in.Id, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &user.UpdateUserRsp{}, nil
}
