package handler

import (
	"net/http"

	"github.com/openui-backend-go/service/user-api/internal/logic"
	"github.com/openui-backend-go/service/user-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewConfigLogic(r.Context(), svcCtx)
		resp, err := l.Config()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
