package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UpdateBuildPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBuildPlanLogic {
	return &UpdateBuildPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBuildPlanLogic) UpdateBuildPlan(in *pb.UpdateBuildPlanReq) (*pb.UpdateBuildPlanRsp, error) {
	plan := &model.BuildPlan{Id: in.Id}
	if err := l.svcCtx.DB.First(plan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "BuildPlanNotFound"))
		}
		logx.Error("Fail to GetBuildPlan:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	if in.Name != "" {
		plan.Name = in.Name
	}
	if in.Tag != "" {
		plan.Tag = in.Tag
	}
	if in.DockerFile != "" {
		plan.Dockerfile = in.DockerFile
	}

	if err := l.svcCtx.DB.Save(plan).Error; err != nil {
		logx.Error("Fail to UpdateBuildPlan:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	return &pb.UpdateBuildPlanRsp{}, nil
}
