package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListProjectLogic {
	return ListProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProjectLogic) ListProject(req types.ListProjectReq) (resp *types.ListProjectReply, err error) {
	rsp, err := l.svcCtx.ProjectRpc.ListProject(l.ctx, &pb.ListProjectReq{Filter: req.Filter})
	if err != nil {
		logx.Error("Fail to ListProject on ListProject: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	res := make([]types.Project, rsp.Count)

	for i, v := range rsp.Result {
		res[i] = types.Project{
			ID:      v.Id,
			Name:    v.Name,
			Desc:    v.Desc,
			Url:     v.Url,
			OwnerID: v.OwnerID,
			Status:  v.Status,
			Owner: types.User{
				ID:      v.Owner.Id,
				Email:   v.Owner.Email,
				Account: v.Owner.Account,
				Name:    v.Owner.Name,
			},
		}
	}

	return &types.ListProjectReply{
		Count:  rsp.Count,
		Result: res,
	}, nil
}
