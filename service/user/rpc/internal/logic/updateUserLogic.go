package logic

import (
	"context"
	"e5Code-Service/common/cryptx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserRsp, error) {
	u, err := l.svcCtx.UserModel.FindOne(in.Id)
	if err != nil {
		logx.Errorf("Fail to get user(id: %s), err: %s", in.Id, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, status.Error(codesx.UserNotFound, "UserNotFound")
		}
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	if in.Name != "" {
		u.Name = in.Name
	}
	if in.Password != "" {
		u.Password = cryptx.EncryptPwd(in.Password, l.svcCtx.Config.Salt)
	}
	if err = l.svcCtx.UserModel.Update(*u); err != nil {
		l.Logger.Errorf("Fail to update user(id: %s)", in.Id)
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &user.UpdateUserRsp{}, nil
}
