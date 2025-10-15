package tilestream

import (
	"context"
	"errors"
)

type DummyTileWriter struct {
}

var err = errors.New("dummy TileWriter DO NOT support write operations")

func (d DummyTileWriter) StartWriting(ctx context.Context) error {
	return err
}

func (d DummyTileWriter) StopWriting(ctx context.Context) error {
	return err
}

func (d DummyTileWriter) WriteTile(ctx context.Context, x, y, z int32, tile []byte) error {
	return err
}

func (d DummyTileWriter) WriteInfo(ctx context.Context, info *TileInfo) error {
	return err
}
