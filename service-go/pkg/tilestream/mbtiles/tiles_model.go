package mbtiles

import (
	"context"
	"fmt"
	"sync"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var mt map[string]*TilesModel = make(map[string]*TilesModel)
var mtLock sync.Mutex

type TilesModel struct {
	DB        *db.DB
	Config    *Config
	err       error
	writeOnce sync.Once
	writeErr  error
}

func NewTilesModel(config *Config, layer string) *TilesModel {
	file, err := modelFile(config, layer)
	if err != nil {
		return &TilesModel{Config: config, err: err}
	}
	database, err := openModelDB(file)
	return &TilesModel{DB: database, Config: config, err: err}
}

func GetTilesModel(config *Config, layer string) *TilesModel {
	file, err := modelFile(config, layer)
	if err != nil {
		return &TilesModel{Config: config, err: err}
	}
	mtLock.Lock()
	defer mtLock.Unlock()
	if model, ok := mt[file]; ok {
		return model
	}
	model := NewTilesModel(config, layer)
	if model.err == nil {
		mt[file] = model
	}
	return model
}

func (t *TilesModel) prepareWrite() error {
	if t == nil {
		return fmt.Errorf("invalid MBTiles model")
	}
	if t.err != nil {
		return t.err
	}
	t.writeOnce.Do(func() { t.writeErr = initializeModelSchema(t.DB.DB) })
	return t.writeErr
}

func (t *TilesModel) CreateTile(ctx context.Context, tiles ...*Tiles) error {
	if err := t.prepareWrite(); err != nil {
		return err
	}
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
	if t == nil {
		return nil, fmt.Errorf("invalid MBTiles model")
	}
	if t.err != nil {
		return nil, t.err
	}
	tiles := &Tiles{}
	tx := t.DB.DB.WithContext(ctx)
	return tiles, tx.Where("zoom_level=? and tile_column=? and tile_row=?", level, x, y).Find(tiles).Error
}
