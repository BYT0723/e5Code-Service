package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"e5Code-Service/common/contextx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/common/gitx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/pb"
	"e5Code-Service/service/user/rpc/user"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CommitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommitLogic {
	return &CommitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommitLogic) Commit(in *pb.CommitReq) (*pb.CommitRsp, error) {
	uid, err := contextx.GetValueFromMetadata(l.ctx, contextx.UserID)
	if err != nil {
		logx.Error("Fail to get UserID: ", err.Error())
		return nil, err
	}
	rsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: uid})
	if err != nil {
		logx.Error("Fail to get User: ", err.Error())
		return nil, err
	}
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
	auth := &http.BasicAuth{}
	json.Unmarshal([]byte(p.Auth), auth)

	if err := gitx.Commit(rep, &gitx.CommitOption{
		FilePaths: in.FilePaths,
		Msg:       in.Msg,
		Author:    rsp.Account,
		Email:     rsp.Email,
		Remote:    "origin",
		BasicAuth: auth,
	}); err != nil {
		logx.Error("Fail to Commit:", err.Error())
		return nil, status.Error(codesx.GitError, err.Error())
	}

	return &pb.CommitRsp{}, nil
}
