package handler

import (
	"net/http"

	"github.com/openui-backend-go/service/chat-api/internal/logic"
	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChangelogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewChangelogLogic(r.Context(), svcCtx)
		resp, err := l.Changelog()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
