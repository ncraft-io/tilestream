package mbtiles

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
	ts "github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
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
	if m == nil || m.Config == nil {
		return ""
	}
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
		row, err := tmsRow(x, y, level)
		if err != nil {
			return nil, nil, err
		}
		t, err := GetTilesModel(m.Config, layer).GetTile(ctx, x, row, level)
		if err != nil {
			return nil, nil, err
		}
		return t.TileData, core.NewOptions("Format", m.Config.Format), nil
	}
	return nil, nil, fmt.Errorf("missing MBTiles layer")
}

func (m *Mbtiles) Info(ctx context.Context) (*tilestream.TileInfo, error) {
	if layer := m.getLayer(ctx); len(layer) > 0 {
		data, err := GetMetadataModel(m.Config, layer).ListMetadata(ctx)
		if err != nil {
			return nil, err
		}
		return ToTileInfo(data), nil

	}
	return nil, fmt.Errorf("missing MBTiles layer")
}

func (m *Mbtiles) StartWriting(ctx context.Context) error {
	layer := m.getLayer(ctx)
	if layer == "" {
		return fmt.Errorf("missing MBTiles layer")
	}
	return GetTilesModel(m.Config, layer).prepareWrite()
}

func (m *Mbtiles) StopWriting(ctx context.Context) error {
	return nil
}

func (m *Mbtiles) WriteTile(ctx context.Context, x, y, z int32, tile []byte, options core.Options) error {
	if layer := m.getLayer(ctx); len(layer) > 0 {
		row, err := tmsRow(x, y, z)
		if err != nil {
			return err
		}
		return GetTilesModel(m.Config, layer).CreateTile(ctx, &Tiles{
			ZoomLevel:  z,
			TileColumn: x,
			TileRow:    row,
			TileData:   tile,
		})

	}
	return fmt.Errorf("missing MBTiles layer")
}

func (m *Mbtiles) WriteInfo(ctx context.Context, info *tilestream.TileInfo) error {
	if info == nil {
		return fmt.Errorf("missing MBTiles metadata")
	}
	if layer := m.getLayer(ctx); len(layer) > 0 {
		data := FromTileInfo(info)
		return GetMetadataModel(m.Config, layer).CreateMetadata(ctx, data...)
	}
	return fmt.Errorf("missing MBTiles layer")
}

// Tile and WriteTile use XYZ coordinates; MBTiles stores TMS rows.
func tmsRow(x, y, z int32) (int32, error) {
	if z < 0 || z > 31 {
		return 0, fmt.Errorf("invalid tile zoom: %d", z)
	}
	max := (int64(1) << uint(z)) - 1
	if x < 0 || y < 0 || int64(x) > max || int64(y) > max {
		return 0, fmt.Errorf("invalid tile coordinates: %d/%d/%d", z, x, y)
	}
	return int32(max - int64(y)), nil
}
