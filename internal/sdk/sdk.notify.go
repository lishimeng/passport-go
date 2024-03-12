package sdk

import (
	"github.com/lishimeng/x/container"
)

type SendSms func(code, mobile string) error
type SendEmail func(code, email string) error

func GetEmailSender() SendEmail {
	var sender SendEmail
	_ = container.Get(&sender)
	return sender
}

func GetSmsSender() SendSms {
	var sender SendSms
	_ = container.Get(&sender)
	return sender
}
