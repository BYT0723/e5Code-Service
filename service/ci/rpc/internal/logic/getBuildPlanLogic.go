package logic

import (
	"context"

	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBuildPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBuildPlanLogic {
	return &GetBuildPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBuildPlanLogic) GetBuildPlan(in *pb.GetBuildPlanReq) (*pb.GetBuildPlanRsp, error) {
	res := &model.BuildPlan{Id: in.Id}
	l.svcCtx.DB.First(res)

	return &pb.GetBuildPlanRsp{
		Id:         res.Id,
		ProjectID:  res.ProjectID,
		Name:       res.Name,
		Tag:        res.Tag,
		DockerFile: res.Dockerfile,
	}, nil
}
