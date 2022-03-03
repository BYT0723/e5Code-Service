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

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserLogic {
	return UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req types.UpdateUserReq) (*types.UpdateUserReply, error) {
	if _, err := l.svcCtx.UserRpc.UpdateUser(l.ctx, &user.UpdateUserReq{
		Id:       req.Id,
		Name:     req.Name,
		Password: req.Password,
	}); err != nil {
		logx.Errorf("Fail to Update User(id:%s), err: %s", req.Id, err.Error())
		return &types.UpdateUserReply{Result: false}, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.UpdateUserReply{Result: true}, nil
}
