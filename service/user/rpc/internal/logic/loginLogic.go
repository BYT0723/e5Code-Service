package logic

import (
	"context"
	"time"

	"e5Code-Service/common"
	"e5Code-Service/errorx"
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
		return nil, errorx.NewRpcError(errorx.UserNotFound, err.Error())
	}
	if !common.ComparePwd(u.Password, in.Password) {
		l.Logger.Errorf("Password isn't match")
		return nil, errorx.NewRpcError(errorx.PasswordNoMatch, "password isn't match")
	}

	now := time.Now().Unix()
	var accessExpire int64

	// 从redis获取token
	conn := l.svcCtx.RedisClient
	token, err := conn.Get(in.Email).Result()
	if err != nil {
		// 否则生成新token
		info := make(map[string]interface{})
		info[common.UserID] = u.Id

		token, err = common.GenerateToken(l.svcCtx.Config.Token.JwtKey, now, accessExpire, info)
		accessExpire = l.svcCtx.Config.Token.Expire
		if err != nil {
			logx.Error("Fail to generate token, err: ", err.Error())
			return nil, errorx.NewRpcError(errorx.TokenGenerateError, err.Error())
		}

		// 将新token放入redis
		if err := conn.Set(in.Email, token, time.Duration(accessExpire*int64(time.Second))).Err(); err != nil {
			logx.Error("Fail to save token to redis, err: ", err.Error())
		}
	} else {
		accessExpire = conn.TTL(in.Email).Val().Microseconds()
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
