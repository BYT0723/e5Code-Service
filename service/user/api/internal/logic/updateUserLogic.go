package logic

import (
	"context"
	"e5Code-Service/errorx"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserLogic {
	return UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req types.UpdateUserReq) (*types.UpdateUserReply, error) {
	gRsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: req.Id})
	if err != nil {
		l.Logger.Errorf("Fail to get user(id: %s)", req.Id)
		return &types.UpdateUserReply{Result: false}, errorx.NewCodeError(errorx.ServiceError, err.Error())
	}
	if gRsp.Result == nil {
		l.Logger.Errorf("User(id: %s) is not exist", req.Id)
		return &types.UpdateUserReply{Result: false}, nil
	}
	rsp, err := l.svcCtx.UserRpc.UpdateUser(l.ctx, &user.UpdateUserReq{Payload: &user.User{
		Id:       req.Id,
		Email:    req.Email,
		Name:     req.Name,
		Password: gRsp.Result.Password,
	}})
	if err != nil {
		logx.Errorf("Fail to Update User(id:%s), err: %s", req.Id, err.Error())
		return &types.UpdateUserReply{Result: false}, nil
	}
	return &types.UpdateUserReply{Result: rsp.Result}, nil
}
