package app

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/passport-go/internal/db/model"
)

func getAppInfo(appId string) (ai AppInfo, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		c, e := _getClientByAppId(appId, ctx)
		if e != nil {
			return
		}
		if e != nil {
			return
		}
		ai.AppId = c.Code
		ai.Secret = c.Secret
		return
	})
	return
}

func _getClientByAppId(appId string, ctx persistence.TxContext) (c model.Application, err error) {
	err = ctx.Context.QueryTable(new(model.Application)).Filter("Code", appId).One(&c)
	return
}

func _getTenantById(id int, ctx persistence.TxContext) (t model.Tenant, err error) {
	t.Id = id
	err = ctx.Context.Read(&t)
	return
}
