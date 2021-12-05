package logic

import (
	"context"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
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
	rsp, err := l.svcCtx.UserRpc.DeleteUser(l.ctx, &user.DeleteUserReq{Id: req.Id})
	if err != nil {
		l.Logger.Errorf("Fail to delete user(id: %s)", req.Id)
		return &types.DeleteUserReply{Result: false}, err
	}

	return &types.DeleteUserReply{Result: rsp.Result}, nil
}
