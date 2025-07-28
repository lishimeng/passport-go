package page

import "github.com/lishimeng/app-starter/server"

func error401(ctx server.Context) {
	data := HtmlError{
		Header:     "401 Unauthorized",
		Info:       "您的登录信息已失效，请重新登录。",
		Title:      "401 Unauthorized",
		ShowButton: true,
		Button:     "确定",
	}
	path := ctx.C.URLParam("path")
	data.Path = checkParams(path)
	ctx.C.ViewLayout("layout/main")
	err := ctx.C.View("error.html", data)
	if err != nil {
		_, _ = ctx.C.HTML("<h3>%s</h3>", err.Error())
	}
	return
}

func error403(ctx server.Context) {
	data := HtmlError{
		Header:     "403 Forbidden",
		Info:       "对不起，您没有权限访问此应用。",
		Title:      "403 Forbidden",
		ShowButton: true,
		Button:     "退出登录",
	}
	path := ctx.C.URLParam("path")
	data.Path = checkParams(path)
	ctx.C.ViewLayout("layout/main")
	err := ctx.C.View("error.html", data)
	if err != nil {
		_, _ = ctx.C.HTML("<h3>%s</h3>", err.Error())
	}
	return
}

func error404(ctx server.Context) {
	data := HtmlError{
		Header: "404 Not Found",
		Info:   "对不起，您访问的页面不存在。",
		Title:  "404 Not Found",
	}
	ctx.C.ViewLayout("layout/main")
	err := ctx.C.View("error.html", data)
	if err != nil {
		_, _ = ctx.C.HTML("<h3>%s</h3>", err.Error())
	}
	return
}
