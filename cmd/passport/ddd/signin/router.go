package signin

import (
	"github.com/lishimeng/app-starter/server"
)

func Route(root server.Router) {
	root.Post("/", passwdSignIn) // 默认
	root.Post("/wx", wechatSignIn)
	root.Post("/sms", smsSignIn)
}
