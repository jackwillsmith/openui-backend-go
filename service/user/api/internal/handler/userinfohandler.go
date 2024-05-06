package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"openui-backend-go/service/user/api/internal/logic"
	"openui-backend-go/service/user/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
