package logic

import (
	"context"
	"e5Code-Service/api/pb/user"
	"e5Code-Service/common/copierx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GetUserByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByEmailLogic {
	return &GetUserByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByEmailLogic) GetUserByEmail(in *user.GetUserByEmailReq) (*user.GetUserRsp, error) {
	u := &model.User{}
	if err := l.svcCtx.Db.Where("email = ?", in.Email).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "UserNotFound")
		}
		logx.Errorf("Fail to get user(email: %s): %s", in.Email, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	result := &user.UserModel{}
	if err := copierx.Copy(&result, &u); err != nil {
		logx.Error("Fail to Copy on GetUserByEmail:", err.Error())
		return nil, status.Error(codesx.CopierError, err.Error())
	}

	return &user.GetUserRsp{
		Result: result,
	}, nil
}
