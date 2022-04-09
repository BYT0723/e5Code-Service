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

type GetProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProjectLogic {
	return GetProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectLogic) GetProject(req types.GetProjectReq) (*types.GetProjectReply, error) {
	rsp, err := l.svcCtx.ProjectRpc.GetProject(l.ctx, &project.GetProjectReq{
		Id: req.ID,
	})
	if err != nil {
		logx.Error("Fail to GetProject: ", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}

	return &types.GetProjectReply{
		ID:      rsp.Id,
		Name:    rsp.Name,
		Desc:    rsp.Desc,
		Url:     rsp.Url,
		OwnerID: rsp.OwnerID,
	}, nil
}
