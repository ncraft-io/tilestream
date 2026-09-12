package mbtiles

import (
	"context"
	"fmt"
	"sync"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var mm map[string]*MetadataModel = make(map[string]*MetadataModel)
var mmLock sync.Mutex

type MetadataModel struct {
	DB        *db.DB
	Config    *Config
	err       error
	writeOnce sync.Once
	writeErr  error
}

func NewMetadataModel(config *Config, layer string) *MetadataModel {
	file, err := modelFile(config, layer)
	if err != nil {
		return &MetadataModel{Config: config, err: err}
	}
	database, err := openModelDB(file)
	return &MetadataModel{DB: database, Config: config, err: err}
}

func GetMetadataModel(config *Config, layer string) *MetadataModel {
	file, err := modelFile(config, layer)
	if err != nil {
		return &MetadataModel{Config: config, err: err}
	}
	mmLock.Lock()
	defer mmLock.Unlock()
	if model, ok := mm[file]; ok {
		return model
	}
	model := NewMetadataModel(config, layer)
	if model.err == nil {
		mm[file] = model
	}
	return model
}

func (m *MetadataModel) prepareWrite() error {
	if m == nil {
		return fmt.Errorf("invalid MBTiles model")
	}
	if m.err != nil {
		return m.err
	}
	m.writeOnce.Do(func() { m.writeErr = initializeModelSchema(m.DB.DB) })
	return m.writeErr
}

func (m *MetadataModel) CreateMetadata(ctx context.Context, data ...*Metadata) error {
	if err := m.prepareWrite(); err != nil {
		return err
	}
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
	if m == nil {
		return nil, fmt.Errorf("invalid MBTiles model")
	}
	if m.err != nil {
		return nil, m.err
	}
	var data []*Metadata
	tx := m.DB.DB.WithContext(ctx)
	err := tx.Find(&data).Error
	return data, err
}
