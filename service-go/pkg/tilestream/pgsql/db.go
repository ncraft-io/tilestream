package pgsql

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	ts "github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
	"sync"
)

var d *db.DB
var dOnce sync.Once

func GetDB() *db.DB {
	dOnce.Do(func() {
		cfg := &db.Config{
			Driver:              "postgres",
			Dsn:                 ts.Conf.DefaultDbUri,
			Debug:               true,
			MaxIdleConnections:  0,
			MaxOpenConnections:  0,
			ConnectionKeepAlive: nil,
		}
		//err := config.Get("db").Scan(cfg)
		//if err != nil {
		//	logs.Errorw("failed to get the db config", "error", err.Error())
		//	panic("failed to get the db config")
		//}

		if d = db.New(cfg); d == nil {
			panic("create the db failed")
		}
	})
	return d
}
