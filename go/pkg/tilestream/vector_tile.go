package tilestream

import (
	"bytes"
	"compress/gzip"
	"errors"

	"github.com/mojo-lang/geom/go/pkg/mojo/geom"
	"google.golang.org/protobuf/proto"
)

var (
	DefaultVersion  = uint32(2)
	DefaultExtent   = uint32(4096)
	DefaultLayerTag = "@layer"
)

func NewVectorTileFrom(geojson *geom.GeoJson, id *TileId) *VectorTile {
	tile := &VectorTile{}
	if err := tile.EncodeFrom(geojson, id); err != nil {
		return nil
	}
	return tile
}

// MarshalCompressed output mvt format
func (x *VectorTile) MarshalCompressed(compress bool) ([]byte, error) {
	if x != nil {
		bs, err := proto.Marshal(x)
		if err != nil {
			return nil, err
		}

		if !compress {
			return bs, nil
		}

		bf := bytes.Buffer{}
		w := gzip.NewWriter(&bf)
		_, err = w.Write(bs)
		if err != nil {
			return nil, err
		}
		w.Flush()
		w.Close()

		return bf.Bytes(), nil
	}
	return nil, nil
}

// EncodeFrom encodes vector-tile from geom.GeoBuf
func (x *VectorTile) EncodeFrom(geojson *geom.GeoJson, id *TileId) error {
	feats, err := geojson.ToFeatures()
	if err != nil {
		return err
	}

	groups := groupFeatures(feats)
	x.Layers = make([]*VectorTile_Layer, 0, len(groups))
	for name, fs := range groups {
		layer := EncodeLayer(name, fs, id, DefaultExtent)
		if layer != nil {
			x.Layers = append(x.Layers, layer)
		}
	}

	return nil
}

func groupFeatures(src []*geom.Feature) map[string][]*geom.Feature {
	groups := make(map[string][]*geom.Feature, len(src))
	for _, f := range src {
		name, _ := f.GetStringProperty(DefaultLayerTag)
		groups[name] = append(groups[name], f)
	}
	return groups
}

// Decode decodes vector-tile to geom.GeoBuf
func (x *VectorTile) Decode(tid *TileId) (*geom.GeoJson, error) {
	var features []*geom.Feature
	for _, layer := range x.Layers {
		feats := layer.Decode(tid)
		features = append(features, feats...)
	}

	if len(features) == 0 {
		return nil, errors.New("invalid tile")
	}

	if len(features) == 1 {
		return geom.NewGeoJson(features[0]), nil
	}
	return geom.NewFeatureCollectionGeoJson(features...), nil
}

//
//func DecodeGeometry(cmds []uint32, gtype VectorTile_GeomType, id *TileId, extent float64) *geom.Geometry {
//	var geometry *geom.Geometry
//	z, x, y := id.Level, id.X, id.Y
//	switch gtype {
//	case Tile_UNKNOWN:
//		_ = gtype
//	case Tile_POINT:
//		points := decodePoints(cmds, z, x, y, extent)
//		if len(points) == 1 {
//			geometry = geom.NewGeometry(&geom.Point{Coordinates: points[0]})
//		} else if len(points) > 1 {
//			geometry = geom.NewGeometry(&geom.MultiPoint{Points: points})
//		}
//	case Tile_LINESTRING:
//		lines := decodeLineStrings(cmds, z, x, y, extent)
//		if len(lines) == 1 {
//			geometry = geom.NewGeometry(lines[0])
//		} else if len(lines) > 1 {
//			geometry = geom.NewGeometry(&geom.MultiLineString{LineStrings: lines})
//		}
//	case Tile_POLYGON:
//		polygons := decodePolygons(cmds, z, x, y, extent)
//		if len(polygons) == 1 {
//			geometry = geom.NewGeometry(polygons[0])
//		} else if len(polygons) > 1 {
//			geometry = geom.NewGeometry(&geom.MultiPolygon{Polygons: polygons})
//		}
//	}
//
//	return geometry
//}
