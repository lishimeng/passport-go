package oauth

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

type PasswordReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AuthResponse struct {
	app.Response
	Token string `json:"token"`
}

func passwordAuth(ctx server.Context) {
	var req PasswordReq
	var resp AuthResponse
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		ctx.Json(&resp)
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
	tokenContent, err := gentoken.GenToken(info.Username, "passportAuth")
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token获取失败"
		ctx.Json(resp)
		return
	}
	//cache token
	go func() {
		_ = gentoken.SaveToken(tokenContent)
	}()
	resp.Code = 200
	resp.Message = "success"
	resp.Token = string(tokenContent)
	ctx.Json(&resp)
}
