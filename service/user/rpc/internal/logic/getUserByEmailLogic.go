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

type GetUserByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByEmailLogic {
	return &GetUserByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByEmailLogic) GetUserByEmail(in *user.GetUserByEmailReq) (*user.GetUserRsp, error) {
	rsp, err := l.svcCtx.UserModel.FindOneByEmail(in.Email)
	if err != nil {
		logx.Errorf("Fail to get user(email: %s)", in.Email)
		if err == sqlx.ErrNotFound {
			return nil, status.Error(codesx.NotFound, fmt.Sprintf("UserNotFound(email: %v)", in.Email))
		}
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	return &user.GetUserRsp{
		Id:          rsp.Id,
		CreatedTime: timestamppb.New(rsp.CreateTime),
		UpdatedTime: timestamppb.New(rsp.UpdateTime),
		Email:       in.Email,
		Name:        rsp.Name,
	}, nil
}
