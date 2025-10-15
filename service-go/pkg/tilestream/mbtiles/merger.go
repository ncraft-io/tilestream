package mbtiles

import (
	"context"
	"fmt"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
	"os"
)

func Merge(src string, dest string) error {
	var files []string

	if stat, err := os.Stat(src); err != nil {
		return err
	} else if stat.IsDir() {

	} else if IsMbtilesFile(src) {
		files = append(files, src)
	} else {
		return fmt.Errorf("invalid src: %s", src)
	}

	d := NewFromFile(dest)
	var info *tilestream.TileInfo
	for _, file := range files {
		f := NewFromFile(file)
		i, err := f.Info(context.Background())
		if err != nil {
			continue
		}

		if info == nil {
			info = i
		} else {
			info.Merge(i)
		}

		for l := i.MinZoom; l <= i.MaxZoom; l++ {
			lb := tilestream.GetTileId(i.Bounds.LeftBottom.Longitude, i.Bounds.LeftBottom.Latitude, l)
			rt := tilestream.GetTileId(i.Bounds.RightTop.Longitude, i.Bounds.RightTop.Latitude, l)

			for x := tilestream.MinX(lb, rt); x <= tilestream.MaxX(lb, rt); x++ {
				for y := tilestream.MinY(lb, rt); y <= tilestream.MaxY(lb, rt); y++ {
					tile, _, err := f.Tile(context.Background(), x, y, l)
					if err != nil {
						return err
					}
					if len(tile) > 0 {
						err = d.WriteTile(context.Background(), x, y, l, tile)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	if err := d.WriteInfo(context.Background(), info); err != nil {
		return err
	}
	return nil
}
