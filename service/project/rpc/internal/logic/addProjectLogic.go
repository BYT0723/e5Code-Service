package logic

import (
	"context"
	"database/sql"
	"fmt"

	"e5Code-Service/common"
	"e5Code-Service/common/contextx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type AddProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProjectLogic {
	return &AddProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProjectLogic) AddProject(in *project.AddProjectReq) (*project.AddProjectRsp, error) {
	id := common.GenUUID()
	ownerID, err := contextx.GetValue(l.ctx, contextx.UserID)
	if err != nil {
		logx.Error("Fail to getUserID from Context: ", err.Error())
		return nil, status.Error(codesx.ContextError, err.Error())
	}
	url := ""
	if in.Url == "" {
		u, _ := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: ownerID})
		if res, err := l.svcCtx.GitCli.CreateRegistry(u.Name, in.Name); err != nil {
			logx.Error("Fail to CreateRegistry on AddProject: ", err.Error())
			return nil, status.Error(codesx.GitError, res)
		}
		url = fmt.Sprintf("git@git.byt0723.xyz:%s/%s.git", u.Name, in.Name)
	} else {
		url = in.Url
	}
	payload := model.Project{
		Id:      id,
		Name:    in.Name,
		Url:     url,
		Desc:    sql.NullString{String: in.Desc, Valid: true},
		OwnerId: ownerID,
	}
	if _, err := l.svcCtx.ProjectModel.Insert(payload); err != nil {
		logx.Errorf("Fail to insert Project(Name: %s), err: %s", in.Name, err.Error())
		return nil, err
	}
	return &project.AddProjectRsp{
		Id: id,
	}, nil
}
