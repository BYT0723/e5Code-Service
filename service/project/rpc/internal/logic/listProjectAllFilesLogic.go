package logic

import (
	"context"
	"fmt"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/common/gitx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/pb"

	git "github.com/go-git/go-git/v5"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ListProjectAllFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProjectAllFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProjectAllFilesLogic {
	return &ListProjectAllFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProjectAllFilesLogic) ListProjectAllFiles(in *pb.ListProjectAllFilesReq) (*pb.ListProjectFilesRsp, error) {
	// 判断Project是否存在
	p := &model.Project{}
	if err := l.svcCtx.DB.Where("id = ?", in.Id).First(p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "ProjectNotFound")
		}
		logx.Error("Fail to Find Project:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	// 打开本地仓库
	rep, err := git.PlainOpen(fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.RepositoryConf.Repositories, p.OwnerId, p.ID))
	if err != nil {
		logx.Error("Fail to Open repository: ", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}

	// 拉取最新镜像
	if err := gitx.Pull(rep, "origin"); err != nil {
		if err != git.NoErrAlreadyUpToDate {
			logx.Error("Fail to Pull Latest Commit for Repository: ", err.Error())
			return nil, status.Error(codesx.GitError, err.Error())
		}
	}

	// 获取指定路径的
	files, err := gitx.ListFile(rep, "", true, in.IsWork)
	if err != nil {
		logx.Error("Fail to List File for Repository: ", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}
	count := len(files)
	result := make([]*pb.FileModel, count)
	for i := 0; i < count; i++ {
		result[i] = &pb.FileModel{}
	}
	copier.Copy(&result, &files)

	return &pb.ListProjectFilesRsp{
		Count:  int64(count),
		Result: result,
	}, nil
}
