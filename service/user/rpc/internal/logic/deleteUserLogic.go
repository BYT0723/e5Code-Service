package logic

import (
	"context"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
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
		logx.Errorf("Fail to delete user(%s), err: %v", in.Id, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, status.Error(codesx.UserNotFound, "UserNotFound")
		}
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	return &user.DeleteUserRsp{}, nil
}
