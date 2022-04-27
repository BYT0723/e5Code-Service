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
	"google.golang.org/grpc/status"
)

type GetImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetImageLogic {
	return GetImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetImageLogic) GetImage(req types.GetImageReq) (resp *types.GetImageReply, err error) {
	rsp, err := l.svcCtx.CIRpc.GetImage(l.ctx, &ci.GetImageReq{ID: req.ID})
	if err != nil {
		logx.Error("Fail to GetImage:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	res := types.Image{}
	if err := copier.CopyWithOption(&res, &rsp.Result, copier.Option{DeepCopy: true}); err != nil {
		logx.Error("Fail to CopyWithOption on GetImage:", err.Error())
		return nil, status.Error(codesx.CopierError, err.Error())
	}

	return &types.GetImageReply{Result: res}, nil
}
