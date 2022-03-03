package logic

import (
	"context"
	"database/sql"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectLogic {
	return &UpdateProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProjectLogic) UpdateProject(in *project.UpdateProjectReq) (*project.UpdateProjectRsp, error) {
	pj, err := l.svcCtx.ProjectModel.FindOne(in.Id)
	if err != nil {
		logx.Error("Fail to GetProject on UpdateProject: ", err.Error())
		return nil, status.Error(codesx.NotFound, err.Error())
	}
	if in.Name != "" {
		pj.Name = in.Name
	}
	if in.Desc != "" {
		pj.Desc = sql.NullString{String: in.Desc, Valid: true}
	}
	if in.Url != "" {
		pj.Url = in.Url
	}
	if err := l.svcCtx.ProjectModel.Update(*pj); err != nil {
		logx.Error("Fail to UpdateProject : ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &project.UpdateProjectRsp{}, nil
}
