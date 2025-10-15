package tilestream

import "github.com/mojo-lang/core/go/pkg/mojo/core"

type Config struct {
	Name    string       `json:"name"`
	Type    string       `json:"type"`
	Options core.Options `json:"options"`
}
