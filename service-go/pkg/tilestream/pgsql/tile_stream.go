package pgsql

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream/providers/postgis"
	ts "github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
)

type TileStream struct {
	Config *postgis.Config

	GeometryField string
	Fields        string
	Filter        string
	Table         string
}

func init() {
	//createReader := func(options core.Options) ts.TileReader {
	//	p, err := New(options)
	//	if err != nil {
	//		logs.Errorf("failed to create reader", "error", err)
	//		panic(err)
	//	}
	//
	//	return p
	//}
	//_ = ts.RegisterReader("postgis", createReader)
	//
	//createTileStream := func(options core.Options) ts.TileStream {
	//	p, err := New(options)
	//	if err != nil {
	//		logs.Errorf("failed to create tilestream", "error", err)
	//		panic(err)
	//	}
	//	return p
	//}
	//_ = ts.RegisterTileStream("postgis", createTileStream)
}

func New(options core.Options) (*TileStream, error) {
	cfg, err := postgis.NewConfig(options)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid options for postgis provider. err: %s", err.Error()))
	}
	if cfg.Provider == nil {
		return nil, errors.New("invalid options for postgis provider: should has provider")
	}

	if len(cfg.Provider.Uri) == 0 {
		cfg.Provider.Uri = ts.Conf.DefaultDbUri
		if len(cfg.Provider.Uri) == 0 {
			return nil, errors.New("invalid options for postgis provider: should has provider's sql uri")
		}
	}

	stream := &TileStream{
		Config:        cfg,
		GeometryField: cfg.Provider.GeometryField,
		Table:         cfg.Provider.Sql.Table,
		Filter:        cfg.Provider.Sql.Filter,
	}
	if len(stream.Table) == 0 {
		stream.Table = cfg.Name
	}

	fields := []string{", id as feat_id"}
	for _, field := range cfg.Provider.Sql.Fields {
		if field.Name == cfg.Provider.GeometryField || field.Name == cfg.Provider.IdField {
			continue
		}
		if len(field.Alias) > 0 {
			fields = append(fields, fmt.Sprintf("%s as %s", field.Name, field.Alias))
		} else {
			fields = append(fields, field.Name)
		}
	}
	stream.Fields = strings.Join(fields, ",")
	if len(stream.Filter) > 0 {
		stream.Filter = " and ( " + stream.Filter + " )"
	}

	//tx := ts.DB.Select("Select postgis_version()").Scan()
	return stream, nil
}

type Mvt struct {
	Tile []byte `json:"tile"`
}

func (s *TileStream) Tile(ctx context.Context, x, y, level int32) ([]byte, core.Options, error) {
	sql := fmt.Sprintf(`WITH mvtgeom AS (
		SELECT ST_AsMVTGeom(%s, ST_TRANSFORM(ST_TileEnvelope(%d, %d, %d), 4326), extent => 4096, buffer => 64) AS geom %s
		FROM %s
		WHERE %s && ST_TRANSFORM(ST_TileEnvelope(%d, %d, %d, margin => (64.0 / 4096)), 4326) %s)
	SELECT ST_AsMVT(mvtgeom.*, '%s') as tile
	FROM mvtgeom;`, s.GeometryField, level, x, y, s.Fields, s.Table, s.GeometryField, level, x, y, s.Filter, s.Config.Name)

	mvt := &Mvt{}
	tx := GetDB().Raw(sql).Scan(mvt)
	if tx.RowsAffected == 0 {
		return []byte{}, core.NewOptions(
			"Format", "mvt",
			"Content-Type", "application/vnd.mapbox-vector-tile",
		), nil
	}
	if tx.Error != nil {
		return nil, nil, core.NewNotFoundError("failed to found the tile")
	}

	buf := bytes.NewBuffer(nil)
	writer := gzip.NewWriter(buf)
	if _, err := writer.Write(mvt.Tile); err != nil {
		return nil, nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	options := core.NewOptions(
		"Format", "mvt",
		"Content-Type", "application/vnd.mapbox-vector-tile",
		"Content-Encoding", "gzip",
	)
	return buf.Bytes(), options, nil
}

func (s *TileStream) Info(ctx context.Context) (*tilestream.TileInfo, error) {
	return nil, nil
}

func (s *TileStream) StartWriting(ctx context.Context) error {
	return errors.New("not implemented")
}

func (s *TileStream) StopWriting(ctx context.Context) error {
	return errors.New("not implemented")
}

func (s *TileStream) WriteTile(ctx context.Context, x, y, z int32, tile []byte, options core.Options) error {
	return errors.New("not implemented")
}

func (s *TileStream) WriteInfo(ctx context.Context, info *tilestream.TileInfo) error {
	return errors.New("not implemented")
}
