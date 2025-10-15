package tilestream

import (
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadTileStream(t *testing.T) {
	options := make(map[string]interface{})
	err := config.ScanFrom(options, "tileStreams")
	assert.NoError(t, err)
}
