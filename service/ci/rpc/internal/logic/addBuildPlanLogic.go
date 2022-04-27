package logic

import (
	"context"
	"fmt"
	"strings"

	"e5Code-Service/common"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"
	"e5Code-Service/service/project/rpc/project"

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
	// 获取Project
	pj, err := l.svcCtx.ProjectRpc.GetProject(l.ctx, &project.GetProjectReq{Id: in.ProjectID})
	if err != nil {
		logx.Error("Fail to GetProject on BuildImage: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	id := common.GenUUID()
	tag := strings.ToLower(fmt.Sprintf("%s/%s/%s:%s", l.svcCtx.Config.ImageConf.BaseUrl, pj.Owner.Account, pj.Name, in.Version))
	if err := l.svcCtx.DB.Create(&model.BuildPlan{
		ID:         id,
		Name:       in.Name,
		ProjectID:  in.ProjectID,
		Context:    in.Context,
		Dockerfile: in.Dockerfile,
		Version:    in.Version,
		Tag:        tag,
		Already:    false,
	}).Error; err != nil {
		logx.Error("Fail to CreateBuildPlan: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &pb.AddBuildPlanRsp{
		Id:  id,
		Tag: tag,
	}, nil
}
