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

type DeleteSSHKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSSHKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSSHKeyLogic {
	return &DeleteSSHKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSSHKeyLogic) DeleteSSHKey(req *types.DeleteSSHKeyReq) (resp *types.DeleteSSHKeyReply, err error) {
  if _,err := l.svcCtx.UserRpc.DeleteSSHKey(l.ctx, &user.DeleteSSHKeyReq{Id: req.ID});err != nil {
    logx.Error("Fail to DeleteSSHKey:",err.Error())
    return nil,errorx.NewCodeError(codesx.RPCError, err.Error())
  }

	return &types.DeleteSSHKeyReply{Result: true},nil
}
