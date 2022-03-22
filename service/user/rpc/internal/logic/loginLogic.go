package logic

import (
	"context"

	"e5Code-Service/common/cryptx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginRsp, error) {
	// 判断用户是否存在
	u := &model.User{}
	if err := l.svcCtx.Db.Where("email = ?", in.Email).First(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "UserNotFound")
		}
		logx.Errorf("Fail to get User(email: %s), err: %s", in.Email, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	// 密码验证
	if u.Password != cryptx.EncryptPwd(in.Password, l.svcCtx.Config.Salt) {
		return nil, status.Error(codesx.PasswordNotMatch, "PasswordNotMatch")
	}

	return &user.LoginRsp{
		Id:      u.ID,
		Account: u.Accout,
		Name:    u.Name,
	}, nil
}
