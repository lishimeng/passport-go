package signin

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/db/model"
)

// 手机验证码登录

func smsSignIn(ctx server.Context) {
	var resp LoginResp
	var req CodeLoginReq
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		ctx.Json(resp)
		return
	}
	var info model.UserInfo
	cond := orm.NewCondition()
	cond = cond.And("Phone", req.UserName)
	err = app.GetOrm().Context.QueryTable(new(model.UserInfo)).SetCond(cond).One(&info)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "请先绑定邮箱/手机"
		ctx.Json(resp)
		return
	}
	//验证码匹配激活邮箱或手机
	var value string
	switch req.CodeLoginType {
	case string(model.SmsNotifyType):
		key := string(model.SmsSighIn) + req.UserName
		err = app.GetCache().Get(key, &value)
		if err != nil {
			resp.Code = tool.RespCodeError
			resp.Message = "请先获取验证码"
			ctx.Json(resp)
			return
		}
		log.Info("code:%s,%s", value, req.Code)
		if value != req.Code {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码不正确"
			ctx.Json(resp)
			return
		}
		//登录成功时清除验证码缓存
		app.GetCache().Del(key)
	case string(model.MailNotifyType):
		key := string(model.EmailSighIn) + req.UserName
		err = app.GetCache().Get(key, &value)
		if err != nil {
			resp.Code = tool.RespCodeError
			resp.Message = "请先获取验证码"
			ctx.Json(resp)
			return
		}
		log.Info("code:%s,%s", value, req.Code)
		if value != req.Code {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码不正确"
			ctx.Json(resp)
			return
		}
		//登录成功时清除验证码缓存
		app.GetCache().Del(key)
	default:
		resp.Code = tool.RespCodeError
		resp.Message = "未匹配到登录平台"
		ctx.Json(resp)
		return
	}
	tokenContent, err := getToken(info.UserCode, req.LoginType)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token获取失败"
		ctx.Json(resp)
		return
	}
	//cache token
	go func() {
		_ = saveToken(tokenContent)
	}()
	resp.Code = tool.RespCodeSuccess
	resp.Token = string(tokenContent)
	ctx.Json(resp)

}
