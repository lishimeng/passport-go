package signin

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/x/util"
)

func svsGetUnionId(unionId string) (item model.MfaItem, err error) {

	err = app.GetOrm().Context.
		QueryTable(new(model.MfaItem)).
		Filter("Status", app.Enable).
		Filter("Category", model.MfaWechat).
		Filter("Sn", unionId).
		One(&item)

	return
}

func genToken(uid string) (tk string) {
	tk = util.RandStr(32)
	return
}
