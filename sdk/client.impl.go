package sdk

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/utils"
)

const ApiPasswordAuth = "/open/oauth2/password"

const ApiCredential = "/open/oauth2/app/token"

const (
	CodeNotAllow int = 401
	CodeNotFound int = 404
	CodeSuccess  int = 200
)

type passportClient struct {
	host   string // 主机地址. 如, "http://127.0.0.1"
	appId  string
	secret string
}

func (p *passportClient) PasswordAuth(request PasswordRequest) (response AuthResponse, err error) {
	if debugEnable {
		log.Debug("PasswordAuth: %s", request.UserName)
	}

	err = NewRpc(p.host).Auth(p.appId, p.secret).BuildReq(func(rest *utils.RestClient) (int, error) {
		code, e := rest.Path(ApiPasswordAuth).ResponseJson(&response).Post(request)
		if response.Code == CodeNotAllow { // 拦截业务异常
			code = CodeNotAllow
		}
		return code, e
	}).Exec()

	if err != nil {
		log.Debug(err)
		return
	}
	return
}
