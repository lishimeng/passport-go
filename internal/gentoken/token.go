package gentoken

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/etc"
	"github.com/lishimeng/x/container"
)

func GenToken(uCode, loginType string) (tokenVal []byte, err error) {
	var tokenPayload token.JwtPayload
	tokenPayload.Uid = uCode
	tokenPayload.Client = loginType
	tokenVal, err = _genToken(tokenPayload)
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

func SaveToken(tokenContent []byte) (err error) {
	key := token.Digest(tokenContent)
	log.Info("缓存tokenContent：%s,%s", key, etc.TokenTTL)
	err = app.GetCache().SetTTL(key, string(tokenContent), etc.TokenTTL)
	return
}
