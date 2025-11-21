package mbtiles

import (
	"context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
)

var mt map[string]*TilesModel = make(map[string]*TilesModel)
var mtLock sync.Mutex

type TilesModel struct {
	DB     *db.DB
	Config *Config
}

func NewTilesModel(config *Config, layer string) *TilesModel {
	if config == nil || len(config.Paths) == 0 || len(layer) == 0 {
		return nil
	}

	file := findMbtiles(config, layer)
	DB, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		logs.Warnw("failed to open the mbtiles files", "file", file, "error", err)
		return nil
	}
	return &TilesModel{
		DB: &db.DB{
			DB:     DB,
			Config: nil,
		},
		Config: config,
	}
}

func GetTilesModel(config *Config, layer string) *TilesModel {
	mtLock.Lock()
	defer mtLock.Unlock()

	if m, ok := mt[layer]; ok {
		return m
	}

	m := NewTilesModel(config, layer)
	mt[layer] = m
	return m
}

func (t *TilesModel) CreateTile(ctx context.Context, tiles ...*Tiles) error {
	length := len(tiles)
	var executionResult *gorm.DB

	if length == 0 {
		return nil
	} else if length == 1 {
		executionResult = t.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(tiles[0])
	} else {
		executionResult = t.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(tiles, len(tiles))
	}

	return executionResult.Error
}

func (t *TilesModel) GetTile(ctx context.Context, x, y, level int32) (*Tiles, error) {
	tiles := &Tiles{}
	tx := t.DB.DB.WithContext(ctx)
	return tiles, tx.Find(tiles).Error
}
