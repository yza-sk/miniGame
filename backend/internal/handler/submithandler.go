// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"encoding/json"
	"net/http"

	"rank_list/internal/logic"
	"rank_list/internal/svc"
	"rank_list/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SubmitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubmitReq
		
		// 对于 POST 请求，直接从 JSON body 解码；对于 GET 或表单使用 httpx.Parse
		if r.Method == http.MethodPost {
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
		} else {
			if err := httpx.Parse(r, &req); err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
		}

		l := logic.NewSubmitLogic(r.Context(), svcCtx)
		resp, err := l.Submit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
