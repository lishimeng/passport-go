package sdk

import (
	"errors"
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/utils"
	"net/url"
	"strconv"
)

const ApiOAuthPrefix = "/open/oauth2"

const ApiCredential = "/open/oauth2/app/token"

const PathOAuth = "/oauth"

const PathRevoke = "/open/oauth2/revoke"

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

func (p *passportClient) Authorize(callback string, scope string) (authUrl string, err error) {
	// 拼接URL PathOAuth
	path, err := url.JoinPath(p.host, PathOAuth)
	if err != nil {
		return "", err
	}
	authUrl = fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&scope=%s",
		path, p.appId, url.QueryEscape(callback), scope)
	return
}

func (p *passportClient) AccessToken(generator AccessTokenGeneratorFunc) (resp OAuthResponse, err error) {
	var config AccessTokenGenerator
	var response _oauthResponse
	var request interface{}
	generator(&config)
	path, err := url.JoinPath(ApiOAuthPrefix, string(config.Category))
	switch config.Category {
	case AuthorizeCode:
		if debugEnable {
			log.Debug("CodeAuth: %s", config.AuthorizeCode)
		}
		request = CodeRequest{
			Code: config.AuthorizeCode,
		}
	case Password:
		if debugEnable {
			log.Debug("PasswordAuth: %s", config.UserName)
		}
		request = PasswordRequest{
			UserName: config.UserName,
			Password: config.Password,
		}
	case RefreshToken:
		if debugEnable {
			log.Debug("RefreshToken: %s", config.RefreshToken)
		}
		request = RefreshRequest{
			RefreshToken: config.RefreshToken,
		}
	}

	err = NewRpc(p.host).Auth(p.appId, p.secret).BuildReq(func(rest *utils.RestClient) (int, error) {
		code, e := rest.Path(path).ResponseJson(&response).Post(request)
		if response.Code == CodeNotAllow {
			code = CodeNotAllow
		}
		return code, e
	}).Exec()

	if err != nil {
		log.Debug(err)
		return
	}

	if response.Code == CodeSuccess {
		resp = response.OAuthResponse
		return
	} else {
		err = errors.New(strconv.Itoa(response.Code) + ": " + response.Message)
		return
	}
}

func (p *passportClient) CancelAuthorize(refreshToken string) (err error) {
	// 删除缓存中的refreshToken
	var request RevokeRequest
	var response Response
	request.RefreshToken = refreshToken
	err = NewRpc(p.host).Auth(p.appId, p.secret).BuildReq(func(rest *utils.RestClient) (int, error) {
		code, e := rest.Path(PathRevoke).ResponseJson(&response).Post(request)
		if response.Code == CodeNotAllow {
			code = CodeNotAllow
		}
		return code, e
	}).Exec()
	return
}
