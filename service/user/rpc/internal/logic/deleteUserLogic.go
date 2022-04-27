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
	// 判断用户是否存在
	if err := l.svcCtx.Db.Where("id = ?", in.Id).Find(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "UserNotFound")
		}
		logx.Errorf("Fail to get user(%s) on DeleteUser, err: %v", in.Id, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	// 开启自动事务
	if err := l.svcCtx.Db.Transaction(func(tx *gorm.DB) error {
		// 删除user-project关系(Permission)
		if err := tx.Where("user_id = ?", in.Id).Delete(&model.Permission{}).Error; err != nil {
			logx.Error("Fail to DeletePermission on DeleteUser:", err.Error())
			return status.Error(codesx.SQLError, err.Error())
		}
		// 删除用户
		if err := tx.Delete(&model.User{ID: in.Id}).Error; err != nil {
			logx.Errorf("Fail to delete user(%s), err: %v", in.Id, err.Error())
			return status.Error(codesx.SQLError, err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// 删除仓库
	if res, err := l.svcCtx.GitCli.DestoryUser(u.Account); err != nil {
		logx.Error("Fail to DestoryGitUser on DeleteUser: ", err.Error())
		return nil, status.Error(codesx.GitError, res)
	}
	return &user.DeleteUserRsp{}, nil
}
