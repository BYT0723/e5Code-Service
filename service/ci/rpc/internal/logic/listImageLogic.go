package logic

import (
	"context"
	"time"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"
	userModel "e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/user"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ListImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListImageLogic {
	return &ListImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListImageLogic) ListImage(in *pb.ListImageReq) (*pb.ListImageRsp, error) {
	images := []*model.Image{}
	if err := l.svcCtx.DB.Find(&images, "project_id = ?", in.ProjectID).Error; err != nil {
		logx.Error("Fail to Find Images:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	ids := []string{}
	for _, im := range images {
		ids = append(ids, im.BuilderID)
	}
	rsp, err := l.svcCtx.UserRpc.ListUser(l.ctx, &user.ListUserReq{
		Ids: ids,
	})
	if err != nil {
		logx.Error("Fail to ListUser on ListImage:", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}
	users := make(map[string]*userModel.User)
	for _, user := range rsp.Result {
		users[user.ID] = &userModel.User{
			ID:      user.ID,
			Email:   user.Email,
			Account: user.Account,
			Name:    user.Name,
			Bio:     user.Bio,
		}
	}
	for _, im := range images {
		im.Builder = users[im.BuilderID]
	}

	imageModels := []*pb.Image{}
	if err := copier.CopyWithOption(&imageModels, &images, copier.Option{
		DeepCopy: true,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: &timestamp.Timestamp{},
				Fn: func(src interface{}) (interface{}, error) {
					temp, _ := src.(time.Time)
					return timestamppb.New(temp), nil
				},
			},
		},
	}); err != nil {
		logx.Error("Fail to CopyWithOption on ListImage:", err.Error())
		return nil, status.Error(codesx.CopierError, err.Error())
	}
	return &pb.ListImageRsp{
		Count:  int64(len(imageModels)),
		Result: imageModels,
	}, nil
}
