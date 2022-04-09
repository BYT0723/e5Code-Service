package logic

import (
	"context"

	"e5Code-Service/common"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type AddBuildPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBuildPlanLogic {
	return &AddBuildPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBuildPlanLogic) AddBuildPlan(in *pb.AddBuildPlanReq) (*pb.AddBuildPlanRsp, error) {
	id := common.GenUUID()
	if err := l.svcCtx.DB.Create(&model.BuildPlan{
		Id:         id,
		Name:       in.Name,
		ProjectID:  in.ProjectID,
		Tag:        in.Tag,
		Dockerfile: in.DockerFile,
	}).Error; err != nil {
		logx.Error("Fail to CreateBuildPlan: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &pb.AddBuildPlanRsp{Id: id}, nil
}
