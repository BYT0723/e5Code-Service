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
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ListProjectFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProjectFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProjectFilesLogic {
	return &ListProjectFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProjectFilesLogic) ListProjectFiles(in *pb.ListProjectFilesReq) (*pb.ListProjectFilesRsp, error) {
	// 判断Project是否存在
	p := &model.Project{}
	if err := l.svcCtx.DB.Where("id = ?", in.Id).First(p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, err.Error())
		}
		logx.Error("Fail to Find Project:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	// 打开本地仓库
	rep, err := git.PlainOpen(fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.RegistryConf.Local, p.OwnerId, p.ID))
	if err != nil {
		logx.Error("Fail to Open repository: ", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}

	if err := gitx.Pull(rep, "origin"); err != nil {
	}

	w, err := rep.Worktree()
	w.Pull(&git.PullOptions{RemoteName: "origin"})

	ref, err := rep.Head()
	if err != nil {
		logx.Error("Fail to get Head from repository:", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}
	rep.CommitObject(ref.Hash())

	return &pb.ListProjectFilesRsp{}, nil
}
