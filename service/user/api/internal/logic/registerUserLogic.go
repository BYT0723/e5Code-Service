package logic

import (
	"context"
	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
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
	rsp, err := l.svcCtx.UserRpc.AddUser(l.ctx, &user.AddUserReq{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("Fail to register User(email: %s), err: %s", req.Email, err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.RegisterUserReply{
		Id:    rsp.Id,
		Email: rsp.Email,
		Name:  rsp.Name,
	}, nil
}
