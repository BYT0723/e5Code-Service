package logic

import (
	"context"
	"fmt"
	"os"

	"e5Code-Service/common/contextx"
	"e5Code-Service/common/dockerx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"
	modelProject "e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/project"

	"github.com/docker/docker/api/types"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type BuildImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBuildImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuildImageLogic {
	return &BuildImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BuildImageLogic) BuildImage(in *pb.BuildReq) (*pb.BuildRsp, error) {
	// 获取Project
	pj, err := l.svcCtx.ProjectRpc.GetProject(l.ctx, &project.GetProjectReq{Id: in.ProjectID})
	if err != nil {
		logx.Error("Fail to GetProject on BuildImage: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	// 获取UserID
	userID, err := contextx.GetValueFromMetadata(l.ctx, contextx.UserID)
	if err != nil {
		logx.Error("Fail to GetUserID on BuildImage: ", err.Error())
		return nil, status.Error(codesx.ContextError, err.Error())
	}

	// 获取BuildPlan
	plan := &model.BuildPlan{Id: in.BuildPlanID}
	if err := l.svcCtx.DB.First(plan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "BuildPlanNotFound")
		}
		logx.Error("Fail to GetBuildPlan on BuildProject:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	// 更新项目状态
	if _, err := l.svcCtx.ProjectRpc.UpdateProject(l.ctx, &project.UpdateProjectReq{
		Id:     pj.Id,
		Status: modelProject.Building,
	}); err != nil {
		logx.Error("Fail to UpdateProject Status:", err.Error())
	}

	// 打包Registry
	local := fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.RepositoryConf.Repositories, userID, pj.Id)
	tarLocal := fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.RepositoryConf.Tars, userID, pj.Id)
	if err := dockerx.TarProject(tarLocal, local); err != nil {
		logx.Error("Fail to TarProject:", err.Error())
		return nil, status.Error(codesx.DockeError, err.Error())
	}

	// 构建镜像
	reader, err := l.svcCtx.DockerClient.BuildImage(l.ctx, tarLocal, types.ImageBuildOptions{
		Tags: []string{plan.Tag},
	})
	if err != nil {
		logx.Error("Fail to BuildImage: ", err.Error())
		return nil, status.Error(codesx.DockeError, err.Error())
	}

	// 记录日志
	go func(svcCtx *svc.ServiceContext) {
		logLocal := fmt.Sprintf("%s/%s/%s-%s.log", svcCtx.Config.RepositoryConf.BuildLogs, userID, pj.Id, plan.Tag)
		logFile, err := os.OpenFile(logLocal, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			logx.Error("Fail to OpenLogFile on BuildImage: ", err.Error())
			return
		}
		defer logFile.Close()

		if _, err = reader.WriteTo(logFile); err != nil {
			logx.Error("Fail to WriteTo Log File on BuildImage: ", err.Error())
			return
		}
	}(l.svcCtx)
	return &pb.BuildRsp{}, nil
}
