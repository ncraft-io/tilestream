package tilestream

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
)

const (
	moveTo    = 1
	lineTo    = 2
	closePath = 7
)

// EncodeGeometry encodes geom.Geometry into commands(moveTo, lineTo, closePath)
func EncodeGeometry(g *geom.Geometry, id *TileId, extent float64) ([]uint32, VectorTile_GeomType) {
	z, x, y := id.Level, id.X, id.Y
	switch vt := g.Geometry.(type) {
	case *geom.Geometry_Point:
		return encodePoints([]*geom.LngLat{vt.Point.Coordinate}, z, x, y, extent)
	case *geom.Geometry_MultiPoint:
		return encodePoints(vt.MultiPoint.GetLngLats(), z, x, y, extent)
	case *geom.Geometry_LineString:
		return encodeLineStrings([]*geom.LineString{vt.LineString}, z, x, y, extent)
	case *geom.Geometry_MultiLineString:
		return encodeLineStrings(vt.MultiLineString.LineStrings, z, x, y, extent)
	case *geom.Geometry_Polygon:
		return encodePolygons([]*geom.Polygon{vt.Polygon}, z, x, y, extent)
	case *geom.Geometry_MultiPolygon:
		return encodePolygons(vt.MultiPolygon.Polygons, z, x, y, extent)
	}
	return nil, VectorTile_GEOM_TYPE_UNKNOWN
}

// DecodeGeometry decodes commands(moveTo, lineTo, closePath) to geom.Geometry
func DecodeGeometry(cmds []uint32, gtype VectorTile_GeomType, id *TileId, extent float64) *geom.Geometry {
	var geometry *geom.Geometry
	z, x, y := id.Level, id.X, id.Y
	switch gtype {
	case VectorTile_GEOM_TYPE_UNKNOWN:
		_ = gtype
	case VectorTile_GEOM_TYPE_POINT:
		points := decodePoints(cmds, z, x, y, extent)
		if len(points) == 1 {
			geometry = geom.NewPointGeometry(points[0])
		} else if len(points) > 1 {
			geometry = geom.NewMultiPointGeometryFrom(points...)
		}
	case VectorTile_GEOM_TYPE_LINESTRING:
		lines := decodeLineStrings(cmds, z, x, y, extent)
		if len(lines) == 1 {
			geometry = geom.NewGeometry(lines[0])
		} else if len(lines) > 1 {
			geometry = geom.NewGeometry(&geom.MultiLineString{LineStrings: lines})
		}
	case VectorTile_GEOM_TYPE_POLYGON:
		polygons := decodePolygons(cmds, z, x, y, extent)
		if len(polygons) == 1 {
			geometry = geom.NewGeometry(polygons[0])
		} else if len(polygons) > 1 {
			geometry = geom.NewGeometry(&geom.MultiPolygon{Polygons: polygons})
		}
	}

	return geometry
}

func encodePoints(points []*geom.LngLat, z, x, y int32, extent float64) ([]uint32, VectorTile_GeomType) {
	c := len(points)
	if c < 1 {
		return nil, VectorTile_GEOM_TYPE_UNKNOWN
	}

	geometry := make([]uint32, 0, 1+2*c)
	geometry = append(geometry, encodeCmdCount(moveTo, c))

	var refPx, refPy int32
	for _, p := range points {
		u, v := LonLat2XY(p.Longitude, p.Latitude, x, y, z)
		px, py := int32(u*extent/DefaultVectorTilePixelCount), int32(v*extent/DefaultVectorTilePixelCount)
		dpx, dpy := px-refPx, py-refPy
		refPx, refPy = px, py
		geometry = append(geometry, zigzagEncode(dpx), zigzagEncode(dpy))
	}
	return geometry, VectorTile_GEOM_TYPE_POINT
}

func decodePoints(geometry []uint32, z, x, y int32, extent float64) []*geom.LngLat {
	c := len(geometry)
	if c < 3 {
		return nil // should not reach here
	}
	_, n := decodeCmdCount(geometry[0])
	points := make([]*geom.LngLat, 0, n)

	var u, v int32

	for i := 0; i < n; i++ {
		j := 2*i + 1 // begins from 1, for geometry[0] == moveTo
		u += zigzagDecode(geometry[j])
		v += zigzagDecode(geometry[j+1])
		lon, lat := XY2LonLat(float64(u)*DefaultVectorTilePixelCount/extent, float64(v)*DefaultVectorTilePixelCount/extent, x, y, z)
		points = append(points, &geom.LngLat{
			Longitude: lon,
			Latitude:  lat,
		})
	}

	return points
}

func encodeLineStrings(lines []*geom.LineString, z, x, y int32, extent float64) ([]uint32, VectorTile_GeomType) {
	var geometry []uint32
	var refPx, refPy int32
	for _, line := range lines {
		var gline []uint32
		gline, refPx, refPy = encodeLine(line, refPx, refPy, z, x, y, extent)
		if len(gline) > 0 {
			geometry = append(geometry, gline...)
		}
	}
	if len(geometry) == 0 {
		return nil, VectorTile_GEOM_TYPE_UNKNOWN
	}
	return geometry, VectorTile_GEOM_TYPE_LINESTRING
}

func decodeLineStrings(geometry []uint32, z, x, y int32, extent float64) []*geom.LineString {
	var (
		startPos     int
		baseU, baseV int32
		lines        []*geom.LineString
	)

	n := len(geometry)
	for startPos < n {
		coords, offset := lineCoords(geometry[startPos:], baseU, baseV)
		lines = append(lines, &geom.LineString{
			Coordinates: decodeCoords(coords, z, x, y, extent),
		})
		startPos += offset
		nc := len(coords)
		if nc < 2 {
			break // should not reach here
		}
		baseU = coords[nc-2]
		baseV = coords[nc-1]
	}

	return lines
}

func encodePolygons(polygons []*geom.Polygon, z, x, y int32, extent float64) ([]uint32, VectorTile_GeomType) {
	var rings []*geom.LineString
	for _, p := range polygons {
		for i, r := range p.LineStrings {
			cw := clockwise(r)
			if (i == 0 && !cw) ||
				(i != 0 && cw) {
				r = invert(r)
			}
			rings = append(rings, r)
		}
	}
	geometry := encodeRings(rings, z, x, y, extent)
	if len(geometry) == 0 {
		return nil, VectorTile_GEOM_TYPE_UNKNOWN
	}

	return geometry, VectorTile_GEOM_TYPE_POLYGON
}

func decodePolygons(geometry []uint32, z, x, y int32, extent float64) []*geom.Polygon {
	rings := decodeRings(geometry, z, x, y, extent)

	polygons := make([]*geom.Polygon, 0, len(rings))
	var polygon *geom.Polygon

	for _, r := range rings {
		if clockwise(r) {
			// mapbox use closewise ring as exterior ring
			polygon = &geom.Polygon{}
			polygon.LineStrings = append(polygon.LineStrings, r)
			polygons = append(polygons, polygon)
		} else {
			// interior ring
			if polygon == nil {
				polygon = &geom.Polygon{}
			}
			polygon.LineStrings = append(polygon.LineStrings, r)
		}
	}

	return polygons
}

func encodeRings(rings []*geom.LineString, z, x, y int32, extent float64) []uint32 {
	var geometry []uint32
	var refPx, refPy int32
	for _, ring := range rings {
		var gring []uint32
		gring, refPx, refPy = encodeRing(ring, refPx, refPy, z, x, y, extent)
		if len(gring) > 0 {
			geometry = append(geometry, gring...)
		}
	}
	return geometry
}

func decodeRings(geometry []uint32, z, x, y int32, extent float64) []*geom.LineString {
	var (
		startPos     int
		baseU, baseV int32
		rings        []*geom.LineString // use line-string to represent ring
	)

	n := len(geometry)
	for startPos < n {
		coords, offset := ringCoords(geometry[startPos:], baseU, baseV)
		rings = append(rings, &geom.LineString{
			Coordinates: decodeCoords(coords, z, x, y, extent),
		})
		startPos += offset
		nc := len(coords)
		if nc < 4 {
			break // at least 4 points: [a, b, c, a]
		}
		baseU = coords[nc-4]
		baseV = coords[nc-3]
	}

	return rings
}

func encodeLine(line *geom.LineString, pxRef, pyRef int32, z, x, y int32, extent float64) (geometry []uint32, refPx, refPy int32) {
	refPx, refPy = pxRef, pyRef
	if line == nil || len(line.Coordinates) < 2 {
		return
	}

	c := len(line.Coordinates)
	geometry = make([]uint32, 0, c*2+2) // [moveTo, lineTo, n*[x, y]]
	geometry = append(geometry, encodeCmdCount(moveTo, 1))
	for i, p := range line.Coordinates {
		u, v := LonLat2XY(p.Longitude, p.Latitude, x, y, z)
		px, py := int32(u*extent/DefaultVectorTilePixelCount), int32(v*extent/DefaultVectorTilePixelCount)
		dpx, dpy := px-refPx, py-refPy
		refPx, refPy = px, py
		geometry = append(geometry, zigzagEncode(dpx), zigzagEncode(dpy))
		if i == 0 {
			geometry = append(geometry, encodeCmdCount(lineTo, c-1))
		}
	}
	return
}

func lineCoords(geometry []uint32, baseU, baseV int32) (coords []int32, offset int) {
	c := len(geometry)
	if c < 6 { // at least [moveTo, x, y, lineTo, x1, y1]
		return nil, -1 // should not reach here
	}

	u, v := baseU, baseV
	coords = make([]int32, 2, 8)
	for offset < c {
		cmd, n := decodeCmdCount(geometry[offset])
		switch cmd {
		case moveTo:
			offset++
			for i := 0; i < n; i++ {
				u += zigzagDecode(geometry[offset])
				v += zigzagDecode(geometry[offset+1])
				coords[0], coords[1] = u, v
				offset += 2
			}
		case lineTo:
			offset++
			for i := 0; i < n; i++ {
				u += zigzagDecode(geometry[offset])
				v += zigzagDecode(geometry[offset+1])
				coords = append(coords, u, v)
				offset += 2
			}
			return
		}
	}

	return
}

func encodeRing(ring *geom.LineString, pxRef, pyRef int32, z, x, y int32, extent float64) (geometry []uint32, refPx, refPy int32) {
	refPx, refPy = pxRef, pyRef
	if ring == nil || len(ring.Coordinates) < 4 {
		return
	}

	c := len(ring.Coordinates) - 1
	geometry = make([]uint32, 0, c*2+3) // [moveTo, lineTo, n*[x, y], closePath]
	geometry = append(geometry, encodeCmdCount(moveTo, 1))
	for i, p := range ring.Coordinates[:c] {
		u, v := LonLat2XY(p.Longitude, p.Latitude, x, y, z)
		px, py := int32(u*extent/DefaultVectorTilePixelCount), int32(v*extent/DefaultVectorTilePixelCount)
		dpx, dpy := px-refPx, py-refPy
		refPx, refPy = px, py
		geometry = append(geometry, zigzagEncode(dpx), zigzagEncode(dpy))
		if i == 0 {
			geometry = append(geometry, encodeCmdCount(lineTo, c-1))
		}
	}
	geometry = append(geometry, encodeCmdCount(closePath, 1))
	return
}

func ringCoords(geometry []uint32, baseU, baseV int32) (coords []int32, offset int) {
	c := len(geometry)
	if c < 9 { // at least [moveTo, x, y, lineTo, x1, y1, x2, y2, closePath]
		return nil, -1 // should not reach here
	}

	u, v := baseU, baseV
	coords = make([]int32, 2, 8)
	for offset < c {
		cmd, n := decodeCmdCount(geometry[offset])
		switch cmd {
		case moveTo:
			offset++
			for i := 0; i < n; i++ {
				u += zigzagDecode(geometry[offset])
				v += zigzagDecode(geometry[offset+1])
				coords[0], coords[1] = u, v
				offset += 2
			}
		case lineTo:
			offset++
			for i := 0; i < n; i++ {
				u += zigzagDecode(geometry[offset])
				v += zigzagDecode(geometry[offset+1])
				coords = append(coords, u, v)
				offset += 2
			}
		case closePath:
			offset++
			coords = append(coords, coords[0], coords[1])
			return
		}
	}

	return
}

func decodeCoords(coords []int32, z, x, y int32, extent float64) []*geom.LngLat {
	n := len(coords) / 2
	points := make([]*geom.LngLat, 0, n)
	for i := 0; i < n; i++ {
		j := i * 2
		ui := coords[j]
		vi := coords[j+1]
		lon, lat := XY2LonLat(float64(ui)*DefaultVectorTilePixelCount/extent, float64(vi)*DefaultVectorTilePixelCount/extent, x, y, z)
		points = append(points, &geom.LngLat{
			Longitude: lon,
			Latitude:  lat,
		})
	}
	return points
}

func clockwise(ring *geom.LineString) bool {
	return area(ring) > 0
}

// shoelace algorithm
func area(ring *geom.LineString) float64 {
	cs := ring.Coordinates
	n := len(cs) - 1
	xs := make([]float64, 0, n)
	ys := make([]float64, 0, n)
	for _, c := range cs[:n] {
		xs = append(xs, c.Longitude)
		ys = append(ys, c.Latitude)
	}
	return _area(xs, ys, n)
}

func _area(xs, ys []float64, n int) float64 {
	var sum float64
	var i = 0
	var j = n - 1
	for ; i < n; i++ {
		sum += (ys[j] - ys[i]) * (xs[i] + xs[j])
		j = i
	}

	return sum / 2
}

func invert(ring *geom.LineString) *geom.LineString {
	n := len(ring.Coordinates)
	coords := make([]*geom.LngLat, 0, n)
	for i := n - 1; i >= 0; i-- {
		coords = append(coords, ring.Coordinates[i])
	}

	return &geom.LineString{
		Coordinates: coords,
	}
}

func encodeCmdCount(cmd, count int) uint32 {
	return uint32(count<<3) | uint32(cmd&0x07)
}

func decodeCmdCount(v uint32) (cmd, count int) {
	return int(v & 0x07), int(v >> 3)
}

func zigzagEncode(u int32) uint32 {
	return uint32((u << 1) ^ (u >> 31))
}

func zigzagDecode(x uint32) int32 {
	y := int32(x)
	return (y >> 1) ^ (-(y & 1))
}
