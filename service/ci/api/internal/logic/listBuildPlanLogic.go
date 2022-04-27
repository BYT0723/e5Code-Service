package logic

import (
	"context"

	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/api/internal/svc"
	"e5Code-Service/service/ci/api/internal/types"
	"e5Code-Service/service/ci/rpc/ci"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListBuildPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListBuildPlanLogic {
	return ListBuildPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBuildPlanLogic) ListBuildPlan(req types.ListBuildPlanReq) (resp *types.ListBuildPlanReply, err error) {
	rsp, err := l.svcCtx.CIRpc.ListBuildPlan(l.ctx, &ci.ListBuildPlanReq{
		ProjectID: req.ProjectID,
	})
	if err != nil {
		logx.Error("Fail to ListBuildPlan:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}

	res := make([]types.BuildPlan, rsp.Count)

	copier.Copy(&res, rsp.Result)

	return &types.ListBuildPlanReply{
		Count:  rsp.Count,
		Result: res,
	}, nil
}
