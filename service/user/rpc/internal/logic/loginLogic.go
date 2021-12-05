package logic

import (
	"context"
	"errors"
	"time"

	"e5Code-Service/common"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
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
	u, err := l.svcCtx.UserModel.FindOneByEmail(in.Email)
	if err != nil {
		logx.Errorf("Fail to get User(email: %s), err: %s", in.Email, err.Error())
		return nil, errors.New("UserNotExist")
	}
	if !common.ComparePwd(u.Password, in.Password) {
		l.Logger.Errorf("Password error")
		return nil, errors.New("Password error")
	}

	now := time.Now().Unix()
	info := make(map[string]interface{})
	info[common.UserID] = u.Id
	accessExpire := l.svcCtx.Config.Token.Expire
	conn := l.svcCtx.Config.Redis.NewRedis()
	token, err := conn.Get(in.Email)
	if token == "" {
		token, err = common.GenerateToken(l.svcCtx.Config.Token.JwtKey, now, accessExpire, info)
		conn.Set(in.Email, token)
	}
	if err != nil {
		logx.Error("Fail to generate token, err: ", err.Error())
		return nil, err
	}

	return &user.LoginRsp{
		Result: &user.User{
			Id:       u.Id,
			Email:    u.Email,
			Name:     u.Name,
			Password: u.Password,
		},
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
