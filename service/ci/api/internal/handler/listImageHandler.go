package handler

import (
	"net/http"

	"e5Code-Service/service/ci/api/internal/logic"
	"e5Code-Service/service/ci/api/internal/svc"
	"e5Code-Service/service/ci/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func listImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListImageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListImageLogic(r.Context(), svcCtx)
		resp, err := l.ListImage(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
