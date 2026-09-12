package mbtiles

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func fixtureDB(t *testing.T, file string, metadata map[string]string, tiles ...Tiles) *sql.DB {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		t.Fatal(err)
	}
	database, err := sql.Open("sqlite3", sqliteDSN(file, "rwc"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { database.Close() })
	// Deliberately omit indexes: they are optional in source MBTiles files.
	for _, statement := range schemaStatements[:2] {
		if _, err := database.Exec(statement); err != nil {
			t.Fatal(err)
		}
	}
	for key, value := range metadata {
		if _, err := database.Exec(`INSERT INTO metadata VALUES (?,?)`, key, value); err != nil {
			t.Fatal(err)
		}
	}
	for _, tile := range tiles {
		if _, err := database.Exec(`INSERT INTO tiles VALUES (?,?,?,?)`, tile.ZoomLevel, tile.TileColumn, tile.TileRow, tile.TileData); err != nil {
			t.Fatal(err)
		}
	}
	return database
}

func pngMetadata(bounds string) map[string]string {
	m := map[string]string{"name": "test", "format": "png", "custom": "preserved"}
	if bounds != "" {
		m["bounds"] = bounds
	}
	return m
}

func openResult(t *testing.T, file string) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", sqliteDSN(file, "ro"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func assertStoredTile(t *testing.T, db *sql.DB, z, x, y int, expected []byte) {
	t.Helper()
	var data []byte
	if err := db.QueryRow(`SELECT tile_data FROM tiles WHERE zoom_level=? AND tile_column=? AND tile_row=?`, z, x, y).Scan(&data); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, expected) {
		t.Fatalf("tile %d/%d/%d = %q, want %q", z, x, y, data, expected)
	}
}

func TestMergeSinglePreservesStorageAndMetadata(t *testing.T) {
	dir := t.TempDir()
	// Same basename in different directories must never write back to the source.
	src, dest := filepath.Join(dir, "src", "map.mbtiles"), filepath.Join(dir, "dest", "map.mbtiles")
	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		t.Fatal(err)
	}
	metadata := pngMetadata("110,30,111,31")
	metadata["center"] = "110.5,30.5,5"
	metadata["minzoom"], metadata["maxzoom"] = "0", "0" // Intentionally stale.
	source := fixtureDB(t, src, metadata,
		Tiles{ZoomLevel: 5, TileColumn: 25, TileRow: 18, TileData: []byte("north")},
		Tiles{ZoomLevel: 8, TileColumn: 1, TileRow: 200, TileData: []byte{}}, // Outside advertised bounds, and empty.
	)
	if err := Merge(src, dest); err != nil {
		t.Fatal(err)
	}
	result := openResult(t, dest)
	assertStoredTile(t, result, 5, 25, 18, []byte("north"))
	assertStoredTile(t, result, 8, 1, 200, []byte{})
	assertStoredTile(t, source, 5, 25, 18, []byte("north"))
	got, err := readMergeMetadata(context.Background(), result)
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"name", "format", "custom", "center", "bounds"} {
		if got[key] != metadata[key] {
			t.Errorf("metadata %s=%q, want %q", key, got[key], metadata[key])
		}
	}
	if got["minzoom"] != "5" || got["maxzoom"] != "8" {
		t.Fatalf("zoom range: %v", got)
	}
}

func TestMergeDirectoryExistingDestination(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "0-output.mbtiles")
	fixtureDB(t, dest, pngMetadata("-10,-10,10,10"), Tiles{ZoomLevel: 0, TileData: []byte("existing")})
	fixtureDB(t, filepath.Join(dir, "a.mbtiles"), pngMetadata("110,30,111,31"), Tiles{ZoomLevel: 5, TileColumn: 25, TileRow: 18, TileData: []byte("first")})
	fixtureDB(t, filepath.Join(dir, "nested", "z.mbtiles"), pngMetadata("120,40,121,41"),
		Tiles{ZoomLevel: 5, TileColumn: 25, TileRow: 18, TileData: []byte("last")},
		Tiles{ZoomLevel: 6, TileColumn: 50, TileRow: 40, TileData: []byte("new")})
	if err := os.WriteFile(filepath.Join(dir, "ignored.txt"), []byte("ignore"), 0644); err != nil {
		t.Fatal(err)
	}
	files, err := mergeFiles(dir, dest)
	if err != nil || len(files) != 2 {
		t.Fatalf("sources=%v, err=%v", files, err)
	}
	if err := Merge(dir, dest); err != nil {
		t.Fatal(err)
	}
	result := openResult(t, dest)
	assertStoredTile(t, result, 0, 0, 0, []byte("existing"))
	assertStoredTile(t, result, 5, 25, 18, []byte("last"))
	assertStoredTile(t, result, 6, 50, 40, []byte("new"))
	got, err := readMergeMetadata(context.Background(), result)
	if err != nil {
		t.Fatal(err)
	}
	if got["bounds"] != "-10,-10,121,41" || got["minzoom"] != "0" || got["maxzoom"] != "6" {
		t.Fatalf("merged metadata=%v", got)
	}
	// Re-running must be deterministic and must not duplicate rows.
	if err := Merge(dir, dest); err != nil {
		t.Fatal(err)
	}
	var count int
	if err := result.QueryRow(`SELECT COUNT(*) FROM tiles`).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 3 {
		t.Fatalf("tile count=%d", count)
	}
}

func TestMergeSourceViewsAndMissingBounds(t *testing.T) {
	dir := t.TempDir()
	src, dest := filepath.Join(dir, "source.mbtiles"), filepath.Join(dir, "dest.mbtiles")
	source := fixtureDB(t, src, pngMetadata(""), Tiles{ZoomLevel: 2, TileColumn: 1, TileRow: 3, TileData: []byte("view")})
	for _, q := range []string{`ALTER TABLE tiles RENAME TO stored_tiles`, `CREATE VIEW tiles AS SELECT * FROM stored_tiles`, `ALTER TABLE metadata RENAME TO stored_metadata`, `CREATE VIEW metadata AS SELECT * FROM stored_metadata`} {
		if _, err := source.Exec(q); err != nil {
			t.Fatal(err)
		}
	}
	fixtureDB(t, dest, pngMetadata("0,0,1,1"), Tiles{ZoomLevel: 0, TileData: []byte("old")})
	if err := Merge(src, dest); err != nil {
		t.Fatal(err)
	}
	result := openResult(t, dest)
	assertStoredTile(t, result, 2, 1, 3, []byte("view"))
	metadata, err := readMergeMetadata(context.Background(), result)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := metadata["bounds"]; ok {
		t.Fatalf("misleading partial bounds retained: %v", metadata)
	}
}

func TestMergeFailureRollsBack(t *testing.T) {
	for _, failure := range []string{"format", "missing tiles", "corrupt", "bad bounds", "bad coordinate", "missing name", "missing pbf json"} {
		t.Run(failure, func(t *testing.T) {
			dir := t.TempDir()
			src, dest := filepath.Join(dir, "sources"), filepath.Join(dir, "dest.mbtiles")
			original := fixtureDB(t, dest, pngMetadata("0,0,1,1"), Tiles{ZoomLevel: 0, TileData: []byte("original")})
			fixtureDB(t, filepath.Join(src, "a.mbtiles"), pngMetadata("0,0,2,2"), Tiles{ZoomLevel: 0, TileData: []byte("replacement")})
			badFile := filepath.Join(src, "z.mbtiles")
			bad := fixtureDB(t, badFile, pngMetadata("0,0,3,3"))
			var query string
			switch failure {
			case "format":
				query = `UPDATE metadata SET value='jpg' WHERE name='format'`
			case "missing tiles":
				query = `DROP TABLE tiles`
			case "corrupt":
				bad.Close()
				if err := os.WriteFile(badFile, []byte("not a database"), 0644); err != nil {
					t.Fatal(err)
				}
			case "bad bounds":
				query = `UPDATE metadata SET value='NaN,0,3,3' WHERE name='bounds'`
			case "bad coordinate":
				query = `INSERT INTO tiles VALUES (1,5,0,X'00')`
			case "missing name":
				query = `DELETE FROM metadata WHERE name='name'`
			case "missing pbf json":
				query = `UPDATE metadata SET value='pbf' WHERE name='format'`
			}
			if query != "" {
				if _, err := bad.Exec(query); err != nil {
					t.Fatal(err)
				}
			}
			if err := Merge(src, dest); err == nil {
				t.Fatal("expected merge failure")
			}
			assertStoredTile(t, original, 0, 0, 0, []byte("original"))
			got, err := readMergeMetadata(context.Background(), original)
			if err != nil || got["bounds"] != "0,0,1,1" {
				t.Fatalf("destination metadata changed: %v %v", got, err)
			}
			newDest := filepath.Join(dir, "new.mbtiles")
			if err := Merge(src, newDest); err == nil {
				t.Fatal("expected merge failure")
			}
			if _, err := os.Stat(newDest); !os.IsNotExist(err) {
				t.Fatalf("failed output left behind: %v", err)
			}
		})
	}
}

func TestMergeRejectsInvalidPaths(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "source.mbtiles")
	fixtureDB(t, src, pngMetadata(""))
	empty := filepath.Join(dir, "empty")
	if err := os.Mkdir(empty, 0755); err != nil {
		t.Fatal(err)
	}
	symlink, hardlink := filepath.Join(dir, "alias.mbtiles"), filepath.Join(dir, "hardlink.mbtiles")
	if err := os.Symlink(src, symlink); err != nil {
		t.Fatal(err)
	}
	if err := os.Link(src, hardlink); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][2]string{{src, src}, {src, symlink}, {src, hardlink}, {empty, filepath.Join(dir, "dest.mbtiles")}, {src, ""}, {src, filepath.Join(dir, "bad.db")}, {"", filepath.Join(dir, "dest.mbtiles")}} {
		if err := Merge(args[0], args[1]); err == nil {
			t.Errorf("Merge(%q,%q) should fail", args[0], args[1])
		}
	}
	files, err := mergeFiles(dir, src)
	if err == nil || len(files) != 0 {
		t.Fatalf("destination aliases should be excluded: %v %v", files, err)
	}
}

func TestMergeCancelled(t *testing.T) {
	dir := t.TempDir()
	src, dest := filepath.Join(dir, "source.mbtiles"), filepath.Join(dir, "dest.mbtiles")
	fixtureDB(t, src, pngMetadata(""), Tiles{TileData: []byte("tile")})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := MergeContext(ctx, src, dest); err == nil {
		t.Fatal("expected cancellation")
	}
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Fatalf("cancelled output left behind: %v", err)
	}
}

func TestMergePBFMetadata(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "sources")
	left := `{"vector_layers":[{"id":"roads","fields":{"name":"String"},"minzoom":4,"maxzoom":10}],"custom_left":1}`
	right := `{"vector_layers":[{"id":"roads","fields":{"speed":"Number"},"minzoom":2,"maxzoom":12},{"id":"water","fields":{}}],"custom_right":2}`
	for i, value := range []string{left, right} {
		metadata := map[string]string{"name": "vectors", "format": "pbf", "json": value}
		fixtureDB(t, filepath.Join(src, string(rune('a'+i))+".mbtiles"), metadata, Tiles{ZoomLevel: int32(i), TileData: []byte("opaque vector bytes")})
	}
	dest := filepath.Join(dir, "dest.mbtiles")
	if err := Merge(src, dest); err != nil {
		t.Fatal(err)
	}
	got, err := readMergeMetadata(context.Background(), openResult(t, dest))
	if err != nil {
		t.Fatal(err)
	}
	object, err := parseMetadataJSON(got["json"])
	if err != nil {
		t.Fatal(err)
	}
	layers, err := parseVectorLayers(object["vector_layers"])
	if err != nil || len(layers) != 2 {
		t.Fatalf("layers=%v err=%v", layers, err)
	}
	var fields map[string]string
	if err := json.Unmarshal(layers[0]["fields"], &fields); err != nil {
		t.Fatal(err)
	}
	if fields["name"] != "String" || fields["speed"] != "Number" || string(layers[0]["minzoom"]) != "2" || string(layers[0]["maxzoom"]) != "12" || object["custom_left"] == nil || object["custom_right"] == nil {
		t.Fatalf("lost vector metadata: %s", got["json"])
	}
	_, err = mergeMetadataJSON(left, strings.ReplaceAll(left, `"String"`, `"Number"`))
	if err == nil {
		t.Fatal("conflicting field types must fail")
	}
}
