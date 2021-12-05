package logic

import (
	"context"
	"database/sql"

	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
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
	payload := model.Project{}
	if in.Payload.Name != "" {
		payload.Name = in.Payload.Name
	}
	if in.Payload.Desc != "" {
		payload.Desc = sql.NullString{String: in.Payload.Desc, Valid: true}
	}
	if in.Payload.Url != "" {
		payload.Url = in.Payload.Url
	}
	if in.Payload.OwnerID != "" {
		payload.OwnerId = in.Payload.OwnerID
	}
	if err := l.svcCtx.ProjectModel.Update(payload); err != nil {
		logx.Error("Fail to Update Project, err: ", err.Error())
		return &project.UpdateProjectRsp{Result: false}, err
	}
	return &project.UpdateProjectRsp{Result: true}, nil
}
