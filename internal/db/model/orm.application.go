package model

import "github.com/lishimeng/app-starter"

type Application struct {
	app.Pk
	Code        string `orm:"column(code);unique"`
	Name        string `orm:"column(name)"`
	Description string `orm:"column(description)"`
	Secret      string `orm:"column(secret)"`
	app.TableChangeInfo
}
