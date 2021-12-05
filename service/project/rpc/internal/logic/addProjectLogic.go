package logic

import (
	"context"
	"database/sql"

	"e5Code-Service/common"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
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
	id := common.GetUUID()
	ownerID := "nobody"
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		ownerID = md.Get(common.UserID)[0]
	}
	payload := model.Project{
		Id:      id,
		Name:    in.Name,
		Desc:    sql.NullString{String: in.Desc, Valid: true},
		OwnerId: ownerID,
	}
	if _, err := l.svcCtx.ProjectModel.Insert(payload); err != nil {
		logx.Errorf("Fail to insert Project(Name: %s), err: %s", in.Name, err.Error())
		return nil, err
	}
	return &project.AddProjectRsp{
		Result: &project.Project{
			Id:      id,
			Name:    in.Name,
			Desc:    in.Desc,
			Url:     in.Url,
			OwnerID: ownerID,
		},
	}, nil
}
