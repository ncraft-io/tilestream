package tilestream

import (
	"context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
)

type TileInfo = tilestream.TileInfo

type TileReader interface {
	// Tile when initialization is incomplete, this will fail always.
	//
	// obtains tile and the options:
	//
	// error is set when the tile does not exist or when retrieval failed.
	// If the tile does not exist and that's OK, the error message should
	// explicitly read 'Tile does not exist' in order to be handled correctly
	// by copy.
	// otherwise, tile is a buffer containing the compressed image data
	Tile(ctx context.Context, x, y, level int32) ([]byte, core.Options, error)

	// Info when initialization is incomplete, this will fail always.
	//
	// obtains tile and calls callback:
	//
	// err is set when information retrieval failed.
	// otherwise, data is a hash containing all the information.
	Info(ctx context.Context) (*TileInfo, error)
}
