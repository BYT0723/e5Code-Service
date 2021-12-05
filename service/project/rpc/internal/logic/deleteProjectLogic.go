package logic

import (
	"context"

	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
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
		return &project.DeleteProjectRsp{Result: false}, err
	}
	return &project.DeleteProjectRsp{Result: true}, nil
}
