package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/pb"
	"e5Code-Service/service/project/rpc/project"
	"e5Code-Service/service/user/rpc/user"

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

	rsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: p.OwnerId})
	if err != nil {
		logx.Error("Fail to GetUser on GetProject:", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	return &project.GetProjectRsp{
		Id:        p.ID,
		Name:      p.Name,
		Desc:      p.Desc,
		Url:       p.Url,
		Status:    p.Status,
		OwnerID:   p.OwnerId,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
		Owner: &pb.UserModel{
			ID:      rsp.Id,
			Email:   rsp.Email,
			Account: rsp.Account,
			Name:    rsp.Name,
			Bio:     rsp.Bio,
		},
	}, nil
}
