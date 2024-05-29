package handler

import (
	"net/http"

	"github.com/openui-backend-go/service/chat-api/internal/logic"
	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPromptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetPromptLogic(r.Context(), svcCtx)
		resp, err := l.GetPrompt()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
