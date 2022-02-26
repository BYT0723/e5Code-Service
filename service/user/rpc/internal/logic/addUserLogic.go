package logic

import (
	"context"
	"e5Code-Service/common"
	"e5Code-Service/common/cryptx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserReq) (*user.AddUserRsp, error) {
	id := common.GenUUID()

	// 判断email是否已被注册
	if _, err := l.svcCtx.UserModel.FindOneByEmail(in.Email); err != nil {
		if err != sqlx.ErrNotFound {
			return nil, status.Error(codesx.SQLError, err.Error())
		}
	} else {
		return nil, status.Error(codesx.UserAlreadyExist, "UserAlreadyExist")
	}

	if _, err := l.svcCtx.UserModel.Insert(model.User{
		Id:       id,
		Email:    in.Email,
		Name:     in.Name,
		Password: cryptx.EncryptPwd(in.Password, l.svcCtx.Config.Salt),
	}); err != nil {
		l.Logger.Errorf("Fail to add user(%s)", in.Email)
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &user.AddUserRsp{
		Id:    id,
		Email: in.Email,
		Name:  in.Name,
	}, nil
}
