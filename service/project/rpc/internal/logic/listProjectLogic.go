package logic

import (
	"context"

	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProjectLogic {
	return &ListProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProjectLogic) ListProject(in *pb.ListProjectReq) (*pb.ListProjectRsp, error) {
	return &pb.ListProjectRsp{}, nil
}
