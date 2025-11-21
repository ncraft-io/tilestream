package copy

import (
	"context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	ts "github.com/ncraft-io/tilestream/service-go/pkg/tilestream"
)

// Copy
// tiles
//
//	minLevel: 8
//	maxLevel: 12
//	boundary: // geometry
//	    type: Polygon
//	source:
//	destination:
func Copy(src, dest string, geometry *geom.Geometry, minLevel, maxLevel int32) (int, error) {
	ss := ts.NewTileStream(src)
	ds := ts.NewTileStream(dest)

	ranger := NewRanger(geometry, minLevel, maxLevel)
	count := 0
	for {
		tid := ranger.GetNext()
		if tid == nil {
			break
		}

		tile, _, err := ss.Tile(context.Background(), tid.X, tid.Y, tid.Level)
		if err != nil {
			logs.Warnw("failed to read the tile from the src tilestream", "tid", tid, "src", src)
			continue
		}

		err = ds.WriteTile(context.Background(), tid.X, tid.Y, tid.Level, tile)
		if err != nil {
			logs.Warnw("failed to write the tile to dest tilestream", "tid", tid, "dest", dest)
			continue
		}

		count++
	}

	return count, nil
}

func ParallelCopy() {

}
