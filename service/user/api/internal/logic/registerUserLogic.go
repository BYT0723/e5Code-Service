package logic

import (
	"context"
	"e5Code-Service/common"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterUserLogic {
	return RegisterUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterUserLogic) RegisterUser(req types.RegisterUserReq) (*types.RegisterUserReply, error) {
	password, err := common.EncryptPwd(req.Password)
	if err != nil {
		logx.Errorf("Fail to encrypt Password, err: ", err.Error())
		return nil, err
	}
	rsp, err := l.svcCtx.UserRpc.AddUser(l.ctx, &user.AddUserReq{
		Email:    req.Email,
		Name:     req.Name,
		Password: password,
	})
	if err != nil {
		logx.Errorf("Fail to register User(email: %s), err: %s", req.Email, err.Error())
		return nil, err
	}
	return &types.RegisterUserReply{
		Result: types.User{
			ID:       rsp.Result.Id,
			Email:    rsp.Result.Email,
			Name:     rsp.Result.Name,
			Password: rsp.Result.Password,
		},
	}, nil
}
