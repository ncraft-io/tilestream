package mbtiles

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

// Merge copies stored TMS coordinates verbatim. Directory inputs are recursive
// and sorted by path; later sources replace earlier tiles at the same coordinate.
func Merge(src, dest string) error {
	return MergeContext(context.Background(), src, dest)
}

func MergeContext(ctx context.Context, src, dest string) (err error) {
	if !IsMbtilesFile(dest) {
		return fmt.Errorf("destination must be an .mbtiles file: %q", dest)
	}
	dest, err = canonicalFile(dest)
	if err != nil {
		return err
	}
	files, err := mergeFiles(src, dest)
	if err != nil {
		return err
	}

	// Reserve a new destination so failure cleanup cannot remove a pre-existing file.
	created := false
	f, openErr := os.OpenFile(dest, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if openErr == nil {
		created = true
		if err = f.Close(); err != nil {
			os.Remove(dest)
			return err
		}
	} else if !os.IsExist(openErr) {
		return openErr
	}
	defer func() {
		if err != nil && created {
			_ = os.Remove(dest)
		}
	}()

	output, err := sql.Open("sqlite3", sqliteDSN(dest, "rw"))
	if err != nil {
		return err
	}
	defer output.Close()
	tx, err := output.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, statement := range schemaStatements {
		if _, err = tx.ExecContext(ctx, statement); err != nil {
			return fmt.Errorf("initialize destination: %w", err)
		}
	}
	metadata, err := readMergeMetadata(ctx, tx)
	if err != nil {
		return fmt.Errorf("read destination metadata: %w", err)
	}
	var existingTiles int
	if err = tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM tiles)`).Scan(&existingTiles); err != nil {
		return err
	}
	if existingTiles != 0 {
		if err = validateMergeMetadata(metadata); err != nil {
			return fmt.Errorf("destination: %w", err)
		}
	}
	// A partial bounds union would falsely exclude tiles from a source without bounds.
	completeBounds := existingTiles == 0 || metadata["bounds"] != ""
	insert, err := tx.PrepareContext(ctx, `INSERT INTO tiles (zoom_level,tile_column,tile_row,tile_data) VALUES (?,?,?,?) ON CONFLICT(zoom_level,tile_column,tile_row) DO UPDATE SET tile_data=excluded.tile_data`)
	if err != nil {
		return err
	}
	defer insert.Close()
	for _, file := range files {
		if err = mergeSource(ctx, file, insert, metadata, &completeBounds); err != nil {
			return fmt.Errorf("merge %q: %w", file, err)
		}
	}
	if !completeBounds {
		delete(metadata, "bounds")
	}
	var minZoom, maxZoom sql.NullInt64
	if err = tx.QueryRowContext(ctx, `SELECT MIN(zoom_level), MAX(zoom_level) FROM tiles`).Scan(&minZoom, &maxZoom); err != nil {
		return err
	}
	if minZoom.Valid {
		metadata["minzoom"] = strconv.FormatInt(minZoom.Int64, 10)
		metadata["maxzoom"] = strconv.FormatInt(maxZoom.Int64, 10)
	}
	keys := make([]string, 0, len(metadata))
	for name := range metadata {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	if _, err = tx.ExecContext(ctx, `DELETE FROM metadata`); err != nil {
		return err
	}
	for _, name := range keys {
		if _, err = tx.ExecContext(ctx, `INSERT INTO metadata(name,value) VALUES (?,?)`, name, metadata[name]); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func mergeFiles(src, dest string) ([]string, error) {
	stat, err := os.Stat(src)
	if err != nil {
		return nil, err
	}
	destStat, err := os.Stat(dest)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	var files []string
	add := func(file string, fromDirectory bool) error {
		if !IsMbtilesFile(file) {
			return fmt.Errorf("invalid source: %q", file)
		}
		absolute, err := canonicalFile(file)
		if err != nil {
			return err
		}
		sourceStat, err := os.Stat(absolute)
		if err != nil {
			return err
		}
		if !sourceStat.Mode().IsRegular() {
			return fmt.Errorf("source is not a regular file: %q", file)
		}
		if absolute == dest || (destStat != nil && os.SameFile(sourceStat, destStat)) {
			if fromDirectory {
				return nil
			}
			return fmt.Errorf("source and destination are the same file: %q", file)
		}
		files = append(files, absolute)
		return nil
	}
	if stat.IsDir() {
		err = filepath.WalkDir(src, func(file string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() || !IsMbtilesFile(file) {
				return nil
			}
			return add(file, true)
		})
	} else {
		err = add(src, false)
	}
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no source .mbtiles files in %q", src)
	}
	sort.Strings(files)
	return files, nil
}

type metadataQuerier interface {
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
}

func readMergeMetadata(ctx context.Context, database metadataQuerier) (map[string]string, error) {
	rows, err := database.QueryContext(ctx, `SELECT name,value FROM metadata`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	metadata := make(map[string]string)
	for rows.Next() {
		var name, value string
		if err := rows.Scan(&name, &value); err != nil {
			return nil, err
		}
		metadata[name] = value
	}
	return metadata, rows.Err()
}

func mergeSource(ctx context.Context, file string, insert *sql.Stmt, metadata map[string]string, completeBounds *bool) error {
	source, err := sql.Open("sqlite3", sqliteDSN(file, "ro"))
	if err != nil {
		return err
	}
	defer source.Close()
	snapshot, err := source.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	defer snapshot.Rollback()
	incoming, err := readMergeMetadata(ctx, snapshot)
	if err != nil {
		return err
	}
	if err := validateMergeMetadata(incoming); err != nil {
		return err
	}
	if incoming["bounds"] == "" {
		*completeBounds = false
	}
	if err := combineMergeMetadata(metadata, incoming); err != nil {
		return err
	}
	rows, err := snapshot.QueryContext(ctx, `SELECT zoom_level,tile_column,tile_row,tile_data FROM tiles`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var z, x, y int32
		var data []byte
		if err := rows.Scan(&z, &x, &y, &data); err != nil {
			return err
		}
		if _, err := tmsRow(x, y, z); err != nil {
			return err
		}
		if _, err := insert.ExecContext(ctx, z, x, y, data); err != nil {
			return err
		}
	}
	return rows.Err()
}
