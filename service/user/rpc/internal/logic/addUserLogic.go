package logic

import (
	"context"
	"e5Code-Service/common"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserReq) (*user.AddUserRsp, error) {
	// todo: add your logic here and delete this line
	id := common.GetUUID()
	if _, err := l.svcCtx.UserModel.Insert(model.User{
		Id:       id,
		Email:    in.Email,
		Name:     in.Name,
		Password: in.Password,
	}); err != nil {
		l.Logger.Errorf("Fail to add user(%s)", in.Email)
		return nil, err
	}
	return &user.AddUserRsp{
		Result: &user.User{
			Id:          id,
			Email:       in.Email,
			CreatedTime: timestamppb.Now(),
			UpdatedTime: timestamppb.Now(),
			Name:        in.Name,
			Password:    in.Password,
		},
	}, nil
}
