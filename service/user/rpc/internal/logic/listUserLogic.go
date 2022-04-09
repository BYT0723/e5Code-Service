package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ListUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserLogic) ListUser(in *pb.ListUserReq) (*pb.ListUserRsp, error) {
	us := []model.User{}
	if err := l.svcCtx.Db.Model(&model.User{}).Where(in.Filter).Find(&us).Error; err != nil {
		logx.Error("Fail to Find User: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	res := []*pb.UserModel{}
	for _, v := range us {
		res = append(res, &pb.UserModel{
			Id:        v.ID,
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
			Email:     v.Email,
			Account:   v.Accout,
			Name:      v.Name,
			Bio:       v.Bio,
		})
	}

	return &pb.ListUserRsp{
		Count:  int64(len(us)),
		Result: res,
	}, nil
}
