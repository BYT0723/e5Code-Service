package logic

import (
	"context"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
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
	u, err := l.svcCtx.UserModel.FindOne(in.Payload.Id)
	if err != nil {
		logx.Errorf("Fail to get user(id: %s), err: %s", in.Payload.Id, err.Error())
		return nil, err
	}
	if in.Payload.Email != "" {
		u.Email = in.Payload.Email
	}
	if in.Payload.Name != "" {
		u.Name = in.Payload.Name
	}
	if in.Payload.Password != "" {
		u.Password = in.Payload.Password
	}
	if err = l.svcCtx.UserModel.Update(*u); err != nil {
		l.Logger.Errorf("Fail to update user(id: %s)", in.Payload.Id)
		return &user.UpdateUserRsp{Result: false}, err
	}

	return &user.UpdateUserRsp{Result: true}, nil
}
