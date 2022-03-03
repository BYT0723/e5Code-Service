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

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteUserLogic {
	return DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req types.DeleteUserReq) (*types.DeleteUserReply, error) {
	if _, err := l.svcCtx.UserRpc.DeleteUser(l.ctx, &user.DeleteUserReq{Id: req.Id}); err != nil {
		l.Logger.Errorf("Fail to delete user(id: %s)", req.Id)
		return &types.DeleteUserReply{Result: false}, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.DeleteUserReply{Result: true}, nil
}
