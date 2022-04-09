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

type DeleteProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteProjectLogic {
	return DeleteProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProjectLogic) DeleteProject(req types.DeleteProjectReq) (*types.DeleteProjectReply, error) {
	if _, err := l.svcCtx.ProjectRpc.DeleteProject(l.ctx, &project.DeleteProjectReq{Id: req.ID}); err != nil {
		logx.Error("Fail to DeleteProject: ", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}

	return &types.DeleteProjectReply{Result: true}, nil
}
