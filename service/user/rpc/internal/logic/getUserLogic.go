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

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserRsp, error) {
	u := &model.User{}
	if err := l.svcCtx.Db.Where("id = ?", in.Id).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "UserNotFound")
		}
		logx.Errorf("Fail to get user(id: %s): %s", in.Id, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	res := &user.UserModel{}
	if err := copierx.Copy(&res, &u); err != nil {
		logx.Error("Fail to Copy on GetUser:", err.Error())
		return nil, status.Error(codesx.CopierError, err.Error())
	}

	return &user.GetUserRsp{
		Result: res,
	}, nil
}
