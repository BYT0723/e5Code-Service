package logic

import (
	"context"

	"e5Code-Service/api/pb/user"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DeleteSSHKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSSHKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSSHKeyLogic {
	return &DeleteSSHKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSSHKeyLogic) DeleteSSHKey(in *user.DeleteSSHKeyReq) (*user.DeleteSSHKeyRsp, error) {
  key := &model.SSHKey{ID: in.Id}
  if err := l.svcCtx.Db.First(key).Error; err != nil {
    logx.Error("Fail to First SSHKey:",err.Error())
    return nil,status.Error(codesx.SQLError, err.Error())
  }

  if res,err := l.svcCtx.GitCli.DeleteSSHKey(key.Key);err != nil {
    logx.Error("Fail to DeleteSSHKey:",res)
    return nil,status.Error(codesx.GitError, res)
  }

  if err := l.svcCtx.Db.Delete(key).Error;err != nil {
    logx.Error("Fail to Delete SSHKey:",err.Error())
    return nil,status.Error(codesx.SQLError, err.Error())
  }

  return &user.DeleteSSHKeyRsp{}, nil
}
