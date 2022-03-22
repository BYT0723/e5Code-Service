package logic

import (
	"context"
	"fmt"
	"os"

	"e5Code-Service/common/contextx"
	"e5Code-Service/common/dockerx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/common/gitx"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"
	"e5Code-Service/service/project/rpc/project"

	"github.com/docker/docker/api/types"
	"github.com/go-git/go-git/v5"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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

	// Clone Registry
	local := fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.RegistryConf.Local, userID, pj.Id)
	if err := gitx.Clone(gitx.GitCloneOpt{
		Local: local,
		CloneOptions: &git.CloneOptions{
			URL: pj.Url,
		},
	}); err != nil {
		logx.Error("Fail to Clone Registry: ", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}

	// 打包Registry
	tarLocal := fmt.Sprintf("%s/%s/%s.tar", l.svcCtx.Config.RegistryConf.Tar, userID, pj.Id)
	dockerx.TarProject(tarLocal, local)

	// 构建镜像
	reader, err := l.svcCtx.DockerClient.BuildImage(l.ctx, tarLocal, types.ImageBuildOptions{
		Tags:    []string{in.Tag},
		Version: types.BuilderVersion(in.Version),
	})
	if err != nil {
		logx.Error("Fail to BuildImage: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	// 记录日志
	logLocal := fmt.Sprintf("%s/%s/%s-%s:%s.log", l.svcCtx.Config.RegistryConf.BuildLog, userID, pj.Id, in.Tag, in.Version)
	logFile, err := os.OpenFile(logLocal, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logx.Error("Fail to OpenLogFile on BuildImage: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}
	defer logFile.Close()

	if _, err = reader.WriteTo(logFile); err != nil {
		logx.Error("Fail to WriteTo Log File on BuildImage: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}
	return &pb.BuildRsp{}, nil
}
