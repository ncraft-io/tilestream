package mbtiles

import (
	"fmt"
	"github.com/mojo-lang/geom/go/pkg/mojo/geom"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
	"strconv"
	"strings"
)

type Metadata struct {
	Name  string `gorm:"uniqueIndex"`
	Value string
}

func convertPoints(array []string) []float64 {
	var points []float64
	for _, element := range array {
		point, err := strconv.ParseFloat(strings.TrimSpace(element), 64)
		if err != nil {
			return nil
		}
		points = append(points, point)
	}
	return points
}

func ToTileInfo(data []*Metadata) *tilestream.TileInfo {
	info := &tilestream.TileInfo{}

	for _, d := range data {
		switch d.Name {
		case "bounds":
			segments := strings.Split(d.Value, ",")
			if len(segments) == 4 {
				points := convertPoints(segments)
				if len(points) == 4 {
					info.Bounds = &geom.BoundingBox{
						LeftBottom: &geom.LngLat{
							Longitude: points[0],
							Latitude:  points[1],
						},
						RightTop: &geom.LngLat{
							Longitude: points[2],
							Latitude:  points[3],
						},
					}
				}
			}
		case "center":
			segments := strings.Split(d.Value, ",")
			if len(segments) == 2 {
				points := convertPoints(segments)
				if len(points) == 2 {
					info.Center = &geom.LngLat{
						Longitude: points[0],
						Latitude:  points[1],
					}
				}
			}
		case "minzoom":
			if zoom, err := strconv.ParseInt(strings.TrimSpace(d.Value), 10, 64); err == nil {
				info.MinZoom = int32(zoom)
			}
		case "maxzoom":
			if zoom, err := strconv.ParseInt(strings.TrimSpace(d.Value), 10, 64); err == nil {
				info.MaxZoom = int32(zoom)
			}
		}
	}
	return nil
}

func FromTileInfo(info *tilestream.TileInfo) []*Metadata {
	var data []*Metadata
	if info.Bounds != nil && info.Bounds.LeftBottom != nil && info.Bounds.RightTop != nil {
		bounds := info.Bounds
		data = append(data, &Metadata{
			Name:  "bounds",
			Value: fmt.Sprintf("%.6f,%.6f,%.6f,%.6f", bounds.LeftBottom.Longitude, bounds.LeftBottom.Latitude, bounds.RightTop.Longitude, bounds.RightTop.Latitude),
		})
	}
	if info.Center != nil {
		data = append(data, &Metadata{
			Name:  "center",
			Value: fmt.Sprintf("%.6f,%.6f", info.Center.Longitude, info.Center.Latitude),
		})
	}

	if info.MinZoom >= 0 && info.MaxZoom > 0 {
		data = append(data, &Metadata{
			Name:  "minzoom",
			Value: fmt.Sprintf("%d", info.MinZoom),
		}, &Metadata{
			Name:  "maxzoom",
			Value: fmt.Sprintf("%d", info.MaxZoom),
		})
	}
	return data
}
