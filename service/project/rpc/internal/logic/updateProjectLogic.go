package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	p := &model.Project{}
	if err := l.svcCtx.DB.Where("id = ?", in.Id).First(p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "ProjectNotFound")
		}
		logx.Error("Fail to GetProject on UpdateProject: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	if in.Name != "" {
		p.Name = in.Name
	}
	if in.Desc != "" {
		p.Desc = in.Desc
	}
	if in.Url != "" {
		p.Url = in.Url
	}
	if err := l.svcCtx.DB.Save(p).Error; err != nil {
		logx.Error("Fail to UpdateProject : ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &project.UpdateProjectRsp{}, nil
}
