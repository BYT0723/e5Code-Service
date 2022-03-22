package handler

import (
	"net/http"

	"e5Code-Service/service/user/api/internal/logic"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func deletePermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeletePermissionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeletePermissionLogic(r.Context(), svcCtx)
		resp, err := l.DeletePermission(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
