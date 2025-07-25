package signin

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport-go/internal/midware"
)

func Route(root server.Router) {
	root.Post("/", passwdSignIn) // 默认
	root.Post("/wx", wechatSignIn)
	root.Post("/sms", smsSignIn)
	root.Post("/checkToken", midware.WithAuth(checkToken)...)
}
