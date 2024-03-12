package sdk

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/go-sdk/tianditu"
	"github.com/lishimeng/x/container"
)

func CreateTianditu(config TiandituConfig) {
	log.Info("init tianditu sdk")
	tClient := tianditu.NewClient(tianditu.WithKey(config.Key))
	container.Add(&tClient)
}
