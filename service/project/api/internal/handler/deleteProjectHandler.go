package handler

import (
	"net/http"

	"e5Code-Service/service/project/api/internal/logic"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func deleteProjectHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteProjectLogic(r.Context(), ctx)
		resp, err := l.DeleteProject(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
