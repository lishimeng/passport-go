package model

import "github.com/lishimeng/app-starter"

type Mfa struct { // 特定设计, 应减少批量检索MFA item
	app.Pk
	MfaCode           string      `orm:"column(mfa_code)"`
	MfaType           MfaCategory `orm:"column(mfa_type)"`
	SecretPhoneNumber string      `orm:"column(secret_phone_number)"`
	SecretEmail       string      `orm:"column(secret_email)"`
	app.TableChangeInfo
}

type MfaCategory string

const (
	MfaUnknown     MfaCategory = "unknown"                 // 未指定
	MfaPhoneNumber MfaCategory = "phone_number"            // 手机号
	MfaEmail       MfaCategory = "email"                   //  邮箱
	MfaWechat      MfaCategory = "wechat"                  // 微信(union_id优先)
	MfaGoogle      MfaCategory = "google_authenticator"    // google验证器
	MfaMicrosoft   MfaCategory = "microsoft_authenticator" // 微软验证器
)

type MfaItem struct {
	app.Pk
	MfaCode  string      `orm:"column(mfa_code)"`     // MFA编号
	Sn       string      `orm:"column(sn)"`           // 验证编号
	Category MfaCategory `orm:"column(mfa_category)"` // 验证类型
	app.TableChangeInfo
}

// MfaDevice 受信任设备列表
type MfaDevice struct {
	app.Pk
	MfaCode string `orm:"column(mfa_code)"` // MFA编号
	Key     string `orm:"column(key)"`      // MFA设备
	app.TableInfo
}
