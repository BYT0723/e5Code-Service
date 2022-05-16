// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"e5Code-Service/api/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddSSHKeyReq        = user.AddSSHKeyReq
	AddSSHKeyRsp        = user.AddSSHKeyRsp
	AddUserReq          = user.AddUserReq
	AddUserRsp          = user.AddUserRsp
	DeletePermissionReq = user.DeletePermissionReq
	DeletePermissionRsp = user.DeletePermissionRsp
	DeleteSSHKeyReq     = user.DeleteSSHKeyReq
	DeleteSSHKeyRsp     = user.DeleteSSHKeyRsp
	DeleteUserReq       = user.DeleteUserReq
	DeleteUserRsp       = user.DeleteUserRsp
	GetPermissionReq    = user.GetPermissionReq
	GetPermissionRsp    = user.GetPermissionRsp
	GetPermissionsReq   = user.GetPermissionsReq
	GetPermissionsRsp   = user.GetPermissionsRsp
	GetUserByEmailReq   = user.GetUserByEmailReq
	GetUserReq          = user.GetUserReq
	GetUserRsp          = user.GetUserRsp
	ListSSHKeysReq      = user.ListSSHKeysReq
	ListSSHKeysRsp      = user.ListSSHKeysRsp
	ListUserReq         = user.ListUserReq
	ListUserRsp         = user.ListUserRsp
	LoginReq            = user.LoginReq
	LoginRsp            = user.LoginRsp
	PermissionInfo      = user.PermissionInfo
	SSHKey              = user.SSHKey
	SetPermissionReq    = user.SetPermissionReq
	SetPermissionRsp    = user.SetPermissionRsp
	UpdateUserReq       = user.UpdateUserReq
	UpdateUserRsp       = user.UpdateUserRsp
	UserModel           = user.UserModel

	User interface {
		GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRsp, error)
		GetUserByEmail(ctx context.Context, in *GetUserByEmailReq, opts ...grpc.CallOption) (*GetUserRsp, error)
		AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserRsp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRsp, error)
		DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRsp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRsp, error)
		ListUser(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserRsp, error)
		GetPermission(ctx context.Context, in *GetPermissionReq, opts ...grpc.CallOption) (*GetPermissionRsp, error)
		SetPermission(ctx context.Context, in *SetPermissionReq, opts ...grpc.CallOption) (*SetPermissionRsp, error)
		DeletePermission(ctx context.Context, in *DeletePermissionReq, opts ...grpc.CallOption) (*DeletePermissionRsp, error)
		GetPermissions(ctx context.Context, in *GetPermissionsReq, opts ...grpc.CallOption) (*GetPermissionsRsp, error)
		AddSSHKey(ctx context.Context, in *AddSSHKeyReq, opts ...grpc.CallOption) (*AddSSHKeyRsp, error)
		DeleteSSHKey(ctx context.Context, in *DeleteSSHKeyReq, opts ...grpc.CallOption) (*DeleteSSHKeyRsp, error)
		ListSSHKey(ctx context.Context, in *ListSSHKeysReq, opts ...grpc.CallOption) (*ListSSHKeysRsp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) GetUserByEmail(ctx context.Context, in *GetUserByEmailReq, opts ...grpc.CallOption) (*GetUserRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserByEmail(ctx, in, opts...)
}

func (m *defaultUser) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUser) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) ListUser(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.ListUser(ctx, in, opts...)
}

func (m *defaultUser) GetPermission(ctx context.Context, in *GetPermissionReq, opts ...grpc.CallOption) (*GetPermissionRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetPermission(ctx, in, opts...)
}

func (m *defaultUser) SetPermission(ctx context.Context, in *SetPermissionReq, opts ...grpc.CallOption) (*SetPermissionRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SetPermission(ctx, in, opts...)
}

func (m *defaultUser) DeletePermission(ctx context.Context, in *DeletePermissionReq, opts ...grpc.CallOption) (*DeletePermissionRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeletePermission(ctx, in, opts...)
}

func (m *defaultUser) GetPermissions(ctx context.Context, in *GetPermissionsReq, opts ...grpc.CallOption) (*GetPermissionsRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetPermissions(ctx, in, opts...)
}

func (m *defaultUser) AddSSHKey(ctx context.Context, in *AddSSHKeyReq, opts ...grpc.CallOption) (*AddSSHKeyRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddSSHKey(ctx, in, opts...)
}

func (m *defaultUser) DeleteSSHKey(ctx context.Context, in *DeleteSSHKeyReq, opts ...grpc.CallOption) (*DeleteSSHKeyRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteSSHKey(ctx, in, opts...)
}

func (m *defaultUser) ListSSHKey(ctx context.Context, in *ListSSHKeysReq, opts ...grpc.CallOption) (*ListSSHKeysRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.ListSSHKey(ctx, in, opts...)
}
