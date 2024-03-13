package signin

// 微信Code登录

// 如无用户, 返回微信session_key, 记录session_key

// 如有用户, 返回用户信息Token

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/go-sdk/wechat"
	"github.com/lishimeng/passport-go/internal/data"
	"github.com/lishimeng/x/container"
)

type WxJsCodeReq struct {
	Code string `json:"code,omitempty"`
}
type WxLoginResp struct {
	app.Response
	data.UserToken
	wechat.JsSession
}

func wechatSignIn(ctx server.Context) {
	// 微信js_code
	// 微信openid/unionId
	// 微信session_key

	var err error
	var req WxJsCodeReq
	var resp WxLoginResp
	err = ctx.C.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}

	log.Info("wx login code: %s", req.Code)
	if len(req.Code) == 0 {
		log.Info("微信登录code为空")
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}

	// TODO cache重复请求(code), 至少2小时

	var wxClient wechat.Client
	err = container.Get(&wxClient)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	result, err := wxClient.JsCode2Session(req.Code)
	if result.ErrCode != 0 {
		log.Info("获取wx session_key失败, code: %s, err: %s", result.ErrCode, result.ErrMsg)
		// TODO
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	resp.JsSession = result // wx session
	// 换取union_id
	var unionId = result.UnionId
	mfaItem, err := svsGetUnionId(unionId)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeSuccess
		ctx.Json(resp)
		return
	}
	if mfaItem.Status != app.Enable {
		log.Info("mfaItem.Status != app.Enable")
	}

	resp.Code = tool.RespCodeSuccess
	resp.Uid = mfaItem.MfaCode
	resp.Token = genToken(mfaItem.MfaCode)
	ctx.Json(resp)
}
