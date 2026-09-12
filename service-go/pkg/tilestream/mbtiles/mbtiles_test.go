package mbtiles

import (
	"bytes"
	"context"
	"path/filepath"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
)

func TestMetadataRoundTrip(t *testing.T) {
	info := ToTileInfo([]*Metadata{
		nil, {Name: "name", Value: "world"}, {Name: "format", Value: "png"},
		{Name: "description", Value: "description"}, {Name: "attribution", Value: "source"},
		{Name: "bounds", Value: "110,30,111,31"}, {Name: "center", Value: "110.5,30.5,5"},
		{Name: "minzoom", Value: "0"}, {Name: "maxzoom", Value: "5"},
	})
	if info == nil || info.Name != "world" || info.Format != "png" || info.Center == nil || info.Center.Latitude != 30.5 || info.Bounds == nil || info.MaxZoom != 5 {
		t.Fatalf("info=%v", info)
	}
	info.MaxZoom = 0
	got := ToTileInfo(FromTileInfo(info))
	if got.Name != info.Name || got.Format != info.Format || got.Description != info.Description || got.Attribution != info.Attribution || got.Center == nil || got.Bounds == nil {
		t.Fatalf("metadata not preserved: %v", got)
	}
	foundZero := false
	for _, item := range FromTileInfo(info) {
		if item.Name == "maxzoom" && item.Value == "0" {
			foundZero = true
		}
	}
	if !foundZero {
		t.Fatal("zoom zero metadata lost")
	}
	if FromTileInfo(nil) != nil {
		t.Fatal("nil info should produce no metadata")
	}
}

func TestXYZReadWriteCreatesNewDatabase(t *testing.T) {
	ctx := context.Background()
	file := filepath.Join(t.TempDir(), "fresh.mbtiles")
	stream := NewFromFile(file)
	if err := stream.WriteTile(ctx, 3, 7, 4, []byte("original"), nil); err != nil {
		t.Fatal(err)
	}
	if err := stream.WriteTile(ctx, 3, 7, 4, []byte("replacement"), nil); err != nil {
		t.Fatal(err)
	}
	db := openResult(t, file)
	// XYZ (3,7) at z=4 is TMS column=3,row=8, not (7,3).
	assertStoredTile(t, db, 4, 3, 8, []byte("replacement"))
	tile, _, err := stream.Tile(ctx, 3, 7, 4)
	if err != nil || !bytes.Equal(tile, []byte("replacement")) {
		t.Fatalf("read tile=%q err=%v", tile, err)
	}
	tile, _, err = stream.Tile(ctx, 2, 7, 4)
	if err != nil || len(tile) != 0 {
		t.Fatalf("missing tile=%q err=%v", tile, err)
	}
	info := &tilestream.TileInfo{Name: "name", Format: "png", Center: &geom.LngLat{Longitude: 1, Latitude: 2}, MinZoom: 0, MaxZoom: 4}
	if err := stream.WriteInfo(ctx, info); err != nil {
		t.Fatal(err)
	}
	info.Name = "updated"
	if err := stream.WriteInfo(ctx, info); err != nil {
		t.Fatal(err)
	}
	readInfo, err := stream.Info(ctx)
	if err != nil || readInfo == nil || readInfo.Name != "updated" || readInfo.Format != "png" || readInfo.Center == nil {
		t.Fatalf("info=%v err=%v", readInfo, err)
	}
}

func TestCachesSeparateSameBasename(t *testing.T) {
	ctx := context.Background()
	dir := t.TempDir()
	src, dest := filepath.Join(dir, "src", "map.mbtiles"), filepath.Join(dir, "dest", "map.mbtiles")
	source := fixtureDB(t, src, pngMetadata(""), Tiles{ZoomLevel: 5, TileColumn: 25, TileRow: 18, TileData: []byte("source")})
	destination := fixtureDB(t, dest, pngMetadata(""))
	left, right := NewFromFile(src), NewFromFile(dest)
	data, _, err := left.Tile(ctx, 25, 13, 5)
	if err != nil || string(data) != "source" {
		t.Fatalf("source tile=%q err=%v", data, err)
	}
	if _, err := left.Info(ctx); err != nil {
		t.Fatal(err)
	}
	if err := right.WriteTile(ctx, 4, 4, 4, []byte("destination"), nil); err != nil {
		t.Fatal(err)
	}
	if err := right.WriteInfo(ctx, &tilestream.TileInfo{Name: "destination", Format: "png"}); err != nil {
		t.Fatal(err)
	}
	assertStoredTile(t, destination, 4, 4, 11, []byte("destination"))
	var count int
	if err := source.QueryRow(`SELECT COUNT(*) FROM tiles WHERE zoom_level=4`).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatal("source database was modified")
	}
	sourceInfo, err := left.Info(ctx)
	if err != nil || sourceInfo.Name != "test" {
		t.Fatalf("source metadata modified: %v %v", sourceInfo, err)
	}
	destInfo, err := right.Info(ctx)
	if err != nil || destInfo.Name != "destination" {
		t.Fatalf("destination metadata=%v %v", destInfo, err)
	}
}

func TestReadAndOpenErrorsPropagate(t *testing.T) {
	ctx := context.Background()
	for _, name := range []string{"missing-tables.mbtiles", "missing-parent/file.mbtiles", "invalid.txt"} {
		t.Run(name, func(t *testing.T) {
			stream := NewFromFile(filepath.Join(t.TempDir(), name))
			if _, err := stream.Info(ctx); err == nil {
				t.Fatal("Info must return the read/open error")
			}
			if _, _, err := stream.Tile(ctx, 0, 0, 0); err == nil {
				t.Fatal("Tile must return the read/open error")
			}
		})
	}
	stream := NewFromFile(filepath.Join(t.TempDir(), "nil-info.mbtiles"))
	if err := stream.WriteInfo(ctx, nil); err == nil {
		t.Fatal("nil info must return an error")
	}
}

func TestInvalidCoordinatesReturnErrors(t *testing.T) {
	stream := NewFromFile(filepath.Join(t.TempDir(), "invalid.mbtiles"))
	for _, xyz := range [][3]int32{{0, 0, -1}, {0, 0, 32}, {-1, 0, 1}, {0, -1, 1}, {2, 0, 1}, {0, 2, 1}} {
		if _, _, err := stream.Tile(context.Background(), xyz[0], xyz[1], xyz[2]); err == nil {
			t.Errorf("read accepted %v", xyz)
		}
		if err := stream.WriteTile(context.Background(), xyz[0], xyz[1], xyz[2], []byte("bad"), nil); err == nil {
			t.Errorf("write accepted %v", xyz)
		}
	}
}
