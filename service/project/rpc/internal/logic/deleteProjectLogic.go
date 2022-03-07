package logic

import (
	"context"

	"e5Code-Service/common/contextx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
)

type DeleteProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProjectLogic {
	return &DeleteProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProjectLogic) DeleteProject(in *project.DeleteProjectReq) (*project.DeleteProjectRsp, error) {
	uID, err := contextx.GetValue(l.ctx, contextx.UserID)
	if err != nil {
		logx.Error("Fail to getUserID: ", err.Error())
		return nil, status.Error(codesx.ContextError, err.Error())
	}
	us, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: uID})
	if err != nil {
		logx.Error("Fail to getUser on DeleteProject: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}
	pj, err := l.svcCtx.ProjectModel.FindOne(in.Id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			logx.Error("Fail to GetProject on DeleteProject: ", err.Error())
			return nil, status.Error(codesx.NotFound, "NotFound")
		}
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	if err := l.svcCtx.ProjectModel.Delete(in.Id); err != nil {
		logx.Error("Fail to delete Project, err: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	if res, err := l.svcCtx.GitCli.DestoryRegistry(us.Name, pj.Name); err != nil {
		logx.Error("Fail to DestoryRegistry on deleteProject: ", err.Error())
		return nil, status.Error(codesx.GitError, res)
	}
	return &project.DeleteProjectRsp{}, nil
}
