package userInfo

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport-go/internal/db/model"
)

func GetUserInfoByUserName(userName string) (info model.UserInfo, err error) {
	cond := orm.NewCondition()
	cond = cond.Or("UserName", userName)
	cond = cond.Or("Phone", userName)
	cond = cond.Or("Email", userName)
	err = app.GetOrm().Context.QueryTable(new(model.UserInfo)).SetCond(cond).One(&info)
	return
}
