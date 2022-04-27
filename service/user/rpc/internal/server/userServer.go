// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"e5Code-Service/service/user/rpc/internal/logic"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.GetUserRsp, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServer) GetUserByEmail(ctx context.Context, in *pb.GetUserByEmailReq) (*pb.GetUserRsp, error) {
	l := logic.NewGetUserByEmailLogic(ctx, s.svcCtx)
	return l.GetUserByEmail(in)
}

func (s *UserServer) AddUser(ctx context.Context, in *pb.AddUserReq) (*pb.AddUserRsp, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *UserServer) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserRsp, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UserServer) DeleteUser(ctx context.Context, in *pb.DeleteUserReq) (*pb.DeleteUserRsp, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRsp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) ListUser(ctx context.Context, in *pb.ListUserReq) (*pb.ListUserRsp, error) {
	l := logic.NewListUserLogic(ctx, s.svcCtx)
	return l.ListUser(in)
}

func (s *UserServer) GetPermission(ctx context.Context, in *pb.GetPermissionReq) (*pb.GetPermissionRsp, error) {
	l := logic.NewGetPermissionLogic(ctx, s.svcCtx)
	return l.GetPermission(in)
}

func (s *UserServer) SetPermission(ctx context.Context, in *pb.SetPermissionReq) (*pb.SetPermissionRsp, error) {
	l := logic.NewSetPermissionLogic(ctx, s.svcCtx)
	return l.SetPermission(in)
}

func (s *UserServer) DeletePermission(ctx context.Context, in *pb.DeletePermissionReq) (*pb.DeletePermissionRsp, error) {
	l := logic.NewDeletePermissionLogic(ctx, s.svcCtx)
	return l.DeletePermission(in)
}

func (s *UserServer) GetPermissions(ctx context.Context, in *pb.GetPermissionsReq) (*pb.GetPermissionsRsp, error) {
	l := logic.NewGetPermissionsLogic(ctx, s.svcCtx)
	return l.GetPermissions(in)
}
