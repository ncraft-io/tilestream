package mbtiles

import (
	"context"
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"path"
	"sync"
)

var mm map[string]*MetadataModel = make(map[string]*MetadataModel)
var mmLock sync.Mutex

type MetadataModel struct {
	DB     *db.DB
	Config *Config
}

func findMbtiles(config *Config, layer string) string {
	var file string
	for _, p := range config.Paths {
		f := path.Join(p, layer+".mbtiles")
		stat, err := os.Stat(f)
		if err != nil || os.IsNotExist(err) {
			continue
		}
		if stat.IsDir() {
			continue
		}
		file = f
		break
	}
	if len(file) == 0 {
		return path.Join(config.Paths[0], layer+".mbtiles")
	}
	return ""
}

func NewMetadataModel(config *Config, layer string) *MetadataModel {
	if config == nil || len(config.Paths) == 0 || len(layer) == 0 {
		return nil
	}

	file := findMbtiles(config, layer)
	DB, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		logs.Warnw("failed to open the mbtiles files", "file", file, "error", err)
		return nil
	}
	return &MetadataModel{
		DB: &db.DB{
			DB:     DB,
			Config: nil,
		},
		Config: config,
	}
}

func GetMetadataModel(config *Config, layer string) *MetadataModel {
	mmLock.Lock()
	defer mmLock.Unlock()

	if m, ok := mm[layer]; ok {
		return m
	}

	m := NewMetadataModel(config, layer)
	mm[layer] = m
	return m
}

func (m *MetadataModel) CreateMetadata(ctx context.Context, data ...*Metadata) error {
	length := len(data)
	var executionResult *gorm.DB

	if length == 0 {
		return nil
	} else if length == 1 {
		executionResult = m.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(data[0])
	} else {
		executionResult = m.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(data, len(data))
	}

	return executionResult.Error
}

func (m *MetadataModel) ListMetadata(ctx context.Context) ([]*Metadata, error) {
	var data []*Metadata
	tx := m.DB.DB.WithContext(ctx)
	return data, tx.Find(&data).Error
}
