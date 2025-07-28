package config

import "time"

const (
	CodeTTL   = "ttl_config"
	CodeFront = "page_config"
)

var Config Configuration

type Configuration struct {
	TTL  TTLConfig
	Page PageConfig
}

type TTLConfig struct {
	AccessToken      time.Duration `json:"access_token"`
	Code             time.Duration `json:"code"`
	PwdRefreshToken  time.Duration `json:"pwd_refresh_token"`
	CodeRefreshToken time.Duration `json:"code_refresh_token"`
}

type PageConfig struct {
	EnableSMS       bool
	DefaultRedirect string
}
