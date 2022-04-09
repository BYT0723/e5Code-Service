package logic

import (
	"context"
	"fmt"
	"strings"

	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/model"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/pb"
	userModel "e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProjectLogic {
	return &ListProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProjectLogic) ListProject(in *pb.ListProjectReq) (*pb.ListProjectRsp, error) {
	pms := []userModel.Permission{}
	if err := l.svcCtx.DB.Model(&userModel.Permission{}).Where(in.Filter).Find(&pms).Error; err != nil {
		logx.Error("Fail to Find Permission: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	if len(pms) == 0 {
		return &pb.ListProjectRsp{
			Count:  0,
			Result: nil,
		}, nil
	}

	pids := make([]string, len(pms))
	for i, v := range pms {
		pids[i] = "'" + v.ProjectID + "'"
	}

	ps := []model.Project{}
	if err := l.svcCtx.DB.Model(&model.Project{}).Where(
		fmt.Sprintf("id in (%s)", strings.Join(pids, ",")),
	).Find(&ps).Error; err != nil {
		logx.Error("Fail to Find Project: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}

	uids := make([]string, len(ps))
	for i, v := range ps {
		uids[i] = "'" + v.OwnerId + "'"
	}

	rsp, err := l.svcCtx.UserRpc.ListUser(l.ctx, &user.ListUserReq{
		Filter: fmt.Sprintf("id in (%s)", strings.Join(uids, ",")),
	})
	if err != nil {
		logx.Error("Fail to ListUser on ListProject: ", err.Error())
		return nil, status.Error(codesx.RPCError, err.Error())
	}

	userMap := make(map[string]*pb.UserModel)
	for _, v := range rsp.Result {
		userMap[v.Id] = &pb.UserModel{
			Id:      v.Id,
			Email:   v.Email,
			Account: v.Account,
			Name:    v.Name,
			Bio:     v.Bio,
		}
	}
	res := []*pb.ProjectModel{}
	for _, v := range ps {
		res = append(res, &pb.ProjectModel{
			Id:      v.ID,
			Name:    v.Name,
			Url:     v.Url,
			Desc:    v.Desc,
			Status:  v.Status,
			OwnerID: v.OwnerId,
			Owner:   userMap[v.OwnerId],
		})
	}

	return &pb.ListProjectRsp{
		Count:  int64(len(res)),
		Result: res,
	}, nil
}
