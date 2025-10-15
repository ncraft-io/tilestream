package tilestream

import (
	"math"

	"github.com/golang/geo/s2"
	"github.com/mojo-lang/geom/go/pkg/mojo/geom"
)

const (
	EarthRadius = 6378137.0
	LatitudeMax = 85.051128779806604
	LatitudeMin = -85.051128779806604

	LongitudeMax = 180.0
	LongitudeMin = -180.0

	DegreeMax = 360.0

	Deg2Rad = math.Pi / 180.0
	Rad2Deg = 180.0 / math.Pi
)

func NewTileId(x, y, level int32) *TileId {
	return &TileId{
		X:     x,
		Y:     y,
		Level: level,
	}
}

func NewFromQuadKey(s string) *TileId {
	z := int32(len(s))
	id := &TileId{Level: z}

	for i := id.Level; i > 0; i-- {
		mask := int32(1) << uint(i-1)
		switch s[z-i] {
		case '0':
		case '1':
			id.X |= mask
		case '2':
			id.Y |= mask
		case '3':
			id.X |= mask
			id.Y |= mask
		default:
			panic("invalid string " + s)
		}
	}

	return id
}

func (x *TileId) Vertex(index int32) *geom.LngLat {
	switch index {
	case 0:
		return TopLeftPoint(x)
	case 1:
		return TopLeftPoint(&TileId{X: x.X, Y: x.Y + 1, Level: x.Level})
	case 2:
		return TopLeftPoint(&TileId{X: x.X + 1, Y: x.Y + 1, Level: x.Level})
	case 3:
		return TopLeftPoint(&TileId{X: x.X + 1, Y: x.Y, Level: x.Level})
	default:
		return nil
	}
}

func (x *TileId) Polygon() *geom.Polygon {
	polygon := &geom.Polygon{}
	line := &geom.LineString{}
	line.Coordinates = append(line.Coordinates, x.Vertex(0))
	line.Coordinates = append(line.Coordinates, x.Vertex(1))
	line.Coordinates = append(line.Coordinates, x.Vertex(2))
	line.Coordinates = append(line.Coordinates, x.Vertex(3))
	line.Coordinates = append(line.Coordinates, line.Coordinates[0])
	polygon.LineStrings = []*geom.LineString{line}
	return polygon
}

func (x *TileId) CellCovering(minLevel, maxLevel, maxCells int) s2.CellUnion {
	return x.Polygon().CellCovering(minLevel, maxLevel, maxCells)
}

func (x *TileId) SubTiles() []TileId {
	return x.SubLevelTiles(1)
}

func (x *TileId) SubLevelTiles(sublevel int) []TileId {
	total := int(math.Pow(2, float64(sublevel)))

	tiles := make([]TileId, total*total)
	originX := x.X * int32(total)
	originY := x.Y * int32(total)

	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			tiles[total*j+i].X = originX + int32(i)
			tiles[total*j+i].Y = originY + int32(j)
			tiles[total*j+i].Level = x.Level + int32(sublevel)
		}
	}

	return tiles
}

/*
func (t TileId) HexKey() string {
	if t.Level == 0 {

	} else {
		for i := t.Level; i > 0; i = i - 2 {

		}
	}
}*/

func (x *TileId) QuadKey() string {
	key := ""
	for i := x.Level; i > 0; i-- {
		digit := '0'
		mask := int32(1) << uint(i-1)
		if x.X&mask != 0 {
			digit++
		}
		if x.Y&mask != 0 {
			digit += 2
		}
		key += string(digit)
	}

	return key
}

func (x *TileId) ResetQuadKey(key string) {
	id := NewFromQuadKey(key)
	x.Level = id.Level
	x.X = id.X
	x.Y = id.Y
}

func QuadKeyLevel(key string) int {
	return len(key)
}

func GetTileId(longitude, latitude float64, level int32) *TileId {
	lat := Clip(latitude, LatitudeMin, LatitudeMax)
	lng := Clip(longitude, LongitudeMin, LongitudeMax)

	x := (lng + 180) / 360
	r := lat * Deg2Rad
	y := 0.5 - math.Log(math.Tan(r)+1/math.Cos(r))/(2.0*math.Pi)

	size := float64(uint32(1) << uint32(level))
	tileX := int32(Clip(x*size, 0, size-1))
	tileY := int32(Clip(y*size, 0, size-1))

	return &TileId{X: tileX, Y: tileY, Level: level}
}

func TopLeftPoint(id *TileId) *geom.LngLat {
	if id != nil {
		n := math.Pi - 2.0*math.Pi*float64(id.Y)/math.Pow(2.0, float64(id.Level))
		longitude := float64(id.X)/math.Pow(2.0, float64(id.Level))*DegreeMax - LongitudeMax
		latitude := Rad2Deg * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
		return &geom.LngLat{Longitude: longitude, Latitude: latitude}
	}
	return nil
}

func MinX(lb *TileId, rt *TileId) int32 {
	if lb == nil {
		if rt == nil {
			return 0
		} else {
			return rt.X
		}
	} else {
		if rt == nil {
			return lb.X
		} else {
			return min(lb.X, rt.X)
		}
	}
}

func MaxX(lb *TileId, rt *TileId) int32 {
	if lb == nil {
		if rt == nil {
			return 0
		} else {
			return rt.X
		}
	} else {
		if rt == nil {
			return lb.X
		} else {
			return max(lb.X, rt.X)
		}
	}
}

func MinY(lb *TileId, rt *TileId) int32 {
	if lb == nil {
		if rt == nil {
			return 0
		} else {
			return rt.Y
		}
	} else {
		if rt == nil {
			return lb.Y
		} else {
			return min(lb.Y, rt.Y)
		}
	}
}

func MaxY(lb *TileId, rt *TileId) int32 {
	if lb == nil {
		if rt == nil {
			return 0
		} else {
			return rt.Y
		}
	} else {
		if rt == nil {
			return lb.Y
		} else {
			return max(lb.Y, rt.Y)
		}
	}
}

// Clip
// <summary>
// Clips a number to the specified minimum and maximum values.
// </summary>
// <param name="n">The number to Clip.</param>
// <param name="minValue">Minimum allowable value.</param>
// <param name="maxValue">Maximum allowable value.</param>
// <returns>The Clipped value.</returns>
func Clip(n, minValue, maxValue float64) float64 {
	return math.Min(math.Max(n, minValue), maxValue)
}

// GroundResolution
// <summary>
// Determines the ground resolution (in meters per pixel) at a specified
// latitude and level of detail.
// </summary>
// <param name="latitude">Latitude (in degrees) at which to measure the
// ground resolution.</param>
// <param name="levelOfDetail">Level of detail, from 1 (lowest detail)
// to 23 (highest detail).</param>
// <returns>The ground resolution, in meters per pixel.</returns>
func GroundResolution(latitude float64, level int32) float64 {
	return 0
}

// MapScale
// <summary>
// Determines the map scale at a specified latitude, level of detail,
// and screen resolution.
// </summary>
// <param name="latitude">Latitude (in degrees) at which to measure the
// map scale.</param>
// <param name="levelOfDetail">Level of detail, from 1 (lowest detail)
// to 23 (highest detail).</param>
// <param name="screenDpi">Resolution of the screen, in dots per inch.</param>
// <returns>The map scale, expressed as the denominator N of the ratio 1 : N.</returns>
func MapScale(latitude float64, levelOfDetail int32, screenDpi int32) float64 {
	return GroundResolution(latitude, levelOfDetail) * float64(screenDpi) / 0.0254
}

// LngLatToPixelXY
// <summary>
// Converts a point from latitude/longitude WGS-84 coordinates (in degrees)
// into pixel XY coordinates at a specified level of detail.
// </summary>
// <param name="latitude">Latitude of the point, in degrees.</param>
// <param name="longitude">Longitude of the point, in degrees.</param>
// <param name="levelOfDetail">Level of detail, from 1 (lowest detail)
// to 23 (highest detail).</param>
// <param name="pixelX">Output parameter receiving the X coordinate in pixels.</param>
// <param name="pixelY">Output parameter receiving the Y coordinate in pixels.</param>
func LngLatToPixelXY(longitude, latitude float64, level int32) (int32, int32) {
	return 0, 0
}
