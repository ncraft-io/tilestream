package mbtiles

import (
	"sync"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
)

var d *db.DB
var dOnce sync.Once

type DBConfig struct {
	MaxIdleConn int `json:"MaxIdleConn" default:"2"` // 最大空闲链接
}

func GetDB(layer string) *db.DB {
	dOnce.Do(func() {
		cfg := &db.Config{}
		err := config.Get("db").Scan(cfg)
		if err != nil {
			logs.Errorw("failed to get the db config", "error", err.Error())
			panic("failed to get the db config")
		}

		if d = db.New(cfg); d == nil {
			panic("create the db failed")
		}
	})
	return d
}
