package model

import (
	"context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/query"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
)

var layer *Layer
var layerOnce sync.Once

type Layer struct {
	DB *db.DB
}

func GetLayer() *Layer {
	layerOnce.Do(func() {
		layer = NewLayer()
	})

	return layer
}

func NewLayer() *Layer {
	t := &Layer{DB: GetDB()}
	if !t.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&tilestream.Layer{}) {
		if err := d.AutoMigrate(&tilestream.Layer{}); err != nil {
			logs.ErrLog("Init Layer model err: ", err)
			panic(err)
		}
	}
	return t
}

func (a *Layer) Create(ctx context.Context, tables ...*tilestream.Layer) (int64, error) {
	tablesLen := len(tables)
	var executionResult *gorm.DB

	if tablesLen == 0 {
		return 0, nil
	} else if tablesLen == 1 {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(tables[0])
	} else {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(tables, len(tables))
	}

	return executionResult.RowsAffected, executionResult.Error
}

func (a *Layer) Update(ctx context.Context, table *tilestream.Layer) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Updates(table)
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Layer) Get(ctx context.Context, uid string) (*tilestream.Layer, error) {
	table := &tilestream.Layer{}
	return table, a.DB.WithContext(ctx).First(table, "id = ? or (name = ? and original_id = '')", uid, uid).Error
}

func (a *Layer) BatchGet(ctx context.Context, ids []string) ([]*tilestream.Layer, error) {
	var tables []*tilestream.Layer
	return tables, a.DB.WithContext(ctx).Find(&tables, ids).Error
}

func (a *Layer) Query(ctx context.Context, query *query.Query) ([]*tilestream.Layer, error) {
	var tables []*tilestream.Layer

	tx := a.DB.DB.WithContext(ctx)
	if query != nil {
		tx = query.Apply(tx, nil)
	}

	return tables, tx.Find(&tables).Error
}

func (a *Layer) Delete(ctx context.Context, uid string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Where("id = ?", uid).Delete(&tilestream.Layer{})
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Layer) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Delete(&tilestream.Layer{}, ids)
	return executionResult.RowsAffected, executionResult.Error
}
