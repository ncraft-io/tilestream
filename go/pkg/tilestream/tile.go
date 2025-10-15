package tilestream

import (
	"context"
	"errors"
	"net/http"
)

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
