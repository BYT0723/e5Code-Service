package logic

import (
	"context"
	"e5Code-Service/common"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"e5Code-Service/service/user/rpc/user"
	"errors"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatePasswordLogic {
	return UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req types.UpdatePasswordReq) (*types.UpdateUserReply, error) {
	gRsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: req.Id})
	if err != nil {
		logx.Errorf("Fail to get user(id: %s)", req.Id)
		return &types.UpdateUserReply{Result: false}, err
	}
	if gRsp.Result == nil {
		logx.Errorf("User(id: %s) is not exist", req.Id)
		return &types.UpdateUserReply{Result: false}, errors.New("User isn't exist")
	}
	if !common.ComparePwd(gRsp.Result.Password, req.OldPassword) {
		logx.Errorf("Password don't match")
		return &types.UpdateUserReply{Result: false}, errors.New("Password don't match")
	}
	newPassword, err := common.EncryptPwd(req.NewPassword)
	if err != nil {
		logx.Error("Fail to encryptPassword, err: ", err.Error())
		return nil, errors.New("EncryptPwd Failed")
	}
	rsp, err := l.svcCtx.UserRpc.UpdateUser(l.ctx, &user.UpdateUserReq{
		Payload: &user.User{
			Id:       req.Id,
			Email:    gRsp.Result.Email,
			Name:     gRsp.Result.Name,
			Password: newPassword,
		},
	})
	if err != nil {
		l.Logger.Errorf("Fail to Update(id: %s) on updatePassword", req.Id)
		return &types.UpdateUserReply{Result: false}, nil
	}
	return &types.UpdateUserReply{Result: rsp.Result}, nil
}
