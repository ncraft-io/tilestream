package tilestream

import (
	"context"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/tilestream/service-go/pkg/model"
	"sync"
)

var (
	readers        map[string]func(options core.Options) TileReader
	writers        map[string]func(options core.Options) TileWriter
	tileStreams    map[string]func(options core.Options) TileStream
	readersLock    sync.Mutex
	writersLock    sync.Mutex
	tileStreamLock sync.Mutex
)

type TileStream interface {
	TileReader
	TileWriter
}

func init() {
	readers = make(map[string]func(options core.Options) TileReader)
	writers = make(map[string]func(options core.Options) TileWriter)
	tileStreams = make(map[string]func(options core.Options) TileStream)
}

// RegisterReader register a reader function to process a command
func RegisterReader(name string, reader func(options core.Options) TileReader) error {
	readersLock.Lock()
	defer readersLock.Unlock()

	if _, ok := readers[name]; ok {
		// warning
	}

	readers[name] = reader
	return nil
}

func LoadReader(name string, options core.Options) TileReader {
	if reader, ok := readers[name]; ok {
		return reader(options)
	}
	return nil
}

func RegisterWriter(name string, writer func(options core.Options) TileWriter) error {
	writersLock.Lock()
	defer writersLock.Unlock()

	if _, ok := writers[name]; ok {
		// warning
	}

	writers[name] = writer
	return nil
}

func LoadWriter(name string, options core.Options) TileWriter {
	writersLock.Lock()
	defer writersLock.Unlock()

	if writer, ok := writers[name]; ok {
		return writer(options)
	}
	return nil
}

func RegisterTileStream(name string, tileStream func(options core.Options) TileStream) error {
	tileStreamLock.Lock()
	defer tileStreamLock.Unlock()

	if _, ok := tileStreams[name]; ok {
		// warning
	}

	tileStreams[name] = tileStream
	return nil
}

func LoadTileStream(name string, options core.Options) TileStream {
	tileStreamLock.Lock()
	defer tileStreamLock.Unlock()

	if ts, ok := tileStreams[name]; ok {
		return ts(options)
	}
	return nil
}

func NewTileStream(name string) TileStream {
	var confs []*Config
	_ = config.ScanFrom(&confs, "tilestreams")
	for _, conf := range confs {
		if conf.Name == name {
			return LoadTileStream(conf.Type, conf.Options)
		}
	}

	if layer, err := model.GetLayer().Get(context.Background(), name); err == nil {
		options := layer.Config.ToOptions()
		if options != nil {
			return LoadTileStream(layer.Type, options)
		}
	}

	return nil
}

func List() {
}

func Copy() {
}
