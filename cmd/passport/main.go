package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/cache"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/passport-go/cmd/passport/ddd"
	"github.com/lishimeng/passport-go/cmd/passport/setup"
	"github.com/lishimeng/passport-go/cmd/passport/static"
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/passport-go/internal/etc"
	"github.com/lishimeng/x/container"
	"net/http"
	"time"
)
import _ "github.com/lib/pq"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 200)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, app.WithDefaultCallback(configName))
		if err != nil {
			return err
		}
		dbConfig := persistence.PostgresConfig{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.Database,
			InitDb:    true,
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}
		if etc.Config.Token.Enable {
			issuer := etc.Config.Token.Issuer
			tokenKey := []byte(etc.Config.Token.Key)
			builder = builder.EnableTokenValidator(func(inject app.TokenValidatorInjectFunc) {
				provider := token.NewJwtProvider(token.WithIssuer(issuer),
					token.WithKey(tokenKey, tokenKey), // hs256的秘钥必须是[]byte
					token.WithAlg("HS256"),
					token.WithDefaultTTL(etc.TokenTTL),
				)
				var storage token.Storage
				if etc.Config.Token.EnableJwtTokenCache {
					// 使用 Redis 缓存验证 jwt token
					storage = token.NewRedisStorage(app.GetCache())
				} else {
					storage = token.NewLocalStorage(provider)
				}
				container.Add(provider)
				inject(storage)
			})
		}
		redisOpts := cache.RedisOptions{
			Addr:     etc.Config.Redis.Addr,
			Password: etc.Config.Redis.Password,
		}
		cacheOpts := cache.Options{
			MaxSize: 10000,
			Ttl:     etc.TokenTTL,
		}
		builder.
			EnableDatabase(dbConfig.Build(), model.Tables()...).
			SetWebLogLevel("debug").
			EnableStaticWeb(func() http.FileSystem {
				return http.FS(static.Static)
			}).
			EnableOrmLog().
			EnableWeb(etc.Config.Web.Listen, ddd.Route).
			//EnableAmqp(amqp.Connector{Conn: etc.Config.Mq.Conn}).
			EnableCache(redisOpts, cacheOpts).
			ComponentAfter(setup.Setup).
			PrintVersion()
		return err
	})
	return
}
