package main

import (
	"database/sql"
	"io"
	"path/filepath"
	"strings"
	"testing"
)

func TestMergeRequiredFlags(t *testing.T) {
	for _, args := range [][]string{{"mbtiles", "merge"}, {"mbtiles", "merge", "--src", "input.mbtiles"}, {"mbtiles", "merge", "--dest", "output.mbtiles"}} {
		app := newApp()
		app.Writer, app.ErrWriter = io.Discard, io.Discard
		if err := app.Run(args); err == nil || !strings.Contains(err.Error(), "Required flag") {
			t.Fatalf("args=%v err=%v", args, err)
		}
	}
}

func TestMergeCommandAliases(t *testing.T) {
	dir := t.TempDir()
	src, dest := filepath.Join(dir, "input.mbtiles"), filepath.Join(dir, "output.mbtiles")
	database, err := sql.Open("sqlite3", src)
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	for _, q := range []string{
		`CREATE TABLE metadata(name TEXT,value TEXT)`,
		`INSERT INTO metadata VALUES ('name','cli'),('format','png')`,
		`CREATE TABLE tiles(zoom_level INTEGER,tile_column INTEGER,tile_row INTEGER,tile_data BLOB)`,
		`INSERT INTO tiles VALUES (5,25,18,X'616263')`,
	} {
		if _, err := database.Exec(q); err != nil {
			t.Fatal(err)
		}
	}
	if err := newApp().Run([]string{"mbtiles", "m", "-s", src, "-d", dest}); err != nil {
		t.Fatal(err)
	}
	result, err := sql.Open("sqlite3", dest)
	if err != nil {
		t.Fatal(err)
	}
	defer result.Close()
	var tile []byte
	if err := result.QueryRow(`SELECT tile_data FROM tiles WHERE zoom_level=5 AND tile_column=25 AND tile_row=18`).Scan(&tile); err != nil {
		t.Fatal(err)
	}
	if string(tile) != "abc" {
		t.Fatalf("tile=%q", tile)
	}
}
