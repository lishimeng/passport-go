package oauth

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/cache"
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/passport-go/internal/db/service"
	"github.com/lishimeng/passport-go/internal/etc"
	"github.com/lishimeng/passport-go/internal/gentoken"
	"github.com/lishimeng/passport-go/internal/passwd"
)

type PasswordReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

const passwordScope = "read,write"

func passwordAuth(ctx server.Context) {
	var req PasswordReq
	var resp AuthResponse
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	appId := ctx.C.Request().Header.Get(auth.UidKey)

	var info model.UserInfo
	err = app.GetOrm().Context.QueryTable(new(model.UserInfo)).
		Filter("UserName", req.UserName).
		One(&info)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "用户名或密码错误"
		ctx.Json(resp)
		return
	}
	p := passwd.Verify(req.Password, info)
	log.Info("password:%s,%s", info.Password, p)
	if !p {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "用户名或密码错误"
		ctx.Json(resp)
		return
	}
	// 查询可用org
	tenants, err := service.QueryTenants(info.UserCode, appId)
	if err != nil || len(tenants) == 0 {
		resp.Code = 403
		resp.Message = "对不起，您没有权限访问此应用"
		ctx.Json(resp)
		return
	}
	org := tenants[0]
	tokenContent, err := gentoken.GenOpenToken(info.UserCode, appId, org, passwordScope)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "token获取失败"
		ctx.Json(resp)
		return
	}

	gentoken.SaveToken(tokenContent)

	refreshToken, err := cache.GenerateRefreshToken(info.UserCode, appId, org, passwordScope, etc.PwdRefreshTTL)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "refresh token 获取失败"
	}
	resp.Code = 200
	resp.Message = "success"
	resp.AccessToken = string(tokenContent)
	resp.TokenType = "bearer"
	resp.ExpiresIn = int(etc.TokenTTL.Seconds())
	resp.RefreshToken = refreshToken
	resp.Scope = passwordScope
	ctx.Json(&resp)
}
