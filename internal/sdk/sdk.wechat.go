package sdk

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/go-sdk/wechat"
	"github.com/lishimeng/x/container"
)

func CreateWechatClient(config WechatConfig) {
	log.Info("init wechat client")
	var client *wechat.Client
	client = wechat.New(config.AppID, config.AppSecret)
	container.Add(client)
}
