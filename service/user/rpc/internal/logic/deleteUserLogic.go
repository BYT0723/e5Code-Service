package logic

import (
	"context"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserReq) (*user.DeleteUserRsp, error) {
	err := l.svcCtx.UserModel.Delete(in.Id)
	if err != nil {
		l.Logger.Errorf("Fail to delete user(%s)", in.Id)
		return &user.DeleteUserRsp{
			Result: false,
		}, nil
	}

	return &user.DeleteUserRsp{Result: true}, nil
}
