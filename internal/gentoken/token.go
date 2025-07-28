package gentoken

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/common"
	"github.com/lishimeng/passport-go/internal/config"
	"github.com/lishimeng/passport-go/internal/etc"
	"github.com/lishimeng/x/container"
	"time"
)

func GenToken(uCode, loginType string) (tokenVal []byte, err error) {
	var tokenPayload token.JwtPayload
	tokenPayload.Uid = uCode
	tokenPayload.Client = loginType
	tokenPayload.Scope = common.Scope
	tokenVal, err = _genWithTTL(tokenPayload, config.Config.TTL.AccessToken)
	return
}

// GenOpenToken SaaS服务，供第三方使用
func GenOpenToken(userCode string, appID string, org string, scope string) (tokenVal []byte, err error) {
	var tokenPayload token.JwtPayload
	tokenPayload.Uid = userCode
	tokenPayload.Client = appID
	tokenPayload.Org = org
	tokenPayload.Scope = scope
	// 所有access token使用统一ttl，无需传入参数
	tokenVal, err = _genWithTTL(tokenPayload, config.Config.TTL.AccessToken)
	return
}

func _genToken(payload token.JwtPayload) (content []byte, err error) {
	var provider token.JwtProvider
	err = container.Get(&provider)
	if err != nil {
		return
	}
	log.Info("tokenPayload: %s", payload)
	content, err = provider.Gen(payload)
	return
}

func _genWithTTL(payload token.JwtPayload, ttl time.Duration) (content []byte, err error) {
	var provider token.JwtProvider
	err = container.Get(&provider)
	if err != nil {
		return
	}
	log.Info("tokenPayload: %s", payload)
	content, err = provider.GenWithTTL(payload, ttl)
	return
}

func SaveToken(tokenContent []byte) {
	// 只有启动了 EnableJwtTokenCache 才需要 SaveToken
	if etc.Config.Token.EnableJwtTokenCache {
		go func() {
			key := token.Digest(tokenContent)
			log.Info("缓存tokenContent：%s,%s", key, config.Config.TTL.AccessToken)
			err := app.GetCache().SetTTL(key, string(tokenContent), config.Config.TTL.AccessToken)
			if err != nil {
				log.Info(err)
			}
		}()
	}
	return
}
