package sdk

import (
	"errors"
	"github.com/lishimeng/passport-go/utils"
)

type Client interface {
	PasswordAuth(req PasswordRequest) (resp AuthResponse, err error)
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
