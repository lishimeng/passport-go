package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/signin"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/signup"
)

func Route(app server.Router) {
	route(app.Path("/api"))
}

func route(root server.Router) {
	signin.Route(root.Path("/sign_in"))
	signup.Route(root.Path("/signup"))
}
