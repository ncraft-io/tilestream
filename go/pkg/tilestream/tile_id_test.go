package tilestream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTileId(t *testing.T) {
	id := GetTileId(121.458561, 31.193601, 16)
	assert.NotNil(t, id)
}
