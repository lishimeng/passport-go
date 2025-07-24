package oauth

import (
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/passport-go/internal/cache"
	"github.com/lishimeng/passport-go/internal/etc"
	"github.com/lishimeng/passport-go/internal/gentoken"
)

type CodeReq struct {
	Code string `json:"code"`
}

func codeAuth(ctx server.Context) {
	var req CodeReq
	var resp AuthResponse
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	appId := ctx.C.Request().Header.Get(auth.UidKey)

	cached, err := cache.ValidateAuthCode(req.Code)
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

	tokenContent, err := gentoken.GenOpenToken(cached.UserCode, appId, cached.TenantCode, cached.Scope)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "access token 获取失败"
		ctx.Json(resp)
		return
	}

	gentoken.SaveToken(tokenContent)

	refreshToken, err := cache.GenerateRefreshToken(cached.UserCode, appId, cached.TenantCode, cached.Scope, etc.CodeRefreshTTL)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "refresh token 获取失败"
		ctx.Json(resp)
		return
	}
	resp.Code = 200
	resp.Message = "success"
	resp.AccessToken = string(tokenContent)
	resp.TokenType = "bearer"
	resp.ExpiresIn = int(etc.TokenTTL.Seconds())
	resp.RefreshToken = refreshToken
	resp.Scope = cached.Scope
	ctx.Json(&resp)
}
