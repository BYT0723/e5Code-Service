package logic

import (
	"context"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/golang/protobuf/ptypes"
	"github.com/tal-tech/go-zero/core/logx"
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
		l.Logger.Errorf("Fail to get user(email: %s)", in.Email)
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(rsp.CreateTime)
	if err != nil {
		logx.Error("Fail to parse user's CreateAt, err: ", err.Error)
		createTime = timestamppb.Now()
	}
	updateTime, err := ptypes.TimestampProto(rsp.UpdateTime)
	if err != nil {
		logx.Error("Fail to parse user's UpdateAt, err: ", err.Error)
		updateTime = timestamppb.Now()
	}
	return &user.GetUserRsp{
		Result: &user.User{
			Id:          rsp.Id,
			CreatedTime: createTime,
			UpdatedTime: updateTime,
			Email:       in.Email,
			Name:        rsp.Name,
			Password:    rsp.Password,
		},
	}, nil
}
