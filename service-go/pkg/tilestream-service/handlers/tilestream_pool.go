package handlers

import (
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream/cache"
	"sync"
)

var streamOnce sync.Once
var stream *TileStream

func GetTileStream() *TileStream {
	streamOnce.Do(func() {
		stream = NewTileStream()
	})

	return stream
}

type TileStream struct {
	TileStreams sync.Map
	Cache       *cache.Cache
}

func NewTileStream() *TileStream {
	return &TileStream{
		Cache: cache.New(cache.NewConfig()),
	}
}

func (t *TileStream) Get(layer string) tilestream.TileStream {
	if st, ok := t.TileStreams.Load(layer); ok {
		return st.(tilestream.TileStream)
	} else {
		s := tilestream.NewTileStream(layer)
		if s == nil {
			return nil
		}

		t.TileStreams.Store(layer, s)
		return s
	}
}

func (t *TileStream) Delete(layer string) {
	t.TileStreams.Delete(layer)
}
