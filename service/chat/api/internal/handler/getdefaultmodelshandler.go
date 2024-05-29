package handler

import (
	"net/http"

	"github.com/openui-backend-go/service/chat-api/internal/logic"
	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDefaultModelsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DefaultModels
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetDefaultModelsLogic(r.Context(), svcCtx)
		resp, err := l.GetDefaultModels(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
