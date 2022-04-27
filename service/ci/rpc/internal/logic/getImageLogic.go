package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"
	"e5Code-Service/service/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImageLogic {
	return &GetImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetImageLogic) GetImage(in *pb.GetImageReq) (*pb.GetImageRsp, error) {
	image := &model.Image{ID: in.ID}
	if err := l.svcCtx.DB.First(image).Error; err != nil {
		logx.Error("Fail to GetImage: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	rsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: image.BuilderID})
	if err != nil {
		logx.Error("Fail to GetUser on GetImage:", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	imageModel := &pb.Image{}
	if err := copier.Copy(&imageModel, &image); err != nil {
		logx.Error("Fail to Copy on GetImage:", err.Error())
		return nil, status.Error(codesx.CopierError, err.Error())
	}
	imageModel.Builder = &pb.UserModel{
		ID:      rsp.Id,
		Email:   rsp.Email,
		Account: rsp.Account,
		Name:    rsp.Name,
		Bio:     rsp.Bio,
	}

	return &pb.GetImageRsp{Result: imageModel}, nil
}
