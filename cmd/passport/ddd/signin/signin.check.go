package signin

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport-go/internal/db/model"
)

func checkToken(ctx server.Context) {
	var resp app.Response
	userCode := ctx.C.Request().Header.Get(auth.UidKey)
	//loginType := ctx.C.Request().Header.Get(auth.ClientKey)
	resp.Code = 200

	// 查询数据库是否有此 userCode
	var info model.UserInfo
	err := app.GetOrm().Context.QueryTable(new(model.UserInfo)).
		Filter("UserCode", userCode).
		One(&info)
	if err != nil {
		resp.Code = 401
		resp.Message = "Invalid User"
		ctx.Json(&resp)
		return
	}
	ctx.Json(&resp)
	return
}
