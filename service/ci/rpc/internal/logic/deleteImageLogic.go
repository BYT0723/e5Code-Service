package logic

import (
	"context"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/ci/model"
	"e5Code-Service/service/ci/rpc/internal/svc"
	"e5Code-Service/service/ci/rpc/pb"

	"github.com/docker/docker/api/types"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DeleteImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteImageLogic {
	return &DeleteImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteImageLogic) DeleteImage(in *pb.DeleteImageReq) (*pb.DeleteImageRsp, error) {
	image := &model.Image{}
	if err := l.svcCtx.DB.First(&image).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "ImageNotFound")
		}
		logx.Error("Fail to Find Image on DeleteImage:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	// 删除镜像
	if _, err := l.svcCtx.DockerClient.ImageRemove(l.ctx, image.ImageID, types.ImageRemoveOptions{
		Force:         false,
		PruneChildren: false,
	}); err != nil {
		logx.Error("Fail to ImageRemove on DeleteImage:", err.Error())
		return nil, status.Error(codesx.DockerError, err.Error())
	}

	// 删除db中镜像记录
	if err := l.svcCtx.DB.Delete(&model.Image{ID: in.ID}).Error; err != nil {
		logx.Error("Fail to DeleteImage:", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &pb.DeleteImageRsp{}, nil
}
