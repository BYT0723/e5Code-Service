package logic

import (
	"context"

	"e5Code-Service/api/pb/user"
	"e5Code-Service/common"
	"e5Code-Service/common/contextx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type AddSSHKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSSHKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSSHKeyLogic {
	return &AddSSHKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSSHKeyLogic) AddSSHKey(in *user.AddSSHKeyReq) (*user.AddSSHKeyRsp, error) {
  userid,err := contextx.GetValueFromMetadata(l.ctx, contextx.UserID)
  if err != nil {
    logx.Error("Fail to Get UserID on AddSSHKey:",err.Error())
    return nil,status.Error(codesx.ContextError, err.Error())
  }
  id := common.GenUUID()
  if err := l.svcCtx.Db.Create(&model.SSHKey{
    ID: id,
    Name: in.Name,
    Key: in.Key,
    OwnerID: userid,
  }).Error;err != nil {
    logx.Error("Fail to Create SSHKey:",err.Error())
    return nil,status.Error(codesx.SQLError, err.Error())
  }

  if res,err := l.svcCtx.GitCli.AddSSHKey(in.Key); err != nil {
    logx.Error("Fail to AddSSHKey:",res)
    return nil,status.Error(codesx.GitError, res)
  }

  return &user.AddSSHKeyRsp{Id: id}, nil
}
