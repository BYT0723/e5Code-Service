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

type CreateFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFileLogic {
	return &CreateFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFileLogic) CreateFile(in *pb.CreateFileReq) (*pb.CreateFileRsp, error) {
	// 判断Project是否存在
	p := &model.Project{}
	if err := l.svcCtx.DB.Where("id = ?", in.Id).First(p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "ProjectNotFound")
		}
		logx.Error("Fail to Find Project on createFile: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	// 打开本地仓库
	rep, err := git.PlainOpen(fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.RepositoryConf.Repositories, p.OwnerId, p.ID))
	if err != nil {
		logx.Error("Fail to Open Repository on createFile:", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}

	if err := gitx.CreateFile(rep, in.Path); err != nil {
		logx.Error("Fail to createFile:", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}
	return &pb.CreateFileRsp{}, nil
}
