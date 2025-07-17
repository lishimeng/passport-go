package signup

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/passport-go/internal/passwd"
	"github.com/lishimeng/x/util"
)

type RegisterReq struct {
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}
type phoneRegisterReq struct {
	Mobile string `json:"mobile,omitempty"`
	Code   string `json:"code,omitempty"`
}
type emailRegisterReq struct {
	Email string `json:"email,omitempty"`
	Code  string `json:"code,omitempty"`
}

func phoneRegister(ctx server.Context) {
	var resp app.Response
	var req phoneRegisterReq
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		ctx.Json(resp)
		return
	}
	// 检查重复
	// phone重复
	n, err := app.GetOrm().Context.QueryTable(new(model.UserInfo)).
		Filter("Phone", req.Mobile).Count()
	if n > 0 {
		log.Info("phone_number exist: %s", req.Mobile)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	var value string
	key := string(model.SmsSighup) + req.Mobile
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
	var userCode = util.UUIDString()
	var userTable = new(model.UserInfo)
	userTable.UserCode = userCode
	userTable.Status = app.Enable
	userTable.Phone = req.Mobile
	_, err = app.GetOrm().Context.Insert(userTable)
	if err != nil {
		log.Info("create user fail: %s", req.Mobile)
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	go func() {
		_ = app.GetCache().Del(key)
		// 创建用户
		// 创建profile
		e := svsAddPhoneUser(userCode, req.Mobile)
		if e != nil {
			log.Info(e)
		}
	}()
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)

}

func emailRegister(ctx server.Context) {
	var resp app.Response
	var req emailRegisterReq
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		ctx.Json(resp)
		return
	}
	// 检查重复
	// email 重复
	n, err := app.GetOrm().Context.QueryTable(new(model.UserInfo)).
		Filter("Email", req.Email).Count()
	if n > 0 {
		log.Info("email exist: %s", req.Email)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	var value string
	key := string(model.SmsSighup) + req.Email
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
	var userCode = util.UUIDString()
	var userTable = new(model.UserInfo)
	userTable.UserCode = userCode
	userTable.Status = app.Enable
	userTable.Email = req.Email
	_, err = app.GetOrm().Context.Insert(userTable)
	if err != nil {
		log.Info("create user fail: %s", req.Email)
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	go func() {
		_ = app.GetCache().Del(key)
		// 创建用户
		// 创建profile
		e := svsAddEmailUser(userCode, req.Email)
		if e != nil {
			log.Info(e)
		}
	}()
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
func register(ctx server.Context) {
	var resp app.Response
	var req RegisterReq
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		ctx.Json(resp)
		return
	}
	// 检查重复
	// email 重复
	cond := orm.NewCondition().Or("Email", req.Email).Or("Phone", req.Mobile)
	n, err := app.GetOrm().Context.QueryTable(new(model.UserInfo)).
		SetCond(cond).Count()
	if n > 0 {
		log.Info("userInfo exist")
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	var userCode = util.UUIDString()
	var userTable = model.UserInfo{}
	userTable.UserCode = userCode
	userTable.Status = app.Enable
	userTable.Phone = req.Mobile
	userTable.Email = req.Email
	userTable.Username = req.Name
	_, err = app.GetOrm().Context.Insert(&userTable)
	if err != nil {
		log.Info("create user fail: %s", req.Mobile)
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	err = app.GetOrm().Context.QueryTable(new(model.UserInfo)).
		SetCond(cond).
		One(&userTable)
	if err != nil {
		log.Info("create user fail: %s", req.Mobile)
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	p := passwd.Generate(req.Password, userTable)
	userTable.Password = p
	_, err = app.GetOrm().Context.Update(&userTable)
	if err != nil {
		log.Info("create user fail: %s", req.Mobile)
		log.Info(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	go func() {
		// 创建用户
		// 创建profile
		e := svsAddUser(userCode, req.Email, req.Mobile)
		if e != nil {
			log.Info(e)
		}
	}()
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)

}
