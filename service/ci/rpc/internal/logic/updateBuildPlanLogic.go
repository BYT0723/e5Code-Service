package logic

import (
	"context"
	"fmt"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"
	"e5Code-Service/service/project/rpc/project"

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
	plan := &model.BuildPlan{ID: in.Id}
	if err := l.svcCtx.DB.First(plan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "BuildPlanNotFound")
		}
		logx.Error("Fail to GetBuildPlan:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	// 获取Project
	pj, err := l.svcCtx.ProjectRpc.GetProject(l.ctx, &project.GetProjectReq{Id: plan.ProjectID})
	if err != nil {
		logx.Error("Fail to GetProject on BuildImage: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	if in.Name != "" {
		plan.Name = in.Name
	}
	if in.Context != "" {
		plan.Context = in.Context
	}
	if in.Dockerfile != "" {
		plan.Dockerfile = in.Dockerfile
	}
	if in.Version != "" {
		plan.Version = in.Version
		plan.Tag = string(fmt.Sprintf("%s/%s/%s:%s", l.svcCtx.Config.ImageConf.BaseUrl, pj.Owner.Account, pj.Name, in.Version))
	}

	if err := l.svcCtx.DB.Save(plan).Error; err != nil {
		logx.Error("Fail to UpdateBuildPlan:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	return &pb.UpdateBuildPlanRsp{Tag: plan.Tag}, nil
}
