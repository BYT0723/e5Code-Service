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

type GetProjectAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProjectAuthLogic {
	return GetProjectAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectAuthLogic) GetProjectAuth(req types.GetProjectReq) (resp *types.GetProjectAuthReply, err error) {
	rsp, err := l.svcCtx.ProjectRpc.GetProjectAuth(l.ctx, &project.GetProjectReq{Id: req.ID})
	if err != nil {
		logx.Error("Fail to GetProjectAuth:", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.GetProjectAuthReply{UserName: rsp.Username, Password: rsp.Password}, nil
}
