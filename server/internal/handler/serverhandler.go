package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"movies/server/internal/logic"
	"movies/server/internal/svc"
	"movies/server/internal/types"
)

func ServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewServerLogic(r.Context(), svcCtx)
		resp, err := l.Server(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
