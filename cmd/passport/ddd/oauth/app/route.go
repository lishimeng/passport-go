package app

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/midware"
)

func Route(root server.Router) {
	root.Post("/token", genCredential)
	root.Get("/token", midware.WithAuth(verify)...)
}
