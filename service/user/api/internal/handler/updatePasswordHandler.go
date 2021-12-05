package handler

import (
	"e5Code-Service/service/user/api/internal/logic"
	"e5Code-Service/service/user/api/internal/svc"
	"e5Code-Service/service/user/api/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func updatePasswordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdatePasswordLogic(r.Context(), ctx)
		resp, err := l.UpdatePassword(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
