package logic

import (
	"context"

	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateProjectLogic {
	return UpdateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProjectLogic) UpdateProject(req types.UpdateProjectReq) (*types.UpdateProjectReply, error) {
	if _, err := l.svcCtx.ProjectRpc.UpdateProject(l.ctx, &project.UpdateProjectReq{
		Id:   req.ID,
		Name: req.Name,
		Desc: req.Desc,
		Url:  req.Url,
	}); err != nil {
		logx.Error("Fail to UpdateProject: ", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.UpdateProjectReply{Result: true}, nil
}
