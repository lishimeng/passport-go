package cache

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"time"
)

func GenerateRefreshToken(userCode string, clientCode string, tenantCode string, scope string, ttl time.Duration) (string, error) {
	code, err := generateCode()
	if err != nil {
		return "", err
	}
	data := &Data{
		UserCode:   userCode,
		ClientCode: clientCode,
		TenantCode: tenantCode,
		Scope:      scope,
	}
	saveCode("oauth_refresh_token_"+code, ttl, data)
	return code, nil
}

func ValidateRefreshToken(refreshToken string) (data Data, err error) {
	key := "oauth_refresh_token_" + refreshToken
	err = app.GetCache().Get(key, &data)
	// 允许复用，不需要Del
	return
}

func DeleteRefreshToken(refreshToken string) error {
	key := "oauth_refresh_token_" + refreshToken
	log.Debug("delete refresh token")
	return app.GetCache().Del(key)
}
