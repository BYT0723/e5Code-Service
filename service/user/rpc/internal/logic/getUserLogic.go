package logic

import (
	"context"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserRsp, error) {
	rsp, err := l.svcCtx.UserModel.FindOne(in.Id)
	if err != nil {
		logx.Errorf("Fail to get user(id: %s)", in.Id)
		if err == sqlx.ErrNotFound {
			return nil, status.Error(codesx.NotFound, fmt.Sprintf("UserNotFound(id: %v)", in.Id))
		}
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	return &user.GetUserRsp{
		CreatedTime: timestamppb.New(rsp.CreateTime),
		UpdatedTime: timestamppb.New(rsp.UpdateTime),
		Id:          in.Id,
		Email:       rsp.Email,
		Name:        rsp.Name,
	}, nil
}
