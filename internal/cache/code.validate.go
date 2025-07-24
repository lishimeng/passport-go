package cache

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"time"
)

func ValidateAuthCode(code string) (data Data, err error) {
	key := "oauth_code_" + code
	err = app.GetCache().Get(key, &data)
	if err != nil {
		return
	}
	err = app.GetCache().Del(key)
	return
}

func TestCode() error {
	code, err := GenerateAuthCode("user", "client", "tenant", "test")
	if err != nil {
		return err
	}
	time.Sleep(30 * time.Second)
	data, err := ValidateAuthCode(code)
	if err != nil {
		log.Debug("Timeout")
		return err
	}
	log.Debug("Not Timeout")
	log.Debug(data)
	data2, err := ValidateAuthCode(code)
	if err != nil {
		return nil
	}
	log.Debug("Not Deleted")
	log.Debug(data2)
	return nil
}
