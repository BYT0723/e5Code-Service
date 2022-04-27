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

type DeleteBuildPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteBuildPlanLogic {
	return DeleteBuildPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBuildPlanLogic) DeleteBuildPlan(req types.DeleteBuildPlanReq) (resp *types.DeleteBuildPlanReply, err error) {
	if _, err := l.svcCtx.CIRpc.DeleteBuildPlan(l.ctx, &ci.DeleteBuildPlanReq{Id: req.Id}); err != nil {
		logx.Error("Fail to DeleteBuildPlan:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.DeleteBuildPlanReply{Result: true}, nil
}
