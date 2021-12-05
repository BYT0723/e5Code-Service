package logic

import (
	"context"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginReply, error) {
	rsp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("Fail to Login(email: %s), err: %s", req.Email, err.Error())
		return nil, err
	}
	return &types.LoginReply{
		Result: types.User{
			ID:       rsp.Result.Id,
			Email:    rsp.Result.Email,
			Name:     rsp.Result.Name,
			Password: rsp.Result.Password,
		},
		AccessToken:  rsp.AccessToken,
		AccessExpire: rsp.AccessExpire,
		RefreshAfter: rsp.RefreshAfter,
	}, nil
}
