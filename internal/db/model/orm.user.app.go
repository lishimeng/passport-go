package model

import "github.com/lishimeng/app-starter"

type UserAppAuth struct {
	app.Pk
	UserCode   string `orm:"column(user_code)"`
	AppCode    string `orm:"column(app_code)"`
	TenantCode string `orm:"column(tenant_code)"`
	app.TableChangeInfo
}
