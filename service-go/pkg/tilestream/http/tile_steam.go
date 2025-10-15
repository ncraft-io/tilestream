package http

import (
	"compress/gzip"
	"context"

	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	ts "github.com/ncraft-io/tilestream/go/pkg/tilestream"
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
	"gopkg.in/errgo.v2/fmt/errors"
)

func init() {
	createReader := func(options core.Options) tilestream.TileReader {
		return New(options)
	}
	_ = tilestream.RegisterReader("http", createReader)

	createTileStream := func(options core.Options) tilestream.TileStream {
		return New(options)
	}
	_ = tilestream.RegisterTileStream("http", createTileStream)
}

type TileStream struct {
	Config *Config
}

func New(options core.Options) *TileStream {
	ts := &TileStream{
		Config: &Config{},
	}
	if err := options.To(ts.Config); err != nil {
		logs.Warnw("failed to set the options to TileStream.Http's config", "options", options, "err", err)
	}
	return ts
}

func (s *TileStream) getTileUrl(x, y, level int32) (string, error) {
	if s != nil && s.Config != nil && len(s.Config.TileUrl) > 0 {
		url := s.Config.TileUrl
		if strings.Contains(url, "{x}") {
			url = strings.Replace(url, "{x}", strconv.Itoa(int(x)), 1)
			if strings.Contains(url, "{y}") {
				url = strings.Replace(url, "{y}", strconv.Itoa(int(y)), 1)
				if strings.Contains(url, "{level}") {
					url = strings.Replace(url, "{level}", strconv.Itoa(int(level)), 1)
					return url, nil
				}
			}
		}
		return "", errors.Newf("invalid tile url %s", s.Config.TileUrl)
	}
	return "", errors.New("tile url config has not found")
}

func (s *TileStream) Tile(ctx context.Context, x, y, level int32) ([]byte, core.Options, error) {
	url, err := s.getTileUrl(x, y, level)
	if err != nil {
		return nil, nil, err
	}

	r, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode >= 400 {
		return nil, nil, errors.New("http return error " + r.Status)
	}

	var reader io.ReadCloser
	switch r.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(r.Body)
		defer func() {
			_ = reader.Close()
		}()
		if err != nil {
			return nil, nil, errors.New("failed to read the gzip content")
		}
	default:
		reader = r.Body
	}

	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, nil, err
	}

	typ := r.Header.Get("Content-Type")
	options := make(core.Options)
	options.SetValue("Content-Type", typ)

	return content, options, err
}

func (s *TileStream) Info(ctx context.Context) (*ts.TileInfo, error) {
	return nil, nil
}

func (s *TileStream) StartWriting(ctx context.Context) error {
	return errors.New("not implemented")
}

func (s *TileStream) StopWriting(ctx context.Context) error {
	return errors.New("not implemented")
}

func (s *TileStream) WriteTile(ctx context.Context, x, y, z int32, tile []byte) error {
	return errors.New("not implemented")
}

func (s *TileStream) WriteInfo(ctx context.Context, info *ts.TileInfo) error {
	return errors.New("not implemented")
}
