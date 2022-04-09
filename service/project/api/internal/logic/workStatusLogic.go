package logic

import (
	"context"
	"fmt"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type WorkStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) WorkStatusLogic {
	return WorkStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WorkStatusLogic) WorkStatus(req types.WorkStatusReq) (resp *types.WorkStatusReply, err error) {
	rsp, err := l.svcCtx.ProjectRpc.WorkStatus(l.ctx, &project.WorkStatusReq{Id: req.ID})
	if err != nil {
		logx.Error("Fail to get WorkStatus:", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}
	res := make([]types.FileStatus, len(rsp.Status))
	copier.Copy(&res, rsp.Status)
	fmt.Printf("res: %v\n", res)

	return &types.WorkStatusReply{Result: res}, nil
}
