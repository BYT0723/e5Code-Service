package logic

import (
	"context"
	"encoding/json"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/pb"

	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GetProjectAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectAuthLogic {
	return &GetProjectAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectAuthLogic) GetProjectAuth(in *pb.GetProjectReq) (*pb.GetProjectAuthRsp, error) {
	pj := &model.Project{ID: in.Id}
	if err := l.svcCtx.DB.First(pj).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codesx.NotFound, "ProjectNotFound")
		}
		logx.Error("Fail to Find Project On Get Auth: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	auth := &http.BasicAuth{}
	if err := json.Unmarshal([]byte(pj.Auth), auth); err != nil {
		logx.Error("Fail to Unmarshal Project Auth:", err.Error())
		return nil, status.Error(codesx.JSONMarshalError, err.Error())
	}

	return &pb.GetProjectAuthRsp{Username: auth.Username, Password: auth.Password}, nil
}
