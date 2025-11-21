package cache

import (
	"context"
	"errors"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/redis"
	_ "github.com/ncraft-io/ncraft/go/pkg/ncraft/redis/goredis"
	ts "github.com/ncraft-io/tilestream/go/pkg/tilestream"
	"time"
)

type Cache struct {
	Expiration *core.Duration `json:"expiration"`
	Redis      redis.Redis    `json:"redis"`
}

func New(config core.Options) *Cache {
	if config == nil {
		return nil
	}

	c := &Cache{}
	if v, ok := config["expiration"].(*core.Duration); ok {
		c.Expiration = v
	}
	if c.Expiration == nil {
		c.Expiration = core.NewDurationFromSeconds((7 * 24 * time.Hour).Seconds())
	}

	if r, ok := config["redis"].(*redis.Config); ok {
		c.Redis = redis.New(r)
	} else {
		return nil
	}

	return c
}

func getToken(ctx context.Context, x, y, level int32) string {
	layer := ""
	if hash, ok := ctx.Value("hash_key").(string); ok && len(hash) > 0 {
		layer = hash
	} else {
		layer, ok = ctx.Value("layer").(string)
	}
	if len(layer) > 0 {
		return "tilestream-cache-" + layer + "-" + ts.NewTileId(x, y, level).QuadKey()
	}
	return ""
}

func (c *Cache) Tile(ctx context.Context, x, y, level int32) ([]byte, core.Options, error) {
	if key := getToken(ctx, x, y, level); len(key) > 0 {
		tile, err := redis.Get(c.Redis, key)
		if err != nil {
			return nil, nil, err
		}
		return []byte(tile), nil, nil
	}
	return nil, nil, errors.New("the tile key is empty")
}

func (c *Cache) Info(ctx context.Context) (*ts.TileInfo, error) {
	return nil, nil
}

func (c *Cache) StartWriting(ctx context.Context) error {
	return nil
}

func (c *Cache) StopWriting(ctx context.Context) error {
	return nil
}

func (c *Cache) WriteTile(ctx context.Context, x, y, level int32, tile []byte) error {
	if token := getToken(ctx, x, y, level); len(token) > 0 {
		_, err := redis.SetEx(c.Redis, token, string(tile), c.Expiration.ToDuration())
		return err
	}
	return nil
}

func (c *Cache) WriteInfo(ctx context.Context, info *ts.TileInfo) error {
	return nil
}
