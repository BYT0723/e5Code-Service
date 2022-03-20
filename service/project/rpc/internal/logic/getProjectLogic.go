package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
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
	p := &model.Project{}
	if err := l.svcCtx.DB.Where("id = ?", in.Id).First(p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "ProjectNotFound")
		}
		logx.Errorf("Fail to find Project(Id: %s), err: %s", in.Id, err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	return &project.GetProjectRsp{
		Id:        p.ID,
		Name:      p.Name,
		Desc:      p.Desc,
		Url:       p.Url,
		OwnerID:   p.OwnerId,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
	}, nil
}
