package app

import (
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/common"
	"github.com/lishimeng/passport-go/internal/config"
	"github.com/lishimeng/passport-go/internal/gentoken"
	"github.com/lishimeng/x/container"
	"github.com/pkg/errors"
)

type AppInfo struct {
	AppId  string
	Secret string
	Org    string
}

type CredentialReq struct {
	AppId  string `json:"appId,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type CredentialResp struct {
	app.Response
	Token string `json:"token,omitempty"`
}

func genCredential(ctx server.Context) {
	var err error
	var req CredentialReq
	var resp CredentialResp

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		log.Debug(errors.Wrap(err, "读取参数错误"))
		resp.Code = tool.RespCodeError
		resp.Message = "request must be a json format"
		ctx.Json(resp)
		return
	}

	if len(req.AppId) == 0 {
		log.Debug("appId: %s, secret:%s", req.AppId, req.Secret)
		resp.Code = common.CodeAppNotFound
		resp.Message = common.MsgAppNotFound
		ctx.Json(resp)
		return
	}
	c, err := getAppInfo(req.AppId)
	if err != nil {
		log.Debug(errors.Wrap(err, fmt.Sprintf("appId not found:%s", req.AppId)))
		resp.Code = common.CodeAppNotFound
		resp.Message = common.MsgAppNotFound
		ctx.Json(resp)
		return
	}
	if c.Secret != req.Secret {
		log.Debug("appId: %s, secret not match", req.AppId)
		resp.Code = common.CodeSecretNotValid
		resp.Message = common.MsgSecretNotValid
		ctx.Json(resp)
		return
	}

	var provider token.JwtProvider
	err = container.Get(&provider)
	if err != nil {
		log.Debug(errors.Wrap(err, "no jwt provider"))
		resp.Code = tool.RespCodeError
		resp.Message = "Internal Error"
		ctx.Json(resp)
		return
	}

	p := token.JwtPayload{
		Uid:   c.AppId,
		Scope: common.Scope,
	}
	bs, err := provider.GenWithTTL(p, config.Config.TTL.AccessToken)
	if err != nil {
		log.Debug(errors.Wrap(err, "gen credential err"))
		resp.Code = tool.RespCodeError
		resp.Message = "Internal Error"
		ctx.Json(resp)
		return
	}

	gentoken.SaveToken(bs)

	resp.Token = string(bs)
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
	return
}

type TokenVerifyResp struct {
	app.Response
	token.HttpTokenResp
}

func verify(ctx server.Context) {

	var resp TokenVerifyResp
	resp.Valid = true
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
