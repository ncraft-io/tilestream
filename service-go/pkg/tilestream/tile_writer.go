package tilestream

import (
	"context"
)

type TileWriter interface {
	// StartWriting Opens the TileStore in write mode.
	//
	// err is null when write mode could be established successfully.
	//
	// This function must be reentrant: Write mode may only be ended after the
	// same number of calls to .StopWriting(). Use a counter to keep track of
	// how often write mode was started.
	StartWriting(ctx context.Context) error

	// StopWriting Ends the write mode.
	// Call the callback function when the request was successfully completed.
	// This doesn't mean that write mode has been ended (see below).
	//
	// err is null when write mode could be established successfully.
	//
	// When caching and batch-writing tiles, they must be committed to the tile
	// store when this function is called, even when write mode is not ended.
	// This is true for grids as well.
	//
	// This function must be reentrant: Write mode may only be ended after the
	// same number of calls to .StopWriting(). Use a counter to keep track of
	// how often write mode was started.
	StopWriting(ctx context.Context) error

	// WriteTile Stores a tile into the data store. Parameters are in XYZ format.
	// `tile` must be a Buffer containing the compressed image.
	//
	// err is null when the write request was received successfully.
	// This doesn't mean that the tile was already written to the data store.
	//
	// Implementations may decide to cache multiple tile requests and only
	// commit them to the data store periodically to improve performance.
	// Therefore, users MUST NOT rely on this function to persist changes.
	// If you want to make sure that all changes are persisted, call
	// .StopWriting().
	WriteTile(ctx context.Context, x, y, z int32, tile []byte) error

	// WriteInfo Stores metadata into the tile writer. Info is a key-value hash with metadata.
	// Implementations may decide to reject invalid keys.
	//
	// err is null when the metadata could be written successfully.
	WriteInfo(ctx context.Context, info *TileInfo) error
}
