package path

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/passport-go/internal/config"
)

func Route(root server.Router) {
	root.Get("/", getPathInfo) //登录
}

type pathInfo struct {
	app.Response
	Path string `json:"path"`
}

// getPathInfo 登录成功后默认的跳转URL（如果没有URL查询参数）
// 部署时改成profile的URL
func getPathInfo(ctx server.Context) {
	var resp pathInfo
	resp.Path = config.Config.Page.DefaultRedirect
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
