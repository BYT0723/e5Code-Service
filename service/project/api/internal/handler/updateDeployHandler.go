package handler

import (
	"net/http"

	"e5Code-Service/service/project/api/internal/logic"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func updateDeployHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateDeployReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateDeployLogic(r.Context(), ctx)
		resp, err := l.UpdateDeploy(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
