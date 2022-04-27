package logic

import (
	"context"

	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/api/internal/svc"
	"e5Code-Service/service/ci/api/internal/types"
	"e5Code-Service/service/ci/rpc/ci"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBuildPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateBuildPlanLogic {
	return UpdateBuildPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBuildPlanLogic) UpdateBuildPlan(req types.UpdateBuildPlanReq) (resp *types.UpdateBuildPlanReply, err error) {
	rsp, err := l.svcCtx.CIRpc.UpdateBuildPlan(l.ctx, &ci.UpdateBuildPlanReq{
		Id:         req.Id,
		Name:       req.Name,
		Context:    req.Context,
		Dockerfile: req.Dockerfile,
		Version:    req.Version,
	})
	if err != nil {
		logx.Error("Fail to UpdateBuildPlan:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.UpdateBuildPlanReply{Result: true, Tag: rsp.Tag}, nil
}
