package logic

import (
	"context"
	"fmt"

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

	// 判断帐号是否激活
	if !u.Verify {
		if in.Code != "" {
			// 获取验证码
			code, err := l.svcCtx.Redis.Get(fmt.Sprintf("%s_%s", in.Email, "code")).Result()
			if err != nil {
				logx.Error("Fail to Get Verify Code:", err.Error())
				return nil, status.Error(codesx.NotVerify, err.Error())
			}
			// 验证码校验
			if code != in.Code {
				return nil, status.Error(codesx.NotVerify, "CodeNotMatch")
			}
			// 校验成功删除验证码
			if _, err := l.svcCtx.Redis.Del(fmt.Sprintf("%s_%s", in.Email, "code")).Result(); err != nil {
				logx.Error("Fail to Del Verify Code:", err.Error())
			}
			// 更新用户验证字段
			if err := l.svcCtx.Db.Model(&model.User{}).Where("id = ?", u.ID).Update("verify", true).Error; err != nil {
				logx.Error("Fail to Update User's Verify:", err.Error())
			}

			// 创建git服务器用户
			if res, err := l.svcCtx.GitCli.CreateUser(u.Account); err != nil {
				logx.Error("Fail to createGitUser on CreateUser: ", err.Error())
				return nil, status.Error(codesx.GitError, res)
			}
		} else {
			// 如果没有输入验证码则返回用户需要验证
			return nil, status.Error(codesx.NotVerify, "AccountNeedVerify")
		}
	}

	return &user.LoginRsp{
		Id:      u.ID,
		Account: u.Account,
		Name:    u.Name,
	}, nil
}
