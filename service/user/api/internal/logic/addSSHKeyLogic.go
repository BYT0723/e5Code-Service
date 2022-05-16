package logic

import (
	"context"

	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSSHKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSSHKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSSHKeyLogic {
	return &AddSSHKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSSHKeyLogic) AddSSHKey(req *types.AddSSHKeyReq) (resp *types.AddSSHKeyReply, err error) {
  rsp,err := l.svcCtx.UserRpc.AddSSHKey(l.ctx, &user.AddSSHKeyReq{
    Name: req.Name,
    Key: req.Key,
  })
  if err != nil {
    logx.Error("Fail to AddSSHKey:",err.Error())
    return nil,errorx.NewCodeError(codesx.RPCError, err.Error())
  }

	return &types.AddSSHKeyReply{ID: rsp.Id},nil
}
