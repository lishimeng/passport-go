package data

import (
	"github.com/lishimeng/x/util"
	"strings"
)

func CreateRandCode() string {
	code := util.UUIDString()
	code = strings.ToLower(strings.ReplaceAll(code, "-", ""))
	return code
}
