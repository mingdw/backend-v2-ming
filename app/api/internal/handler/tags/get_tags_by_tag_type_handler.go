package tags

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/tags"
	"metaLand/app/api/internal/svc"
)

// 根据类型获取标签列表
func GetTagsByTagTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tags.NewGetTagsByTagTypeLogic(r.Context(), svcCtx)
		resp, err := l.GetTagsByTagType()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
