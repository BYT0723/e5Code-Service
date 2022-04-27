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

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req types.UserInfoReq) (resp *types.UserInfoReply, err error) {
	rsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: req.Id})
	if err != nil {
		logx.Error("Fail to getUser, err: ", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	resp = &types.UserInfoReply{
		Id:      rsp.Id,
		Email:   rsp.Email,
		Account: rsp.Account,
		Name:    rsp.Name,
		Bio:     rsp.Bio,
	}
	return
}
