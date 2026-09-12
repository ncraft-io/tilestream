package mbtiles

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
)

type Metadata struct {
	Name  string `gorm:"uniqueIndex"`
	Value string
}

func convertPoints(array []string) []float64 {
	var points []float64
	for _, element := range array {
		point, err := strconv.ParseFloat(strings.TrimSpace(element), 64)
		if err != nil || math.IsNaN(point) || math.IsInf(point, 0) {
			return nil
		}
		points = append(points, point)
	}
	return points
}

func ToTileInfo(data []*Metadata) *tilestream.TileInfo {
	info := &tilestream.TileInfo{}

	for _, d := range data {
		if d == nil {
			continue
		}
		switch d.Name {
		case "name":
			info.Name = d.Value
		case "format":
			info.Format = d.Value
		case "type":
			info.Type = d.Value
		case "description":
			info.Description = d.Value
		case "version":
			info.Version = d.Value
		case "attribution":
			info.Attribution = d.Value
		case "scheme":
			info.Scheme = d.Value
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
			if len(segments) == 2 || len(segments) == 3 {
				points := convertPoints(segments)
				if len(points) >= 2 {
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
	return info
}

func FromTileInfo(info *tilestream.TileInfo) []*Metadata {
	var data []*Metadata
	if info == nil {
		return nil
	}
	for _, item := range []Metadata{
		{Name: "name", Value: info.Name},
		{Name: "format", Value: info.Format},
		{Name: "type", Value: info.Type},
		{Name: "description", Value: info.Description},
		{Name: "version", Value: info.Version},
		{Name: "attribution", Value: info.Attribution},
		{Name: "scheme", Value: info.Scheme},
	} {
		if item.Value != "" {
			data = append(data, &Metadata{Name: item.Name, Value: item.Value})
		}
	}
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
			Value: fmt.Sprintf("%.6f,%.6f,%d", info.Center.Longitude, info.Center.Latitude, info.MinZoom),
		})
	}

	if info.MinZoom >= 0 && info.MaxZoom >= info.MinZoom {
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
