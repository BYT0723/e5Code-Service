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

type AddBuildPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddBuildPlanLogic {
	return AddBuildPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBuildPlanLogic) AddBuildPlan(req types.AddBuildPlanReq) (resp *types.AddBuildPlanReply, err error) {
	rsp, err := l.svcCtx.CIRpc.AddBuildPlan(l.ctx, &ci.AddBuildPlanReq{
		Name:       req.Name,
		ProjectID:  req.ProjectID,
		Context:    req.Context,
		Dockerfile: req.Dockerfile,
		Version:    req.Version,
	})
	if err != nil {
		logx.Error("Fail to AddBuildPlan:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}

	return &types.AddBuildPlanReply{
		Id:  rsp.Id,
		Tag: rsp.Tag,
	}, nil
}
