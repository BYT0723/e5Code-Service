// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"e5Code-Service/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/updateUser",
				Handler: updateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/deleteUser",
				Handler: deleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/userInfo",
				Handler: userInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/userInfoByEmail",
				Handler: userInfoByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/listUser",
				Handler: listUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/getPermission",
				Handler: getPermissionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/setPermission",
				Handler: setPermissionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/deletePermission",
				Handler: deletePermissionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/getPermissions",
				Handler: getPermissionsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/registerUser",
				Handler: registerUserHandler(serverCtx),
			},
		},
	)
}
