package scatter

import (
	"context"
	"errors"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/kvstore"
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream"

	"github.com/gogo/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	ts "github.com/ncraft-io/tilestream/go/pkg/tilestream"
)

const ScatterInfoKey = "info"

func init() {
	createReader := func(config core.Options) tilestream.TileReader {
		return NewScatter(config)
	}

	createWriter := func(config core.Options) tilestream.TileWriter {
		return NewScatter(config)
	}

	tilestream.RegisterReader("scatter", createReader)
	tilestream.RegisterWriter("scatter", createWriter)
}

// Scatter ...
// 根据基础的Tile，生成高层聚合显示的散点图
type Scatter struct {
	Config *Config
	Store  kvstore.KvStore
}

func NewScatter(config core.Options) *Scatter {
	scatter := &Scatter{}
	scatter.Config = NewConfig(config)

	scatter.Store = kvstore.NewStore(scatter.Config.KvType, config)
	return scatter
}

func (s *Scatter) Tile(ctx context.Context, x, y, level int32) ([]byte, core.Options, error) {
	key := ts.NewTileId(x, y, level).QuadKey()
	tile, err := s.Store.Get(ctx, []byte(key))
	if err != nil {
		return nil, nil, err
	}

	pt := ts.PointTile{}
	err = proto.Unmarshal(tile, &pt)
	if err != nil {
		return nil, nil, err
	}
	if len(pt.Xs) == 0 {
		return nil, nil, errors.New("empty point tile")
	}
	pt.Id = key
	bytes, _ := pt.ToMVT(s.Config.Name, false)
	if len(bytes) == 0 {
		return nil, nil, errors.New("empty vector tile")
	}

	return bytes, nil, nil
}

func (s *Scatter) Info(ctx context.Context) (*ts.TileInfo, error) {
	return nil, nil
}

func (s *Scatter) StartWriting(ctx context.Context) error {
	return nil
}

func (s *Scatter) StopWriting(ctx context.Context) error {
	s.Store.Close()
	return nil
}

func (s *Scatter) WriteTile(ctx context.Context, x, y, level int32, tile []byte) error {
	token := getToken(x, y, level)
	return s.Store.Put(ctx, []byte(token), tile)
}

func (s *Scatter) WriteInfo(ctx context.Context, info *tilestream.TileInfo) error {
	bytes, err := jsoniter.ConfigFastest.Marshal(info)
	if err != nil {
		return err
	}
	return s.Store.Put(ctx, []byte(ScatterInfoKey), bytes)
}

func getToken(x, y, level int32) string {
	return ts.NewTileId(x, y, level).QuadKey()
}
