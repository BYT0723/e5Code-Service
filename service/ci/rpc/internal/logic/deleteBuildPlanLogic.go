package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DeleteBuildPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBuildPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBuildPlanLogic {
	return &DeleteBuildPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBuildPlanLogic) DeleteBuildPlan(in *pb.DeleteBuildPlanReq) (*pb.DeleteBuildPlanRsp, error) {
	if err := l.svcCtx.DB.Delete(&model.BuildPlan{Id: in.Id}).Error; err != nil {
		logx.Error("Fail to Delete BuildPlan: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &pb.DeleteBuildPlanRsp{}, nil
}
