package user

import (
	"net/http"

	"github.com/fleezesd/zero-practice/app/usercenter/cmd/api/internal/logic/user"
	"github.com/fleezesd/zero-practice/app/usercenter/cmd/api/internal/svc"
	"github.com/fleezesd/zero-practice/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WXMiniAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WXMiniAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewWXMiniAuthLogic(r.Context(), svcCtx)
		resp, err := l.WXMiniAuth(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
