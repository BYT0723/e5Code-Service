package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
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
	if err := l.svcCtx.ProjectModel.Delete(in.Id); err != nil {
		logx.Error("Fail to delete Project, err: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &project.DeleteProjectRsp{}, nil
}
