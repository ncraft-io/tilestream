### TileReader Interface

```go
type TileReader interface {
	// when initialization is incomplete, this will fail always.
	//
	// obtains tile and the options:
	//
	// error is set when the tile does not exist or when retrieval failed.
	// If the tile does not exist and that's OK, the error message should
	// explicitly read 'Tile does not exist' in order to be handled correctly
	// by copy.
	// otherwise, tile is a buffer containing the compressed image data
	Tile(x, y, z int32) ([]byte, *ReaderOptions, error)

	// when initialization is incomplete, this will fail always.
	//
	// obtains tile and calls callback:
	//
	// err is set when information retrieval failed.
	// otherwise, data is a hash containing all the information.
	Info() (*TileInfo, error)
}

// Options
type ReaderOptions map[string]interface{}
```



### TileWriter Interface

```go
// TileWriter
type TileWriter interface {
	// Opens the TileStore in write mode.
	//
	// err is null when write mode could be established successfully.
	//
	// This function must be reentrant: Write mode may only be ended after the
	// same number of calls to .StopWriting(). Use a counter to keep track of
	// how often write mode was started.
	StartWriting() error

	// Ends the write mode.
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
	StopWriting() error

	// Stores a tile into the data store. Parameters are in XYZ format.
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
	WriteTile(x, y, z int32, tile []byte) error

	// Stores metadata into the tile writer. Info is a key-value hash with metadata.
	// Implementations may decide to reject invalid keys.
	//
	// err is null when the metadata could be written successfully.
	WriteInfo(info TileInfo) error
}
```

### TileInfo

```go
type TileInfo struct {
  //
	Id string
	
  //
  Name string
	
  //
  Description string
	Version string
	Legend string
	Scheme string

	MinZoom int32
	MaxZoom int32

	Bounds string
	Center string
}
```

