package signin

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/passport-go/internal/gentoken"
	"github.com/lishimeng/passport-go/internal/passwd"
)

type LoginReq struct {
	Password  string `json:"password,omitempty"`
	UserName  string `json:"userName,omitempty"`
	Code      string `json:"code,omitempty"`
	LoginType string `json:"loginType,omitempty"` //登录方式-pc/app/wx
}

type CodeLoginReq struct {
	UserName      string `json:"userName,omitempty"`
	Code          string `json:"code,omitempty"`
	CodeLoginType string `json:"codeLoginType,omitempty"` //登录方式-sms/mail
	LoginType     string `json:"loginType,omitempty"`     //登录方式-pc/app/wx
}

type LoginResp struct {
	app.Response
	Token string `json:"token,omitempty"`
	Uid   int    `json:"uid,omitempty"`
}

// 密码登录

func passwdSignIn(ctx server.Context) {
	var resp LoginResp
	var req LoginReq
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		ctx.Json(resp)
		return
	}
	var info model.UserInfo
	cond := orm.NewCondition()
	cond = cond.And("Username", req.UserName)
	err = app.GetOrm().Context.QueryTable(new(model.UserInfo)).SetCond(cond).One(&info)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "用户名或密码错误"
		ctx.Json(resp)
		return
	}
	p := passwd.Verify(req.Password, info)
	log.Info("password:%s,%s", info.Password, p)
	if !p {
		resp.Code = tool.RespCodeError
		resp.Message = "用户名或密码错误"
		ctx.Json(resp)
		return
	}
	tokenContent, err := gentoken.GenToken(info.UserCode, req.LoginType)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token获取失败"
		ctx.Json(resp)
		return
	}

	gentoken.SaveToken(tokenContent)

	resp.Code = tool.RespCodeSuccess
	resp.Token = string(tokenContent)
	ctx.Json(resp)
}
