package handler

import (
	"net/http"

	"e5Code-Service/service/user/api/internal/logic"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func listSSHKeysHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListSSHKeysReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListSSHKeysLogic(r.Context(), svcCtx)
		resp, err := l.ListSSHKeys(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
