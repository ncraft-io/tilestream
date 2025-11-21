package mbtiles

import (
	"context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
	ts "github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
	"path"
	"strings"
)

type Mbtiles struct {
	Config *Config
}

func IsMbtilesFile(file string) bool {
	return strings.HasSuffix(file, ".mbtiles")
}

func init() {
	createReader := func(options core.Options) ts.TileReader {
		return New(options)
	}
	createWriter := func(options core.Options) ts.TileWriter {
		return New(options)
	}
	createTileStream := func(options core.Options) ts.TileStream {
		return New(options)
	}
	_ = ts.RegisterReader("mbtiles", createReader)
	_ = ts.RegisterWriter("mbtiles", createWriter)
	_ = ts.RegisterTileStream("mbtiles", createTileStream)
}

func New(options core.Options) *Mbtiles {
	mb := &Mbtiles{
		Config: &Config{},
	}
	_ = options.To(mb.Config)
	return mb
}

func NewFromFile(file string) *Mbtiles {
	if !IsMbtilesFile(file) {
		return nil
	}

	file = strings.TrimSuffix(file, ".mbtiles")
	options := core.Options{"layer": path.Base(file), "paths": []string{path.Dir(file)}}
	return New(options)
}

func (m *Mbtiles) getLayer(ctx context.Context) string {
	if len(m.Config.Layer) > 0 {
		return m.Config.Layer
	}
	if layer := ctx.Value("layer"); layer != nil {
		if name, ok := layer.(string); ok {
			return name
		}
	}
	return ""
}

func (m *Mbtiles) Tile(ctx context.Context, x, y, level int32) ([]byte, core.Options, error) {
	if layer := m.getLayer(ctx); len(layer) > 0 {
		t, err := GetTilesModel(m.Config, layer).GetTile(ctx, x, y, level)
		if err != nil {
			return nil, nil, nil
		}
		return t.TileData, nil, nil
	}
	return nil, nil, nil
}

func (m *Mbtiles) Info(ctx context.Context) (*tilestream.TileInfo, error) {
	if layer := m.getLayer(ctx); len(layer) > 0 {
		data, err := GetMetadataModel(m.Config, layer).ListMetadata(ctx)
		if err != nil {
			return nil, nil
		}
		return ToTileInfo(data), nil

	}
	return nil, nil
}

func (m *Mbtiles) StartWriting(ctx context.Context) error {
	return nil
}

func (m *Mbtiles) StopWriting(ctx context.Context) error {
	return nil
}

func (m *Mbtiles) WriteTile(ctx context.Context, x, y, z int32, tile []byte) error {
	if layer := m.getLayer(ctx); len(layer) > 0 {
		return GetTilesModel(m.Config, layer).CreateTile(ctx, &Tiles{
			ZoomLevel:  z,
			TileColumn: y,
			TileRow:    x,
			TileData:   tile,
		})

	}
	return nil
}

func (m *Mbtiles) WriteInfo(ctx context.Context, info *tilestream.TileInfo) error {
	if layer := m.getLayer(ctx); len(layer) > 0 {
		data := FromTileInfo(info)
		return GetMetadataModel(m.Config, layer).CreateMetadata(ctx, data...)
	}
	return nil
}
