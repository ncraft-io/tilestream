package mbtiles

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var schemaStatements = []string{
	`CREATE TABLE IF NOT EXISTS tiles (zoom_level INTEGER, tile_column INTEGER, tile_row INTEGER, tile_data BLOB)`,
	`CREATE TABLE IF NOT EXISTS metadata (name TEXT, value TEXT)`,
	`CREATE UNIQUE INDEX IF NOT EXISTS mbtiles_tile_coordinates ON tiles (zoom_level, tile_column, tile_row)`,
	`CREATE UNIQUE INDEX IF NOT EXISTS mbtiles_metadata_name ON metadata (name)`,
}

func canonicalFile(file string) (string, error) {
	absolute, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	resolved, err := filepath.EvalSymlinks(absolute)
	if err == nil {
		return resolved, nil
	}
	if !os.IsNotExist(err) {
		return "", err
	}
	return absolute, nil
}

func modelFile(config *Config, layer string) (string, error) {
	if config == nil || len(config.Paths) == 0 || layer == "" {
		return "", fmt.Errorf("missing MBTiles path or layer")
	}
	for _, dir := range config.Paths {
		file := filepath.Join(dir, layer+".mbtiles")
		stat, err := os.Stat(file)
		if err == nil && stat.Mode().IsRegular() {
			return canonicalFile(file)
		}
		if err != nil && !os.IsNotExist(err) {
			return "", err
		}
	}
	return canonicalFile(filepath.Join(config.Paths[0], layer+".mbtiles"))
}

func sqliteDSN(file, mode string) string {
	u := url.URL{Scheme: "file", Path: file}
	q := url.Values{"mode": {mode}, "_busy_timeout": {"5000"}}
	u.RawQuery = q.Encode()
	return u.String()
}

func openModelDB(file string) (*db.DB, error) {
	database, err := gorm.Open(sqlite.Open(sqliteDSN(file, "rwc")), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open MBTiles %q: %w", file, err)
	}
	return &db.DB{DB: database}, nil
}

func initializeModelSchema(database *gorm.DB) error {
	return database.Transaction(func(tx *gorm.DB) error {
		for _, statement := range schemaStatements {
			if err := tx.Exec(statement).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
