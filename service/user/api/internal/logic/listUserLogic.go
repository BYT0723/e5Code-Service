package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListUserLogic {
	return ListUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLogic) ListUser(req types.ListUserReq) (resp *types.ListUserReply, err error) {
	rsp, err := l.svcCtx.UserRpc.ListUser(l.ctx, &pb.ListUserReq{Ids: req.IDs})
	if err != nil {
		logx.Error("Fail to ListUser on ListUser: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	res := make([]types.User, rsp.Count)
	copier.Copy(&res, &rsp.Result)

	return &types.ListUserReply{
		Count:  rsp.Count,
		Result: res,
	}, nil
}
