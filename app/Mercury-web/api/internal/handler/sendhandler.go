package handler

import (
	"Mercury/common/result"
	"net/http"

	"Mercury/app/Mercury-web/api/internal/logic"
	"Mercury/app/Mercury-web/api/internal/svc"
	"Mercury/app/Mercury-web/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendLogic(r.Context(), svcCtx)
		resp, err := l.Send(req)
		result.HttpResult(r, w, resp, err)
	}
}
