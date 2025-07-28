package sdk

import (
	"errors"
	"github.com/lishimeng/passport-go/utils"
)

type AccessGeneratorCategory string

const (
	AuthorizeCode AccessGeneratorCategory = "code"
	Password      AccessGeneratorCategory = "password"
	RefreshToken  AccessGeneratorCategory = "refreshToken"
)

type AccessTokenGenerator struct {
	Category      AccessGeneratorCategory
	AuthorizeCode string
	RefreshToken  string
	UserName      string
	Password      string
}
type AccessTokenGeneratorFunc func(config *AccessTokenGenerator)

var WithCode = func(code string) AccessTokenGeneratorFunc {
	return func(config *AccessTokenGenerator) {
		config.Category = AuthorizeCode
		config.AuthorizeCode = code
	}
}

var WithPassword = func(userName string, password string) AccessTokenGeneratorFunc {
	return func(config *AccessTokenGenerator) {
		config.Category = Password
		config.UserName = userName
		config.Password = password
	}
}

var WithRefreshToken = func(refreshToken string) AccessTokenGeneratorFunc {
	return func(config *AccessTokenGenerator) {
		config.Category = RefreshToken
		config.RefreshToken = refreshToken
	}
}

type Client interface {
	// Authorize 返回用于前端的跳转到第三方授权页面的URL
	Authorize(callback string, scope string) (url string, err error)                // 1)
	AccessToken(generator AccessTokenGeneratorFunc) (resp OAuthResponse, err error) //  2)
	CancelAuthorize(refreshToken string) (err error)
}

type Option func(*passportClient)

var (
	InnerServerErr = errors.New("500")
	NotFoundErr    = errors.New("404")
)

var (
	debugEnable = false
)

// Debug 输出sdk中的log
func Debug(enable bool) {
	debugEnable = enable
	utils.DebugEnable = enable
}

func WithHost(host string) Option {
	return func(client *passportClient) {
		client.host = host
	}
}

// WithAuth Credentials配置
func WithAuth(appKey, secret string) Option {
	return func(client *passportClient) {
		client.appId = appKey
		client.secret = secret
	}
}

// New 初始化.
//
// 内部存储了Credentials,应该确保复用,而不是每次新建
//
// Client 内置了刷新credentials功能,不需要考虑credentials的获取问题.
func New(options ...Option) (m Client) {
	c := &passportClient{}
	for _, opt := range options {
		opt(c)
	}
	m = c
	return
}
