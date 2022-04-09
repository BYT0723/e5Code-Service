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

type WorkStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWorkStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkStatusLogic {
	return &WorkStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WorkStatusLogic) WorkStatus(in *pb.WorkStatusReq) (*pb.WorkStatusRsp, error) {
	// 判断Project是否存在
	p := &model.Project{}
	if err := l.svcCtx.DB.Where("id = ?", in.Id).First(p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "ProjectNotFound")
		}
		logx.Error("Fail to Find Project on WorkStatus: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	// 打开本地仓库
	rep, err := git.PlainOpen(fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.RepositoryConf.Repositories, p.OwnerId, p.ID))
	if err != nil {
		logx.Error("Fail to Open Repository on WorkStatus:", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}

	fileStatus, err := gitx.GitStatus(rep)
	if err != nil {
		logx.Error("Fail to get WorkStatus: ", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}
	res := make([]*pb.FileStatus, len(fileStatus))
	for i := 0; i < len(fileStatus); i++ {
		res[i] = &pb.FileStatus{}
	}
	copier.Copy(&res, &fileStatus)

	return &pb.WorkStatusRsp{Status: res}, nil
}
