package comer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/comer"
	"metaLand/app/api/internal/svc"
)

// 更新用户语言
func UpdateComerLanguagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := comer.NewUpdateComerLanguagesLogic(r.Context(), svcCtx)
		resp, err := l.UpdateComerLanguages()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
