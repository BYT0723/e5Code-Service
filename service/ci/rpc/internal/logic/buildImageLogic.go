package logic

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"e5Code-Service/common"
	"e5Code-Service/common/contextx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/common/gitx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"
	modelProject "e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/project"

	"github.com/docker/docker/api/types"
	git "github.com/go-git/go-git/v5"
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

func (l *BuildImageLogic) BuildImage(in *pb.BuildReq, stream pb.Ci_BuildImageServer) error {
	// 获取BuildPlan
	plan := &model.BuildPlan{ID: in.BuildPlanID}
	if err := l.svcCtx.DB.First(plan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Error(codesx.NotFound, "BuildPlanNotFound")
		}
		logx.Error("Fail to GetBuildPlan on BuildProject:", err.Error())
		return status.Error(codesx.SQLError, err.Error())
	}
	// 获取Project
	pj, err := l.svcCtx.ProjectRpc.GetProject(l.ctx, &project.GetProjectReq{Id: plan.ProjectID})
	if err != nil {
		logx.Error("Fail to GetProject on BuildImage: ", err.Error())
		return status.Error(codesx.RPCError, err.Error())
	}

	// 更新项目状态 -- Building
	if _, err := l.svcCtx.ProjectRpc.UpdateProject(l.ctx, &project.UpdateProjectReq{
		Id:     pj.Id,
		Status: modelProject.Building,
	}); err != nil {
		logx.Error("Fail to UpdateProject Status:", err.Error())
	}

	defer func() {
		// 更新项目状态 -- Normal
		if _, err := l.svcCtx.ProjectRpc.UpdateProject(l.ctx, &project.UpdateProjectReq{
			Id:     pj.Id,
			Status: modelProject.Normal,
		}); err != nil {
			logx.Error("Fail to UpdateProject Status:", err.Error())
		}
	}()

	// 打开本地仓库
	repositoryPath := fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.RepositoryConf.Repositories, pj.OwnerID, pj.Id)
	rep, err := git.PlainOpen(repositoryPath)
	if err != nil {
		logx.Error("Fail to Open Repository on createFile:", err.Error())
		return status.Error(codesx.GitError, err.Error())
	}

	// 打包Registry
	buildPath := repositoryPath
	if plan.Context != "" {
		buildPath = fmt.Sprintf("%s/%s", buildPath, plan.Context)
	}

	// TarLocal/ProjectID/ProjectName-PlanName
	tarName := fmt.Sprintf("%s/%s-%s.tar", l.svcCtx.Config.RepositoryConf.Tars, pj.Name, plan.Name)

	if err := gitx.TarProject(rep, tarName, buildPath); err != nil {
		logx.Error("Fail to TarProject:", err.Error())
		return status.Error(codesx.DockerError, err.Error())
	}

	// 构建镜像
	opt := types.ImageBuildOptions{
		Tags: []string{plan.Tag},
	}
	if plan.Dockerfile != "" {
		opt.Dockerfile = plan.Dockerfile
	}
	logStream, err := l.svcCtx.DockerClient.BuildImage(l.ctx, fmt.Sprintf("%s/%s", repositoryPath, tarName), opt)
	if err != nil {
		logx.Error("Fail to BuildImage: ", err.Error())
		return status.Error(codesx.DockerError, err.Error())
	}

	// 创建日志文件
	wt, err := rep.Worktree()
	if err != nil {
		logx.Error("Fail to get Worktree on BuildImage:", err.Error())
		return status.Error(codesx.GitError, err.Error())
	}
	logLocal := fmt.Sprintf("%s/%s-%s.log", l.svcCtx.Config.RepositoryConf.BuildLogs, pj.Name, plan.Version)
	logFile, err := wt.Filesystem.OpenFile(logLocal, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logx.Error("Fail to OpenLogFile on BuildImage: ", err.Error())
		return status.Error(codesx.GitError, err.Error())
	}
	defer logFile.Close()

	// 处理LogStream
	imageID := ""
	reg1 := regexp.MustCompile(`built ([a-z0-9]+)`)
	for {
		line, _, err := logStream.ReadLine()
		if err == io.EOF {
			break
		}
		// log写入日志文件
		logFile.Write(line)

		// Stream 推送Log
		if err := stream.Send(&pb.BuildRsp{LogInfo: string(line)}); err != nil {
			logx.Error("Send BuildRsp Failed...(%s)", string(line))
		}
		if reg1.Match(line) {
			imageID = strings.Split(string(reg1.Find(line)), " ")[1]
			fmt.Printf("imageID: %v\n", imageID)
		}
	}

	// 创建镜像
	userid, _ := contextx.GetValueFromMetadata(l.ctx, contextx.UserID)
	image := &model.Image{Name: plan.Tag}
	if err := l.svcCtx.DB.First(&image).Error; err == gorm.ErrRecordNotFound {
		if errx := l.svcCtx.DB.Create(&model.Image{
			ID:          common.GenUUID(),
			Name:        plan.Tag,
			ProjectID:   plan.ProjectID,
			ImageID:     imageID,
			BuildPlanID: plan.ID,
			BuilderID:   userid,
		}).Error; errx != nil {
			logx.Error("Fail to Create Image:", errx.Error())
		}
	} else {
		if errx := l.svcCtx.DB.Model(&image).Updates(model.Image{BuilderID: userid, ImageID: imageID}).Error; err != nil {
			logx.Error("Fail to Update Image:", errx.Error())
		}
	}

	// 上传镜像
	if _, err := l.svcCtx.DockerClient.UploadImage(l.ctx, plan.Tag, types.AuthConfig{}); err != nil {
		logx.Error("Fail to UploadImage on BuildImage: ", err.Error())
		stream.Send(&pb.BuildRsp{LogInfo: err.Error()})
		return status.Error(codesx.DockerError, err.Error())
	}

	// 更新plan 是否构建过
	if !plan.Already {
		if err := l.svcCtx.DB.Model(&model.BuildPlan{}).Where("id = ?", plan.ID).Update("already", true).Error; err != nil {
			logx.Error("Fail to Update Already on BuildImage:", err.Error())
			return status.Error(codesx.SQLError, err.Error())
		}
	}

	return nil
}
