package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/x/util"
	"strings"
	"time"
)

func genAppId(name string) (code string) {
	now := time.Now().Format(time.RFC3339Nano)
	var tmp = fmt.Sprintf("AppId_%s_%s", now, name)
	sh := sha256.New()
	sh.Write([]byte(tmp))
	bs := sh.Sum(nil)
	code = strings.ToLower(util.BytesToHex(bs))
	return
}

func genSecret(appId string) (code string) {
	now := time.Now().Format(time.RFC3339Nano)
	var tmp = fmt.Sprintf("Secret_%s_%s", now, appId)
	sh := sha256.New()
	sh.Write([]byte(tmp))
	bs := sh.Sum(nil)
	code = strings.ToLower(util.BytesToHex(bs))
	return
}

func CreateApplication(name string, description string) (appInfo model.Application, err error) {
	appId := "app_" + genAppId(name)
	secret := genSecret(appId)
	appInfo = model.Application{
		Code:        appId,
		Name:        name,
		Description: description,
		Secret:      secret,
	}
	_, err = app.GetOrm().Context.Insert(&appInfo)
	return
}

func CreateTenant(name string) (tenant model.Tenant, err error) {
	tenantId := "org_" + genAppId(name)
	tenant = model.Tenant{
		Code: tenantId,
		Name: name,
	}
	_, err = app.GetOrm().Context.Insert(&tenant)
	return
}

func AddAuth(userCode string, appId string, tenantId string) error {
	_, err := app.GetOrm().Context.Insert(&model.UserAppAuth{
		UserCode:   userCode,
		AppCode:    appId,
		TenantCode: tenantId,
	})
	return err
}

func QueryTenants(userCode string, appId string) (tenants []string, err error) {
	var candidates []model.UserAppAuth
	_, err = app.GetOrm().Context.QueryTable(new(model.UserAppAuth)).
		Filter("UserCode", userCode).
		Filter("AppCode", appId).
		//OrderBy(...).
		All(&candidates)
	if err != nil || len(candidates) == 0 {
		return
	}
	tenants = make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		tenants = append(tenants, candidate.TenantCode)
	}
	return tenants, err
}
