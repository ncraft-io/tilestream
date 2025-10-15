package copy

import "github.com/mojo-lang/geom/go/pkg/mojo/geom"

type Config struct {
	ReaderThreads  int           `json:"readerThreads,omitempty"`
	DuplicatedSkip bool          `json:"duplicatedSkip,omitempty"`
	Tiles          []*TileConfig `json:"tiles,omitempty"`
}

type TileConfig struct {
	MinLevel    int                    `json:"minLevel,omitempty"`
	MaxLevel    int                    `json:"maxLevel,omitempty"`
	Boundary    geom.Geometry          `json:"boundary,omitempty"`
	Source      map[string]interface{} `json:"source,omitempty"`
	Destination map[string]interface{} `json:"destination,omitempty"`
}
