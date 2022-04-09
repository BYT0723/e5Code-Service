package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListBuildPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBuildPlanLogic {
	return &ListBuildPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListBuildPlanLogic) ListBuildPlan(in *pb.ListBuildPlanReq) (*pb.ListBuildPlanRsp, error) {
	plans := []*model.BuildPlan{}
	if err := l.svcCtx.DB.Where("project_id = ?", in.ProjectID).Find(&plans).Error; err != nil {
		logx.Error("Fail to ListBuildPlan:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	count := len(plans)
	res := make([]*pb.BuildPlanModel, count)
	for i := 0; i < count; i++ {
		res[i] = &pb.BuildPlanModel{}
	}

	copier.Copy(&res, &plans)

	return &pb.ListBuildPlanRsp{
		Count:  int64(count),
		Result: res,
	}, nil
}
