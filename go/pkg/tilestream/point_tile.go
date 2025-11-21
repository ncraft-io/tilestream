package tilestream

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
)

func (t *PointTile) PointCount() int {
	return len(t.Xs)
}

func (t *PointTile) lng(index int) float64 {
	return float64(t.Xs[index]) / geom.E7
}

func (t *PointTile) lat(index int) float64 {
	return float64(t.Ys[index]) / geom.E7
}

func (t *PointTile) BinLngLat(index int) *geom.BinLngLat {
	if index >= 0 && index < len(t.Xs) {
		return &geom.BinLngLat{Longitude: int32(t.Xs[index]), Latitude: int32(t.Ys[index])}
	}
	return nil
}

func (t *PointTile) LngLat(index int) *geom.LngLat {
	if index >= 0 && index < len(t.Xs) {
		return &geom.LngLat{Longitude: t.lng(index), Latitude: t.lat(index)}
	}
	return nil
}

func (t *PointTile) Point(index int) *geom.Point {
	if ll := t.LngLat(index); ll != nil {
		return geom.NewPoint(ll)
	}
	return nil
}

func (t *PointTile) BinLngLats() []*geom.BinLngLat {
	var lnglats []*geom.BinLngLat
	for i := 0; i < len(t.Xs); i++ {
		lnglats = append(lnglats, &geom.BinLngLat{Longitude: int32(t.Xs[i]), Latitude: int32(t.Ys[i])})
	}
	return lnglats
}

func (t *PointTile) Properties(index int) map[string]*VectorTile_Value {
	if index >= 0 && index < len(t.Tags)/2 {
	}
	return nil
}

func (t *PointTile) ToMVT(layerName string, compress bool) ([]byte, error) {
	var buf *geom.GeoJson
	npoints := len(t.Xs)
	if len(t.Xs) == len(t.Ids) {
		features := make([]*geom.Feature, 0, npoints)
		for i := 0; i < npoints; i++ {
			point := t.LngLat(i)
			features = append(features, &geom.Feature{
				Id:         core.NewIntId(t.Ids[i]),
				Geometry:   geom.NewPointGeometry(point),
				Properties: map[string]*core.Value{DefaultLayerTag: core.NewStringValue(layerName)},
			})
		}
		buf = geom.NewFeatureCollectionGeoJson(features...)
	} else {
		points := make([]*geom.Point, 0, npoints)
		for i := 0; i < npoints; i++ {
			points = append(points, t.Point(i))
		}

		feature := &geom.Feature{
			Geometry:   geom.NewMultiPointGeometry(points...),
			Properties: map[string]*core.Value{DefaultLayerTag: core.NewStringValue(layerName)},
		}
		buf = geom.NewGeoJson(feature)
	}

	vt := &VectorTile{}
	vt.EncodeFrom(buf, NewFromQuadKey(t.Id))
	return vt.MarshalCompressed(compress)
}

func (t *Tile) SetBinLngLats(lnglats []*geom.BinLngLat) {
}

func (t *Tile) AddLngLats(lngLats ...*geom.LngLat) {
}
