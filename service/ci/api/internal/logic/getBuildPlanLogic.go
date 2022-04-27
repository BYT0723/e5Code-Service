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

type GetBuildPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBuildPlanLogic {
	return GetBuildPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBuildPlanLogic) GetBuildPlan(req types.GetBuildPlanReq) (resp *types.GetBuildPlanReply, err error) {
	rsp, err := l.svcCtx.CIRpc.GetBuildPlan(l.ctx, &ci.GetBuildPlanReq{
		Id: req.Id,
	})
	if err != nil {
		logx.Error("Fail to GetBuildPlan:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	res := types.BuildPlan{}
	copier.Copy(&res, rsp)

	return &types.GetBuildPlanReply{Result: res}, nil
}
