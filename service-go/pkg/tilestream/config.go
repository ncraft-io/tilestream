package tilestream

import "github.com/mojo-lang/mojo/go/pkg/mojo/core"

type Config struct {
	DefaultDbUri string       `json:"defaultDbUri,omitempty"`
	Layers       []*LyrConfig `json:"layers,omitempty"`
}

type LyrConfig struct {
	Name    string       `json:"name,omitempty"`
	Type    string       `json:"type,omitempty"`
	Options core.Options `json:"options,omitempty"`
}
