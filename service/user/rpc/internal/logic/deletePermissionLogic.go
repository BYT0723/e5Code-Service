package logic

import (
	"context"
	"fmt"
	"strings"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DeletePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermissionLogic {
	return &DeletePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePermissionLogic) DeletePermission(in *pb.DeletePermissionReq) (*pb.DeletePermissionRsp, error) {
	condition := []string{}
	if in.UserID != "" {
		condition = append(condition, fmt.Sprintf("user_id = '%s'", in.UserID))
	}
	if in.ProjectID != "" {
		condition = append(condition, fmt.Sprintf("project_id = '%s'", in.ProjectID))
	}
	if len(condition) == 0 {
		return nil, status.Error(codesx.APIError, "Args'len be required > 0")
	}
	if err := l.svcCtx.Db.Where(strings.Join(condition, " and ")).Delete(&model.Permission{}).Error; err != nil {
		logx.Error("Fail to deletePermission: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	return &pb.DeletePermissionRsp{}, nil
}
