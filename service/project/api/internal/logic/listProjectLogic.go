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
	rsp, err := l.svcCtx.ProjectRpc.ListProject(l.ctx, &pb.ListProjectReq{UserID: req.UserID})
	if err != nil {
		logx.Error("Fail to ListProject on ListProject: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	res := make([]types.Project, rsp.Count)
	for i, p := range rsp.Result {
		res[i] = types.Project{
			ID:      p.ID,
			Name:    p.Name,
			Desc:    p.Desc,
			Url:     p.Url,
			Status:  p.Status,
			OwnerID: p.OwnerID,
			Owner: types.User{
				ID:      p.Owner.ID,
				Email:   p.Owner.Email,
				Account: p.Owner.Account,
				Name:    p.Owner.Name,
			},
		}
	}

	return &types.ListProjectReply{
		Count:  rsp.Count,
		Result: res,
	}, nil
}
