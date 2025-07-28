package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/oauth"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/path"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/send"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/signin"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/signup"
	"github.com/lishimeng/passport-go/cmd/passport/page"
)

func Route(app server.Router) {
	route(app.Path("/api"))
	oauth.Route(app.Path("/open/oauth2"))

	page.Route(app)
}

func route(root server.Router) {
	signin.Route(root.Path("/signin"))
	signup.Route(root.Path("/signup"))
	send.Route(root.Path("/send"))
	path.Route(root.Path("/path"))
}
