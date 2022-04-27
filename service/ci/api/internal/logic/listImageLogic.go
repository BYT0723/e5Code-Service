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
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ListImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListImageLogic {
	return ListImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListImageLogic) ListImage(req types.ListImageReq) (resp *types.ListImageReply, err error) {
	rsp, err := l.svcCtx.CIRpc.ListImage(l.ctx, &ci.ListImageReq{ProjectID: req.ProjectID})
	if err != nil {
		logx.Error("Fail to ListImage:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	res := []types.Image{}
	if err := copier.CopyWithOption(&res, &rsp.Result, copier.Option{
		DeepCopy: true,
		Converters: []copier.TypeConverter{
			{
				SrcType: &timestamppb.Timestamp{},
				DstType: int64(0),
				Fn: func(src interface{}) (interface{}, error) {
					temp, _ := src.(*timestamppb.Timestamp)
					return temp.AsTime().UnixNano(), nil
				},
			},
		},
	}); err != nil {
		logx.Error("Fail to CopyWithOption on ListImage:", err.Error())
		return nil, status.Error(codesx.CopierError, err.Error())
	}

	return &types.ListImageReply{
		Count:  rsp.Count,
		Result: res,
	}, nil
}
