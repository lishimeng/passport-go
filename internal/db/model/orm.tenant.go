package model

import "github.com/lishimeng/app-starter"

type Tenant struct {
	app.Pk
	Code string `orm:"column(code);unique"`
	Name string `orm:"column(name)"`
	app.TableChangeInfo
}
