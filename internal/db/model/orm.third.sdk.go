package model

import "github.com/lishimeng/app-starter"

type SdkConfig struct {
	app.Pk
	Wechat   string `orm:"column(wechat)"`   // 微信
	Tianditu string `orm:"column(tianditu)"` // 天地图
	app.TableInfo
}
