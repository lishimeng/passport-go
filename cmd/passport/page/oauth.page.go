package page

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport-go/internal/db/model"
)

type openAuthModel struct {
	Model
	AppName      string
	ClientID     string
	RedirectURI  string
	ResponseType string
	Scope        string
	State        string
}

func openAuth(ctx server.Context) {
	var err error
	var data openAuthModel
	appId := ctx.C.URLParam("client_id")
	responseType := ctx.C.URLParam("response_type")
	redirectURI := ctx.C.URLParam("redirect_uri")
	scope := ctx.C.URLParam("scope")

	var appInfo model.Application
	err = app.GetOrm().Context.QueryTable(new(model.Application)).
		Filter("Code", appId).
		One(&appInfo)
	if err != nil || responseType == "" || redirectURI == "" {
		// 参数不合法，显示400页面
		ctx.C.ViewLayout("layout/main")
		err = ctx.C.View("400.html")
		if err != nil {
			_, _ = ctx.C.HTML("<h3>%s</h3>", err.Error())
		}
		return
	}
	data.Title = "passport"
	data.AppName = appInfo.Name
	data.ClientID = appId
	data.RedirectURI = redirectURI
	data.ResponseType = responseType
	data.Scope = scope
	ctx.C.ViewLayout("layout/main")
	err = ctx.C.View("openAuth.html", data)
	if err != nil {
		_, _ = ctx.C.HTML("<h3>%s</h3>", err.Error())
	}
}
