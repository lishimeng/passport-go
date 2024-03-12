package signup

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/go-sdk/wechat"
	"github.com/lishimeng/passport-go/cmd/passport/models"
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/passport-go/internal/store"
	"github.com/lishimeng/x/container"
	"github.com/lishimeng/x/util"
)

// 微信注册 必须有session_key的验证

// union_id, 手机号

type WxMiniReq struct {
	UnionId    string `json:"unionId"`              // 微信ID
	SessionKey string `json:"sessionKey,omitempty"` // 微信登录session
	Code       string `json:"code"`                 // 获取手机号的code
}

type WxMiniResp struct {
	app.Response
	models.BasicSession
}

func wechatRegister(ctx server.Context) {
	var err error
	var req WxMiniReq
	var resp WxMiniResp

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	log.Info("unionId: %s, code:%s", req.UnionId, req.Code)
	if len(req.UnionId) == 0 || len(req.Code) == 0 {
		log.Info("param empty")
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}

	// 获取session_key
	sessionKey, ok := store.GetManager().GetDefaultStore().Load(req.SessionKey)
	if !ok { // 没有登录微信
		log.Info("session_key not found")
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	if sessionKey != req.UnionId { // session与微信ID不匹配
		log.Info("session_key not match")
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}

	// code重复(code放cache, 防止重复请求)

	// 检查重复
	// union id重复
	n, err := app.GetOrm().Context.QueryTable(new(model.MfaItem)).
		Filter("Sn", req.UnionId).Filter("Category", model.MfaWechat).Count()
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	if n > 0 {
		log.Info("union_id exist: %s", req.UnionId)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}

	// 获取phone_number
	// code-->手机号(wx)
	var client = new(wechat.Client)
	err = container.Get(&client)
	if err != nil {
		log.Info("no wechat client")
		log.Info(err)
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}

	wxResp, err := client.GetPhoneNumber(req.Code)
	if err != nil {
		log.Info("get wx phone_number error")
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	if wxResp.ErrCode != 0 {
		log.Info("get wx phone_number error")
		log.Info(wxResp.ErrCode, wxResp.ErrMsg)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}

	phoneNumber := wxResp.PhoneInfo.PurePhoneNumber
	log.Info("phone_number: %s", phoneNumber)

	// code不合法 TODO cache中检查

	// phone number重复
	n, err = app.GetOrm().Context.QueryTable(new(model.MfaItem)).
		Filter("Category", model.MfaPhoneNumber).
		Filter("Sn", phoneNumber).Count()
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	if n > 0 {
		log.Info("phone_number exist: %s", phoneNumber)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}

	var userCode = util.UUIDString()

	var userTable = new(model.UserInfo)
	userTable.UserCode = userCode
	userTable.Status = app.Enable
	userTable.Phone = phoneNumber
	_, err = app.GetOrm().Context.Insert(userTable)
	if err != nil {
		log.Info("create user fail: %s", phoneNumber)
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}

	go func() {
		// 创建用户
		// 创建profile
		e := svsAddWxUser(userCode, phoneNumber, req.UnionId)
		if e != nil {
			log.Info(e)
		}
	}()

	// 创建token
	var token = util.UUIDString() + util.UUIDString()

	// 返回用户信息
	resp.Code = tool.RespCodeSuccess
	resp.Uid = userCode
	resp.Token = token
	ctx.Json(resp)
}
