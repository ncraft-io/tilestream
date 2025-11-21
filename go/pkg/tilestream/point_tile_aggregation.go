package tilestream

import (
	"errors"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"math"
)

var (
	DefaultSamplingRate = math.Sqrt(math.Ln2)
)

func AggregatePointTiles(subTiles []*PointTile, quadKey string, options core.Options) (*PointTile, error) {
	if len(subTiles) == 0 {
		return nil, errors.New("no sub tiles")
	}

	xs := make([]int32, 0, 8*1024)
	ys := make([]int32, 0, 8*1024)

	var (
		minPoints  = 1024
		extent     = float64(1024)
		sampleRate = DefaultSamplingRate
	)
	if v, ok := options["min_point_count"]; ok {
		minPoints = v.(int)
	}

	if v, ok := options["extent"]; ok {
		extent = v.(float64)
	}

	if v, ok := options["sampling_rate"]; ok {
		sampleRate = v.(float64)
	}

	level := len(quadKey)
	zxy := NewFromQuadKey(quadKey)

	for _, sub := range subTiles {
		npoints := len(sub.Xs)
		if npoints <= minPoints {
			xs = append(xs, sub.Xs...)
			ys = append(ys, sub.Ys...)
			continue
		}
		for i := 0; i < npoints; i++ {
			x, y := sub.Xs[i], sub.Ys[i]
			lon, lat := float64(x)/geom.E7, float64(y)/geom.E7
			u, v := LonLat2XY(lon, lat, zxy.X, zxy.Y, zxy.Level)
			if sampled(level, extent, u/DefaultVectorTilePixelCount, v/DefaultVectorTilePixelCount, sampleRate) {
				xs = append(xs, x)
				ys = append(ys, y)
			}
		}
	}

	return &PointTile{
		Id:   quadKey,
		Type: subTiles[0].Type,
		Xs:   xs,
		Ys:   ys,
	}, nil
}

// sample rate: (0, 1]
func sampled(level int, extent float64, u, v float64, sampleRate float64) bool {
	halfSR := 0.5 * sampleRate
	_u := u * extent
	_v := v * extent
	if math.Abs(_u-math.Round(_u)) <= halfSR &&
		math.Abs(_v-math.Round(_v)) <= halfSR {
		return true
	}
	return false
}
