package http

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTileStream_Tile(t *testing.T) {
	ts := New(map[string]interface{}{"tileUrl": "https://webst02.is.autonavi.com/appmaptile?style=6&x={x}&y={y}&z={level}"})
	content, _, err := ts.Tile(context.Background(), 27439, 13393, 15)
	assert.NoError(t, err)
	assert.NotEmpty(t, content)
	//assert.Equal(t, 48657, len(content))
}
