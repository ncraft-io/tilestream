package cache

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/redis"
)

type Config struct {
	Expiration *core.Duration `json:"expiration"`
	Redis      *redis.Config  `json:"redis"`
}

func NewConfig() core.Options {
	config.Get("cache")
	var conf Config
	if err := config.ScanFrom(&conf, "cache"); err != nil {
		logs.ErrLogw("failed to load the tile cache config", "error", err)
		return nil
	}

	if conf.Redis == nil {
		conf.Redis = &redis.Config{}
		if err := config.ScanFrom(conf.Redis, "redis"); err != nil {
			logs.ErrLogw("failed to load the tile cache redis config", "error", err)
			return nil
		}
	}

	return core.Options{
		"expiration": conf.Expiration,
		"redis":      conf.Redis,
	}
}
