package signup

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/db/model"
)

func svsAddWxUser(uid, phone, unionId string) (err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {

		var user = model.UserInfo{
			UserCode: uid,
			Phone:    phone,
		}
		user.Status = app.Enable

		_, e = ctx.Context.Insert(&user)
		if e != nil {
			log.Info("create user fail")
			log.Info(e)
			return
		}

		var mfa = model.Mfa{
			MfaCode:           uid,
			MfaType:           model.MfaUnknown, // 验证方式需要显式打开
			SecretPhoneNumber: phone,
		}
		mfa.Status = app.Enable
		_, e = ctx.Context.Insert(&mfa)
		if e != nil {
			log.Info("create mfa fail")
			log.Info(e)
			return
		}

		var wxMfaItem = model.MfaItem{
			MfaCode:  uid,
			Sn:       unionId,
			Category: model.MfaWechat,
		}
		wxMfaItem.Status = app.Enable
		_, e = ctx.Context.Insert(&wxMfaItem)
		if e != nil {
			log.Info("create mfa item(wechat) fail")
			log.Info(e)
			return
		}

		var phoneMfaItem = model.MfaItem{
			MfaCode:  uid,
			Sn:       phone,
			Category: model.MfaPhoneNumber,
		}
		phoneMfaItem.Status = app.Enable
		_, e = ctx.Context.Insert(&phoneMfaItem)
		if e != nil {
			log.Info("create mfa item(phone_number) fail")
			log.Info(e)
			return
		}

		return
	})
	return
}
