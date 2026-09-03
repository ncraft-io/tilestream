package tilestream

import (
	"context"
	"errors"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"net/http"
)

func NewTile(x, y, level int32, tile []byte, options core.Options) *Tile {
	t := &Tile{
		X:       x,
		Y:       y,
		Level:   level,
		Content: tile,
	}
	if format := options.GetString("Format"); len(format) > 0 {
		t.Format = format
	}
	if encoding := options.GetString("Content-Encoding"); len(encoding) > 0 {
		t.Encoding = encoding
	}
	return t
}

func (x *Tile) GetOptions() core.Options {
	if x != nil {
		options := make(core.Options)
		if len(x.Format) > 0 {
			options.SetValue("Format", x.Format)
		}
		if len(x.Encoding) > 0 {
			options.SetValue("Content-Encoding", x.Encoding)
		}
		return options
	}
	return nil
}

func (x *Tile) WriteHttpResponse(ctx context.Context, writer http.ResponseWriter) error {
	if x != nil {
		switch x.Format {
		case "mvt":
			writer.Header().Set("Content-Type", "application/vnd.mapbox-vector-tile")
		case "png":
			writer.Header().Set("Content-Type", "image/png")
		case "jpeg", "jpg":
			writer.Header().Set("Content-Type", "image/jpeg")
		}

		if len(x.Encoding) > 0 {
			writer.Header().Set("Content-Encoding", x.Encoding)
		}
		size, err := writer.Write(x.Content)
		if err != nil {
			return err
		}
		if size != len(x.Content) {
			return errors.New("")
		}
	}
	return nil
}
