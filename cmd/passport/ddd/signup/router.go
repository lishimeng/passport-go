package signup

import (
	"github.com/lishimeng/app-starter/server"
)

func Route(root server.Router) {
	root.Post("/wx", wechatRegister)
}
