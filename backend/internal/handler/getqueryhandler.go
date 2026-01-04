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

// 注意：goctl 重新生成时可能会覆盖此文件。如果你需要保持自定义逻辑，
// 请在重新生成后合并或将自定义部分放到单独文件中。
func GetQueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq

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

		l := logic.NewGetQueryLogic(r.Context(), svcCtx)
		resp, err := l.GetQuery(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
