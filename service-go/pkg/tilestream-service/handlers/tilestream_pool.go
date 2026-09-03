package handlers

import (
	"github.com/hashicorp/golang-lru/v2/expirable"
	"sync"
	"time"

	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream/cache"
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
	TileStreams *expirable.LRU[string, tilestream.TileStream]
	Cache       *cache.Cache
}

func NewTileStream() *TileStream {
	return &TileStream{
		TileStreams: expirable.NewLRU[string, tilestream.TileStream](100, nil, time.Hour*12),
		Cache:       cache.New(cache.NewConfig()),
	}
}

func (t *TileStream) Get(layer string, hashkey string) tilestream.TileStream {
	key := layer + "/" + hashkey
	if st, ok := t.TileStreams.Get(key); ok {
		return st
	} else {
		s := tilestream.NewTileStream(layer)
		if s == nil {
			return nil
		}

		t.TileStreams.Add(key, s)
		return s
	}
}
