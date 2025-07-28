package cache

import (
	"encoding/base64"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/config"
	"math/rand"
	"time"
)

type Data struct {
	UserCode   string `json:"user_id"`
	ClientCode string `json:"client_id"`
	TenantCode string `json:"tenant_id"`
	Scope      string `json:"scope"`
	//State    string `json:"state"`
}

func generateCode() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes), nil
}

func saveCode(code string, ttl time.Duration, data *Data) {
	go func() {
		err := app.GetCache().SetTTL(code, data, ttl)
		if err != nil {
			log.Debug(err)
		}
	}()
}

func GenerateAuthCode(userCode string, clientCode string, tenantCode string, scope string) (string, error) {
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
	saveCode("oauth_code_"+code, config.Config.TTL.Code, data)
	return code, nil
}
