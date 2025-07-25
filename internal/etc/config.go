package etc

import "time"

var Config Configuration
var TokenTTL = time.Hour * 2
var CodeTTL = time.Minute * 5
var PwdRefreshTTL = time.Hour * 24 * 14
var CodeRefreshTTL = time.Hour * 24 * 30
var EnableJwtTokenCache = false

// todo：改为从数据库加载配置
