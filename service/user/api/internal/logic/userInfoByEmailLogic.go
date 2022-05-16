package logic

import (
	"context"

	"e5Code-Service/common/copierx"
	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoByEmailLogic {
	return UserInfoByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoByEmailLogic) UserInfoByEmail(req types.UserInfoByEmailReq) (resp *types.UserInfoReply, err error) {
	rsp, err := l.svcCtx.UserRpc.GetUserByEmail(l.ctx, &user.GetUserByEmailReq{Email: req.Email})
	if err != nil {
		logx.Error("Fail to GetUserByEmail:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	res := types.User{}
	if err := copierx.Copy(&res, &rsp.Result); err != nil {
		logx.Error("Fail to Copy on UserInfoByEmail:", err.Error())
		return nil, errorx.NewCodeError(codesx.CopierError, err.Error())
	}

	return &types.UserInfoReply{
		Result: res,
	}, nil
}
