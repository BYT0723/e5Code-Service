package logic

import (
	"context"

	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/golang/protobuf/ptypes"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectLogic {
	return &GetProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectLogic) GetProject(in *project.GetProjectReq) (*project.GetProjectRsp, error) {
	p, err := l.svcCtx.ProjectModel.FindOne(in.Id)
	if err != nil {
		logx.Errorf("Fail to find Project(Id: %s), err: %s", in.Id, err.Error())
		return nil, err
	}
	createdTime, err := ptypes.TimestampProto(p.CreateTime)
	updatedTime, err := ptypes.TimestampProto(p.UpdateTime)

	return &project.GetProjectRsp{
		Id:         p.Id,
		Name:       p.Name,
		Desc:       p.Desc.String,
		Url:        p.Url,
		OwnerID:    p.OwnerId,
		CreateTime: createdTime,
		UpdateTime: updatedTime,
	}, nil
}
