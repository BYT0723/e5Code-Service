package logic

import (
	"context"
	"e5Code-Service/common"
	"e5Code-Service/common/cryptx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

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
	if err := l.svcCtx.Db.Where("email = ?", in.Email).First(&model.User{}).Error; err == nil {
		return nil, status.Error(codesx.AlreadyExist, "UserAlreadyExist")
	}

	if res, err := l.svcCtx.GitCli.CreateUser(in.Account); err != nil {
		logx.Error("Fail to createGitUser on CreateUser: ", err.Error())
		return nil, status.Error(codesx.GitError, res)
	}

	if err := l.svcCtx.Db.Model(&model.User{}).Create(&model.User{
		ID:       id,
		Email:    in.Email,
		Accout:   in.Account,
		Name:     in.Name,
		Bio:      in.Bio,
		Password: cryptx.EncryptPwd(in.Password, l.svcCtx.Config.Salt),
	}).Error; err != nil {
		l.Logger.Errorf("Fail to add user(%s): %s", in.Email, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	return &user.AddUserRsp{
		Id: id,
	}, nil
}
