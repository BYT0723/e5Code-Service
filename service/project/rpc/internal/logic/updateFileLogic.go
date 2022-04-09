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

type UpdateFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileLogic {
	return &UpdateFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFileLogic) UpdateFile(in *pb.UpdateFileReq) (*pb.UpdateFileRsp, error) {
	// 判断Project是否存在
	p := &model.Project{}
	if err := l.svcCtx.DB.Where("id = ?", in.Id).First(p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "ProjectNotFound")
		}
		logx.Error("Fail to Find Project on ReadFile: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	// 打开本地仓库
	rep, err := git.PlainOpen(fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.RepositoryConf.Repositories, p.OwnerId, p.ID))
	if err != nil {
		logx.Error("Fail to Open Repository on ReadFile:", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}

	// 更新文件
	if err := gitx.UpdateFile(rep, in.Path, []byte(in.Body)); err != nil {
		logx.Errorf("Fail to UpdateFile(%s): %s", in.Path, err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}

	return &pb.UpdateFileRsp{}, nil
}
