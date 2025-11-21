package kv

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
	"io/ioutil"
	"reflect"

	jsoniter "github.com/json-iterator/go"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/kvstore"
	_ "github.com/ncraft-io/ncraft/go/pkg/ncraft/kvstore/badgerdb"
	_ "github.com/ncraft-io/ncraft/go/pkg/ncraft/kvstore/bbolt"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	ts "github.com/ncraft-io/tilestream/go/pkg/tilestream"
)

const KvInfoKey = "info"

func init() {
	createReader := func(config core.Options) tilestream.TileReader {
		return New(config)
	}
	createWriter := func(config core.Options) tilestream.TileWriter {
		return New(config)
	}
	_ = tilestream.RegisterReader("kv", createReader)
	_ = tilestream.RegisterWriter("kv", createWriter)
}

type Kv struct {
	KvType       string
	CompressType string // zip, gzip,
	Store        kvstore.KvStore
}

func New(config core.Options) *Kv {
	kv := ""
	compressType := ""
	if v, ok := config["kvType"]; ok {
		if kv, ok = v.(string); !ok {
			logs.Error("the kvType field type is not string: ", reflect.TypeOf(v).String())
			return nil
		}
	} else {
		logs.Error("has no kvType field in options")
		return nil
	}

	if v, ok := config["compressType"]; ok {
		if compressType, ok = v.(string); !ok {
			logs.Warn("the compress field type is not string: ", reflect.TypeOf(v).String())
		}
	}

	store := &Kv{}
	store.KvType = kv
	store.CompressType = compressType
	store.Store = kvstore.NewStore(kv, config)
	return store
}

func getToken(x, y, level int32) string {
	return ts.NewTileId(x, y, level).QuadKey()
}

func (k *Kv) Tile(ctx context.Context, x, y, level int32) ([]byte, core.Options, error) {
	key := getToken(x, y, level)

	tile, err := k.Store.Get(ctx, []byte(key))
	if err != nil {
		return nil, nil, err
	}

	if k.CompressType == "zip" || k.CompressType == "gzip" {
		body, err := gzip.NewReader(bytes.NewReader(tile))
		if err != nil {
			fmt.Println("unzip is failed, err:", err)
			return nil, nil, err
		}
		defer body.Close()
		data, err := ioutil.ReadAll(body)
		return data, nil, err
	} else {
		return tile, nil, nil
	}
}

func (k *Kv) Info(ctx context.Context) (*ts.TileInfo, error) {
	info, err := k.Store.Get(ctx, []byte(KvInfoKey))
	if err != nil {
		return nil, err
	}

	tileInfo := &ts.TileInfo{}
	_ = jsoniter.ConfigFastest.Unmarshal(info, tileInfo)

	return tileInfo, nil
}

func (k *Kv) StartWriting(ctx context.Context) error {
	return nil
}

func (k *Kv) StopWriting(ctx context.Context) error {
	k.Store.Close()
	return nil
}

func (k *Kv) WriteTile(ctx context.Context, x, y, level int32, tile []byte) error {
	token := getToken(x, y, level)
	return k.Store.Put(ctx, []byte(token), tile)
}

func (k *Kv) WriteInfo(ctx context.Context, info *ts.TileInfo) error {
	bytes, err := jsoniter.ConfigFastest.Marshal(info)
	if err != nil {
		return err
	}
	return k.Store.Put(ctx, []byte(KvInfoKey), bytes)
}
