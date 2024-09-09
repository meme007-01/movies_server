package home

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"movies_server/server/internal/logic/home"
	"movies_server/server/internal/svc"
	"movies_server/server/internal/types"
)

func GetPlayLineServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPlayLineRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := home.NewGetPlayLineServerLogic(r.Context(), svcCtx)
		resp, err := l.GetPlayLineServer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
