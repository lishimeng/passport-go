package oauth

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/passport-go/internal/cache"
	"github.com/lishimeng/passport-go/internal/db/service"
	"net/url"
	"strings"
)

func authorize(ctx server.Context) {
	var resp app.Response
	clientID := ctx.C.FormValue("client_id")
	redirectURI := ctx.C.FormValue("redirect_uri")
	responseType := ctx.C.FormValue("response_type")
	scope := ctx.C.FormValue("scope")
	token := ctx.C.FormValue("token")
	//state := ctx.C.FormValue("state")
	path := ctx.C.FormValue("path")

	callback, err := url.QueryUnescape(redirectURI)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "redirect uri 解析失败"
		ctx.Json(resp)
		return
	}
	// 验证token信息
	p, err := auth.TokenStorage.Verify(token)
	if err != nil {
		// 登录失效，跳转回登录界面
		ctx.C.Redirect("/401?path="+path, iris.StatusFound)
		return
	}

	userCode := p.Uid

	if responseType == "code" {
		var c string
		// 查询可用org，需要选择则可302跳转组织选择页面
		tenants, err := service.QueryTenants(userCode, clientID)
		if err != nil || len(tenants) == 0 {
			ctx.C.Redirect("/403?path="+path, iris.StatusFound)
			return
		}
		c, err = cache.GenerateAuthCode(userCode, clientID, tenants[0], scope)
		if err != nil {
			resp.Code = tool.RespCodeError
			resp.Message = "code获取失败"
			ctx.Json(resp)
			return
		}
		if strings.Contains(callback, "?") {
			callback = callback + "&code=" + c
		} else {
			callback = callback + "?code=" + c
		}
	}

	// 307 跳转回客户端，使用POST请求方法
	ctx.C.Redirect(callback, iris.StatusTemporaryRedirect)
}
