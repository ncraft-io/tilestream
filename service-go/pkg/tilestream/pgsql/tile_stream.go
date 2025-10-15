package pgsql

import (
	"github.com/mojo-lang/db/go/pkg/mojo/db"
)

type TileStream struct {
	DB *db.DB
}

//func NewTileStream() *TileStream {
//	ts := &TileStream{}
//
//	tx := ts.DB.Select("Select postgis_version()").Scan()
//}
//
//type Mvt struct {
//	Tile []byte `json:"tile"`
//}
//
//func (s *TileStream) Tile(ctx context.Context, x, y, level int32) ([]byte, *core.Options, error) {
//
//	sql := fmt.Sprintf(`WITH mvtgeom AS (
//		SELECT ST_AsMVTGeom(geom, ST_TileEnvelope(%d, %d, %d), extent => 4096, buffer => 64) AS geom %s
//		FROM %s
//		WHERE geom && ST_TileEnvelope(%d, %d, %d, margin => (64.0 / 4096)))
//	SELECT ST_AsMVT(mvtgeom.*) as tile
//	FROM mvtgeom;`, level, x, y, fields, table, level, x, y)
//
//	mvt := &Mvt{}
//	s.DB.Raw(sql).Scan(mvt)
//
//	return nil, nil, nil
//}
//
//func (s *TileStream) Info(ctx context.Context) (*tilestream.TileInfo, error) {
//	return nil, nil
//}
