package logic

import (
	"context"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	return &user.GetUserRsp{
		Id:        u.ID,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
		Email:     in.Email,
		Account:   u.Accout,
		Name:      u.Name,
		Bio:       u.Bio,
	}, nil
}
