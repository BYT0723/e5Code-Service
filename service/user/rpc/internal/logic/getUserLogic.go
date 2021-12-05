package logic

import (
	"context"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/golang/protobuf/ptypes"
	"github.com/tal-tech/go-zero/core/logx"
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
		l.Logger.Errorf("Fail to get user(id: %s)", in.Id)
		return nil, err
	}
	createTime, err := ptypes.TimestampProto(rsp.CreateTime)
	if err != nil {
		logx.Error("Fail to parse user's CreateTime, err: ", err.Error)
		createTime = timestamppb.Now()
	}
	updateTime, err := ptypes.TimestampProto(rsp.UpdateTime)
	if err != nil {
		logx.Error("Fail to parse user's UpdateTime, err: ", err.Error)
		updateTime = timestamppb.Now()
	}

	return &user.GetUserRsp{
		Result: &user.User{
			CreatedTime: createTime,
			UpdatedTime: updateTime,
			Id:          in.Id,
			Email:       rsp.Email,
			Name:        rsp.Name,
			Password:    rsp.Password,
		},
	}, nil
}
