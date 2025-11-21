package postgis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-spatial/geom/slippy"
	"github.com/go-spatial/tegola/atlas"
	"github.com/go-spatial/tegola/dict"
	_ "github.com/go-spatial/tegola/provider/postgis"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	postgis2 "github.com/ncraft-io/tilestream/go/pkg/tilestream/providers/postgis"
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
)

type Postgis struct {
	Config *AtlasConfig
	Atlas  *atlas.Atlas
}

func init() {
	createReader := func(options core.Options) tilestream.TileReader {
		p, err := New(options)
		if err != nil {
			logs.Errorf("failed to create reader", "error", err)
			panic(err)
		}

		return p
	}
	_ = tilestream.RegisterReader("postgis", createReader)

	createTileStream := func(options core.Options) tilestream.TileStream {
		p, err := New(options)
		if err != nil {
			logs.Errorf("failed to create tilestream", "error", err)
			panic(err)
		}
		return p
	}
	_ = tilestream.RegisterTileStream("postgis", createTileStream)
}

func New(options core.Options) (*Postgis, error) {
	cfg, err := postgis2.NewConfig(options)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid options for postgis provider. err: %s", err.Error()))
	}
	if cfg.Provider == nil {
		return nil, errors.New("invalid options for postgis provider: should has provider")
	}

	if len(cfg.Provider.Uri) == 0 {
		cfg.Provider.Uri = tilestream.Conf.DefaultDbUri
		if len(cfg.Provider.Uri) == 0 {
			return nil, errors.New("invalid options for postgis provider: should has provider's sql uri")
		}
	}

	if cfg.Bounds == nil {
		cfg.Bounds = &geom.BoundingBox{}
	}
	if cfg.Bounds.LeftBottom == nil {
		cfg.Bounds.LeftBottom = &geom.LngLat{Longitude: 120, Latitude: 30}
	}
	if cfg.Bounds.RightTop == nil {
		cfg.Bounds.RightTop = &geom.LngLat{Longitude: 123, Latitude: 32}
	}

	postgis := &Postgis{}
	postgis.Config, err = NewAtlasConfig(cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid altas config for postgis provider. err: %s", err.Error()))
	}

	atl := postgis.Config.Atlas
	if err = atl.Validate(); err != nil {
		return nil, err
	}

	// init our providers
	// but first convert []env.Map -> []dict.Dicter
	provArr := make([]dict.Dicter, len(atl.Providers))
	for i := range provArr {
		provArr[i] = atl.Providers[i]
	}

	providers, err := Providers(provArr)
	if err != nil {
		return nil, fmt.Errorf("could not register providers: %v", err)
	}

	postgis.Atlas = &atlas.Atlas{}

	// init our maps
	if err = Maps(postgis.Atlas, atl.Maps, providers); err != nil {
		return nil, fmt.Errorf("could not register maps: %v", err)
	}

	return postgis, nil
}

func (p *Postgis) Tile(ctx context.Context, x, y, level int32) ([]byte, core.Options, error) {
	// lookup our Map
	m, err := p.Atlas.Map(p.Config.MapName)
	if err != nil {
		logs.Errorf("map (%v) not configured. check your config file", p.Config.MapName)
		return nil, nil, err
	}

	// filter down the layers we need for this zoom
	m = m.FilterLayersByZoom(slippy.Zoom(level))
	if len(m.Layers) == 0 {
		logs.Errorf("map (%v) has no layers, at zoom %v", p.Config.MapName, level)
		return nil, nil, err
	}

	if p.Config.LayerName != "" {
		m = m.FilterLayersByName(p.Config.LayerName)
		if len(m.Layers) == 0 {
			//logAndError(w, http.StatusNotFound, "map (%v) has no layers, for LayerName %v at zoom %v", req.mapName, req.layerName, req.z)
			return nil, nil, err
		}
	}

	//atlas.SeedMapTile(ctx, m, uint(level), uint(x), uint(y))
	tile := slippy.Tile{
		Z: slippy.Zoom(level),
		X: uint(x),
		Y: uint(y),
	}

	// encode the tile
	b, err := m.Encode(ctx, tile, nil)
	if err != nil {
		return nil, nil, err
	}

	options := core.NewOptions(
		"Format", "mvt",
		"Content-Type", "application/vnd.mapbox-vector-tile",
		"Content-Encoding", "gzip",
	)

	return b, options, nil
}

func (p *Postgis) Info(ctx context.Context) (*tilestream.TileInfo, error) {
	return nil, nil
}

func (p *Postgis) StartWriting(ctx context.Context) error {
	return errors.New("not implemented")
}

func (p *Postgis) StopWriting(ctx context.Context) error {
	return errors.New("not implemented")
}

func (p *Postgis) WriteTile(ctx context.Context, x, y, z int32, tile []byte) error {
	return errors.New("not implemented")
}

func (p *Postgis) WriteInfo(ctx context.Context, info *tilestream.TileInfo) error {
	return errors.New("not implemented")
}
