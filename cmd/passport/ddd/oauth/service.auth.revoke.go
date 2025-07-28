package oauth

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport-go/internal/cache"
)

type RevokeReq struct {
	RefreshToken string `json:"refresh_token"`
}

func revokeAuth(ctx server.Context) {
	var req RevokeReq
	var resp app.Response
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	appId := ctx.C.Request().Header.Get(auth.UidKey)
	ctx.C.Params()

	cached, err := cache.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		resp.Code = 404
		resp.Message = "unknown code"
		ctx.Json(resp)
		return
	}
	if cached.ClientCode != appId {
		resp.Code = 403
		resp.Message = "invalid code"
		ctx.Json(resp)
		return
	}

	err = cache.DeleteRefreshToken(req.RefreshToken)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = 200
	resp.Message = "success"
	ctx.Json(resp)
}
