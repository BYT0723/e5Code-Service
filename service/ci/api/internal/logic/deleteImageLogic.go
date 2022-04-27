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

type DeleteImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteImageLogic {
	return DeleteImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteImageLogic) DeleteImage(req types.DeleteImageReq) (resp *types.DeleteImageReply, err error) {
	if _, err := l.svcCtx.CIRpc.DeleteImage(l.ctx, &ci.DeleteImageReq{ID: req.ID}); err != nil {
		logx.Error("Fail to DeleteImage:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}

	return &types.DeleteImageReply{Result: true}, nil
}
