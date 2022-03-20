package logic

import (
	"context"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	u := &model.User{}
	if err := l.svcCtx.Db.Where("id = ?", in.Id).Find(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "UserNotFound")
		}
		logx.Errorf("Fail to get user(%s) on DeleteUser, err: %v", in.Id, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	if res, err := l.svcCtx.GitCli.DestoryUser(u.Accout); err != nil {
		logx.Error("Fail to DestoryGitUser on DeleteUser: ", err.Error())
		return nil, status.Error(codesx.GitError, res)
	}

	if err := l.svcCtx.Db.Delete(&model.User{ID: in.Id}).Error; err != nil {
		logx.Errorf("Fail to delete user(%s), err: %v", in.Id, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &user.DeleteUserRsp{}, nil
}
