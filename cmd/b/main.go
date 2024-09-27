package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate(buildscript.Project{
		Namespace: "lishimeng",
	},
		buildscript.Application{
			Name:    "passport-go",
			AppPath: "cmd/passport",
			HasUI:   false,
		},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
