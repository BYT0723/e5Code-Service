package logic

import (
	"context"
	"e5Code-Service/common/contextx"
	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/common/jwtx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginReply, error) {
	rsp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("Fail to Login(email: %s), err: %s", req.Email, err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}

	now := time.Now().Unix()
	var accessExpire int64

	token, err := l.svcCtx.Redis.Get(rsp.Email).Result()
	if err != nil {
		// 否则生成新token
		accessExpire = l.svcCtx.Config.Auth.AccessExpire
		token, err = jwtx.GenerateToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, map[string]interface{}{
			contextx.UserID: rsp.Id,
		})
		if err != nil {
			logx.Error("Fail to generate token, err: ", err.Error())
			return nil, errorx.NewCodeError(codesx.TokenGenerateError, err.Error())
		}

		// 将新token放入redis
		if err = l.svcCtx.Redis.Set(req.Email, token, time.Duration(accessExpire*int64(time.Second))).Err(); err != nil {
			logx.Error("Fail to save token to redis, err: ", err.Error())
		}
	} else {
		accessExpire = int64(l.svcCtx.Redis.TTL(req.Email).Val().Seconds())
	}

	return &types.LoginReply{
		Id:           rsp.Id,
		AccessToken:  token,
		AccessExpire: accessExpire,
	}, nil
}
