package config

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/db/service"
	"time"
)

// setup 调用
func useDefaultConfig() {
	Config.TTL.AccessToken = time.Hour * 2
	Config.TTL.Code = time.Minute * 5
	Config.TTL.PwdRefreshToken = time.Hour * 24 * 14
	Config.TTL.CodeRefreshToken = time.Hour * 24 * 30
	Config.Page.DefaultRedirect = "http://localhost:80/404/"
	Config.Page.EnableSMS = false
}

func saveConfig() (err error) {
	err = service.SaveConfig(CodeTTL, Config.TTL)
	if err != nil {
		return
	}
	err = service.SaveConfig(CodeFront, Config.Page)
	return
}

func _loadConfig() (err error) {
	err = service.LoadConfig(CodeTTL, &Config.TTL)
	if err != nil {
		return
	}
	err = service.LoadConfig(CodeFront, &Config.Page)
	return
}

func LoadFromDB() {
	if err := _loadConfig(); err != nil {
		log.Info("No config in DB")
		useDefaultConfig()
	}
}
