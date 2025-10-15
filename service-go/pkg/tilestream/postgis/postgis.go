package postgis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-spatial/geom/slippy"
	"github.com/go-spatial/tegola/atlas"
	"github.com/go-spatial/tegola/config"
	"github.com/go-spatial/tegola/dict"
	_ "github.com/go-spatial/tegola/provider/postgis"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/kvstore"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
	"strings"
)

type Postgis struct {
	Config *Config
	Atlas  *atlas.Atlas
}

func init() {
	createReader := func(options core.Options) tilestream.TileReader {
		p, _ := New(options)
		return p
	}
	_ = tilestream.RegisterReader("postgis", createReader)

	createTileStream := func(options core.Options) tilestream.TileStream {
		p, _ := New(options)
		return p
	}
	_ = tilestream.RegisterTileStream("postgis", createTileStream)
}

func New(options core.Options) (*Postgis, error) {
	postgis := &Postgis{}
	postgis.Config = &Config{}

	atlasConf := options["atlas"]
	options["atlas"] = nil
	_ = kvstore.ResetOptions(options, postgis.Config)
	if atlasConf == nil {
		panic("atlas config is null")
	}
	if atlasStr, e := resetAtlasOptions(atlasConf.(map[string]interface{})); e == nil {
		r := strings.NewReader(atlasStr)
		postgis.Config.Atlas, e = config.Parse(r, "")
		if e != nil {
			//logs.Error(e)
			logs.ErrLogw("", "err", e)
		}
	}

	//var e error
	//postgis.Config.Atlas, e = config.Load("config.toml")
	//if e != nil{
	//	logs.Error(e)
	//}

	aconf := postgis.Config.Atlas
	aconf.ConfigureTileBuffers()

	if err := aconf.Validate(); err != nil {
		return nil, err
	}

	// init our providers
	// but first convert []env.Map -> []dict.Dicter
	provArr := make([]dict.Dicter, len(aconf.Providers))
	for i := range provArr {
		provArr[i] = aconf.Providers[i]
	}

	providers, err := Providers(provArr)
	if err != nil {
		return nil, fmt.Errorf("could not register providers: %v", err)
	}

	// init our maps
	if err = Maps(nil, aconf.Maps, providers); err != nil {
		return nil, fmt.Errorf("could not register maps: %v", err)
	}

	return postgis, nil
}

func (p Postgis) Tile(ctx context.Context, x, y, level int32) ([]byte, core.Options, error) {
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
	)

	return b, options, nil
}

func (p Postgis) Info(ctx context.Context) (*tilestream.TileInfo, error) {
	return nil, nil
}

func (p Postgis) StartWriting(ctx context.Context) error {
	return errors.New("not implemented")
}

func (p Postgis) StopWriting(ctx context.Context) error {
	return errors.New("not implemented")
}

func (p Postgis) WriteTile(ctx context.Context, x, y, z int32, tile []byte) error {
	return errors.New("not implemented")
}

func (p Postgis) WriteInfo(ctx context.Context, info *tilestream.TileInfo) error {
	return errors.New("not implemented")
}
