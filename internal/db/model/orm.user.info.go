package model

import "github.com/lishimeng/app-starter"

// UserInfo 普通用户
type UserInfo struct {
	app.Pk
	UserCode string `orm:"column(user_code)"`
	Username string `orm:"username"`
	Password string `orm:"password"`
	Email    string `orm:"email"`
	Phone    string `orm:"phone"`
	app.TableChangeInfo
	// 省略其他字段...
}

// EnterpriseUserInfo 企业用户(商户)
type EnterpriseUserInfo struct {
	app.TenantPk
	UserCode string `orm:"column(user_code)"`
	Name     string `orm:"column(name)"`
	app.TableChangeInfo
}

type NotifyType string

const (
	SmsNotifyType  NotifyType = "sms"  //短信验证
	MailNotifyType NotifyType = "mail" //邮箱验证
)
