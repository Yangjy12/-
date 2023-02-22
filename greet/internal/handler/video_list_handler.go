package handler

import (
	"net/http"

	"Bytecode_Project/greet/internal/logic"
	"Bytecode_Project/greet/internal/svc"
	"Bytecode_Project/greet/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVideoListLogic(r.Context(), svcCtx)
		resp, err := l.VideoList(&req,r.Header.Get("id"),r.Header.Get("name"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
