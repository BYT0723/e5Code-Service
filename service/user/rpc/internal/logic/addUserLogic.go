package logic

import (
	"context"
	"e5Code-Service/api/pb/user"
	"e5Code-Service/common"
	"e5Code-Service/common/cryptx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/common/mailx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"fmt"
	"math/rand"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
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
	u := &model.User{}
	if err := l.svcCtx.Db.Where("email = ?", in.Email).First(u).Error; err == nil {
		if !u.Verify {
			l.svcCtx.Db.Delete(u)
		} else {
			return nil, status.Error(codesx.AlreadyExist, "UserAlreadyExist")
		}
	}

	if err := l.svcCtx.Db.Model(&model.User{}).Create(&model.User{
		ID:       id,
		Email:    in.Email,
		Account:  in.Account,
		Name:     in.Name,
		Bio:      in.Bio,
		Password: cryptx.EncryptPwd(in.Password, l.svcCtx.Config.Salt),
	}).Error; err != nil {
		logx.Errorf("Fail to add user(%s): %s", in.Email, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	go func() {
		// 存储验证码
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		code := fmt.Sprintf("%06d", r.Intn(100000))
		if _, err := l.svcCtx.Redis.Set(fmt.Sprintf("%s_%s", in.Email, "code"), code, time.Duration(time.Minute*10)).Result(); err != nil {
			logx.Error("Fail to Set Verify Code:", err.Error())
			return
		}

		// 发送验证邮件
		dialer := mailx.NewDialer()
		mes := mailx.NewMessage(mailx.Admin, in.Email, mailx.VerifyTitle, mailx.GenBody(mailx.VerifyTemplate, code))
		if err := dialer.DialAndSend(mes); err != nil {
			logx.Error("Fail to Send Verify Mail:", err.Error())
			return
		}
	}()

	return &user.AddUserRsp{
		Id: id,
	}, nil
}
