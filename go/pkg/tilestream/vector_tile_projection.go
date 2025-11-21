package tilestream

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"math"
)

const (
	DefaultVectorTilePixelCount = 256
)

/*
 *
 * n = 2 ^ zoom
 * xtile = n * ((lon_deg + 180) / 360)
 * ytile = n * (1 - (log(tan(lat_rad) + sec(lat_rad)) / pi)) / 2
 *       = n * (1 - (log((1+sin(lat_rad)/(1-sin(lat_rad))/pi)/4)
 *
 */

// GetX gets x
func GetX(lon float64, z int32) int32 {
	n := 1 << z // 2^z
	x := int32(math.Floor(float64(n) * (lon + 180) / 360))
	return _clip(x, 0, int32(n-1))
}

// GetY gets y
func GetY(lat float64, z int32) int32 {
	lat = Clip(lat, LatitudeMin, LatitudeMax)
	n := 1 << z // 2^z
	rad := lat * math.Pi / 180.0
	y := int32(math.Floor((1 - math.Log(math.Tan(rad)+1/math.Cos(rad))/math.Pi) / 2.0 * float64(n)))
	return _clip(y, 0, int32(n-1))
}

func _clip(v, min, max int32) int32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}

	return v
}

// GetLon x to lon
func GetLon(x float64, z int32) float64 {
	n := 1 << z // 2^z
	return x/float64(n)*360 - 180
}

// GetLat y to lat
func GetLat(y float64, z int32) float64 {
	n := 1 << z // 2^z
	return math.Atan(math.Sinh(math.Pi*(1-2*y/float64(n)))) * 180 / math.Pi
}

// PointToXY ...
func PointToXY(point *geom.LngLat, tid *TileId) (float64, float64) {
	return LonLat2XY(point.Longitude, point.Latitude, tid.X, tid.Y, tid.Level)
}

// LonLat2XY converts a lonlat to a point x/y for the specified map tile.
func LonLat2XY(lon, lat float64, tileX, tileY, tileZ int32) (float64, float64) {
	lon = Clip(lon, LongitudeMin, LongitudeMax)
	lat = Clip(lat, LatitudeMin, LatitudeMax)

	lx := (lon + 180) / 360
	sinLat := math.Sin(lat * math.Pi / 180)
	ly := 0.5 - math.Log((1+sinLat)/(1-sinLat))/(4*math.Pi)
	mapSize := float64(uint64(DefaultVectorTilePixelCount) << uint(tileZ))
	pixelX := Clip(lx*mapSize, -mapSize/2, mapSize+mapSize/2-1)
	pixelY := Clip(ly*mapSize, -mapSize/2, mapSize+mapSize/2-1)
	return pixelX - float64(tileX*DefaultVectorTilePixelCount), pixelY - float64(tileY*DefaultVectorTilePixelCount)
}

// XY2LonLat converts an point x/y to lon/lat
func XY2LonLat(x, y float64, tileX, tileY, tileZ int32) (lon, lat float64) {
	lon = GetLon(float64(tileX)+x/DefaultVectorTilePixelCount, tileZ)
	lat = GetLat(float64(tileY)+y/DefaultVectorTilePixelCount, tileZ)
	return
}
