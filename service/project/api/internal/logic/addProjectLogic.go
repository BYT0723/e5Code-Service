package logic

import (
	"context"
	"fmt"

	"e5Code-Service/common"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

type AddProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddProjectLogic {
	return AddProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProjectLogic) AddProject(req types.AddProjectReq) (*types.AddProjectReply, error) {
	if userID, ok := l.ctx.Value(common.UserID).(string); ok {
		fmt.Printf("userID: %v\n", userID)
		md := metadata.Pairs("userID", userID)
		l.ctx = metadata.NewOutgoingContext(l.ctx, md)
	} else {
		logx.Error("Fail to get UserID")
	}
	rsp, err := l.svcCtx.ProjectServer.AddProject(l.ctx, &project.AddProjectReq{
		Name: req.Name,
		Desc: req.Desc,
		Url:  req.Url,
	})
	if err != nil {
		logx.Error("Fail to add Project, err: ", err.Error())
		return nil, err
	}
	return &types.AddProjectReply{
		Result: types.Project{
			ID:      rsp.Result.Id,
			Name:    rsp.Result.Name,
			Desc:    rsp.Result.Desc,
			Url:     rsp.Result.Url,
			OwnerId: rsp.Result.OwnerID,
		},
	}, nil
}
