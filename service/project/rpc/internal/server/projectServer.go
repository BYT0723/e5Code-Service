// Code generated by goctl. DO NOT EDIT!
// Source: project.proto

package server

import (
	"context"

	"e5Code-Service/service/project/rpc/internal/logic"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/pb"
)

type ProjectServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProjectServer
}

func NewProjectServer(svcCtx *svc.ServiceContext) *ProjectServer {
	return &ProjectServer{
		svcCtx: svcCtx,
	}
}

//  basic operation
func (s *ProjectServer) GetProject(ctx context.Context, in *pb.GetProjectReq) (*pb.GetProjectRsp, error) {
	l := logic.NewGetProjectLogic(ctx, s.svcCtx)
	return l.GetProject(in)
}

func (s *ProjectServer) AddProject(ctx context.Context, in *pb.AddProjectReq) (*pb.AddProjectRsp, error) {
	l := logic.NewAddProjectLogic(ctx, s.svcCtx)
	return l.AddProject(in)
}

func (s *ProjectServer) UpdateProject(ctx context.Context, in *pb.UpdateProjectReq) (*pb.UpdateProjectRsp, error) {
	l := logic.NewUpdateProjectLogic(ctx, s.svcCtx)
	return l.UpdateProject(in)
}

func (s *ProjectServer) DeleteProject(ctx context.Context, in *pb.DeleteProjectReq) (*pb.DeleteProjectRsp, error) {
	l := logic.NewDeleteProjectLogic(ctx, s.svcCtx)
	return l.DeleteProject(in)
}

func (s *ProjectServer) ListProject(ctx context.Context, in *pb.ListProjectReq) (*pb.ListProjectRsp, error) {
	l := logic.NewListProjectLogic(ctx, s.svcCtx)
	return l.ListProject(in)
}

//  permission manager
func (s *ProjectServer) AddUser(ctx context.Context, in *pb.AddUserReq) (*pb.AddUserRsp, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *ProjectServer) RemoveUser(ctx context.Context, in *pb.RemoveUserReq) (*pb.RemoveUserRsp, error) {
	l := logic.NewRemoveUserLogic(ctx, s.svcCtx)
	return l.RemoveUser(in)
}

func (s *ProjectServer) ModifyPermission(ctx context.Context, in *pb.ModifyPermissionReq) (*pb.ModifyPermissionRsp, error) {
	l := logic.NewModifyPermissionLogic(ctx, s.svcCtx)
	return l.ModifyPermission(in)
}
