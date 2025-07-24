package oauth

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/oauth/app"
	"github.com/lishimeng/passport-go/internal/midware"
)

func Route(root server.Router) {
	app.Route(root.Path("/app"))

	root.Post("/submit", authorize)

	root.Post("/password", midware.WithAuth(passwordAuth)...)
	root.Post("/code", midware.WithAuth(codeAuth)...)
	root.Post("/refreshToken", midware.WithAuth(refreshAuth)...)
}
