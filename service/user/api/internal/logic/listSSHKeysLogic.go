package logic

import (
	"context"

	"e5Code-Service/common/copierx"
	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSSHKeysLogic struct {
  logx.Logger
  ctx    context.Context
  svcCtx *svc.ServiceContext
}

func NewListSSHKeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSSHKeysLogic {
  return &ListSSHKeysLogic{
    Logger: logx.WithContext(ctx),
    ctx:    ctx,
    svcCtx: svcCtx,
  }
}

func (l *ListSSHKeysLogic) ListSSHKeys(req *types.ListSSHKeysReq) (resp *types.ListSSHKeysReply, err error) {
  rsp,err := l.svcCtx.UserRpc.ListSSHKey(l.ctx, &user.ListSSHKeysReq{OwnerID: req.OwnerID})
  if err != nil {
    logx.Error("Fail to ListSSHKey:",err.Error())
    return nil,errorx.NewCodeError(codesx.RPCError, err.Error())
  }
  res := []types.SSHKey{}
  if err := copierx.Copy(&res, &rsp.Result);err != nil {
    logx.Error("Fail to Copy on ListSSHKey:",err.Error())
    return nil,errorx.NewCodeError(codesx.CopierError, err.Error())
  }

  return &types.ListSSHKeysReply{Count: rsp.Count, Result: res},nil
}
