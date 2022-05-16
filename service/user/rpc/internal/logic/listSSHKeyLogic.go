package logic

import (
	"context"

	"e5Code-Service/api/pb/user"
	"e5Code-Service/common/copierx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListSSHKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSSHKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSSHKeyLogic {
	return &ListSSHKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListSSHKeyLogic) ListSSHKey(in *user.ListSSHKeysReq) (*user.ListSSHKeysRsp, error) {
  keys := []*model.SSHKey{}
  if err := l.svcCtx.Db.Find(&keys, "owner_id = ?",in.OwnerID).Error;err != nil {
    logx.Error("Fail to Find SSHKey:",err.Error())
    return nil,status.Error(codesx.SQLError,err.Error())
  }

  res := []*user.SSHKey{}
  if err := copierx.Copy(&res, &keys);err != nil {
    logx.Error("Fail to Copy on ListSSHKey:",err.Error())
    return nil,status.Error(codesx.CopierError, err.Error())
  }

  return &user.ListSSHKeysRsp{
    Count: int64(len(res)),
    Result: res,
  }, nil
}
