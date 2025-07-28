package setup

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/template"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/cmd/passport/static"
	cfg "github.com/lishimeng/passport-go/internal/config"
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/passport-go/internal/sdk"
	"github.com/lishimeng/passport-go/internal/store"
	"github.com/lishimeng/x/container"
)

func Setup(_ context.Context) (err error) {
	initWebServer()
	initStoreManager()
	err = initSdkClient()
	if err != nil {
		return
	}
	err = initNotifySdk()
	if err != nil {
		return
	}
	return
}

func initStoreManager() {
	m := store.NewStoreManager()
	container.Add(&m)
}

func initWebServer() {
	AppProxy := app.GetWebServer().GetApplication()
	if AppProxy == nil {
		log.Info("web server nil")
		return
	} else {
		log.Info("web server start", AppProxy.String())
		engine := iris.HTML(static.Static, ".html")
		template.Init(engine)
		AppProxy.RegisterView(engine)
	}
}

func initSdkClient() (err error) {
	cfg.LoadFromDB()

	var config model.SdkConfig
	err = app.GetOrm().Context.QueryTable(new(model.SdkConfig)).One(&config)
	if err != nil {
		return
	}

	// 初始化微信sdk客户端
	var wx = config.Wechat
	bs, err := base64.StdEncoding.DecodeString(wx)
	if err != nil {
		return
	}
	var wxConfig sdk.WechatConfig
	err = json.Unmarshal(bs, &wxConfig)
	if err != nil {
		return
	}
	sdk.CreateWechatClient(wxConfig)

	// 初始化天地图sdk客户端
	var t = config.Tianditu
	bs, err = base64.StdEncoding.DecodeString(t)
	if err != nil {
		return
	}
	var tdConfig sdk.TiandituConfig
	err = json.Unmarshal(bs, &tdConfig)
	if err != nil {
		return
	}
	sdk.CreateTianditu(tdConfig)
	return
}

func initNotifySdk() (err error) {
	// TODO 初始化notify sdk
	return
}
