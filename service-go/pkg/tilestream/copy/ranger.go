package copy

import (
	"github.com/mojo-lang/geom/go/pkg/mojo/geom"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
)

type Ranger struct {
	Geometry *geom.Geometry
	Boundary *geom.BoundingBox

	MinLevel int32
	MaxLevel int32

	Current *tilestream.TileId
}

func NewRanger(geometry *geom.Geometry, minLevel, maxLevel int32) *Ranger {
	ranger := &Ranger{
		Geometry: geometry,
		MinLevel: minLevel,
		MaxLevel: maxLevel,
	}

	return ranger
}

func (r *Ranger) GetNext() *tilestream.TileId {
	return nil
}
