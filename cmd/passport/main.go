package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/passport-go/cmd/passport/ddd"
	"github.com/lishimeng/passport-go/cmd/passport/setup"
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/passport-go/internal/etc"
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

		builder.
			EnableDatabase(dbConfig.Build(), model.Tables()...).
			SetWebLogLevel("debug").
			EnableOrmLog().
			EnableWeb(etc.Config.Web.Listen, ddd.Route).
			//EnableAmqp(amqp.Connector{Conn: etc.Config.Mq.Conn}).
			ComponentAfter(setup.Setup).
			PrintVersion()
		return err
	})
	return
}
