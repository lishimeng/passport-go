package model

import "github.com/lishimeng/app-starter"

type VerifyFlag int

const (
	Verified   = 1
	UnVerified = 0
)

// UserProfile 用户档案,每个用户必须有
type UserProfile struct {
	app.Pk
	UserCode       string     `orm:"column(user_code)"`        // 用户编号
	RealName       string     `orm:"column(real_name)"`        // 真实姓名
	IdCard         string     `orm:"column(id_card)"`          // 身份证号
	IdCardVerified VerifyFlag `orm:"column(id_card_verified)"` // 身份证号验证标记
	app.TableChangeInfo
}
