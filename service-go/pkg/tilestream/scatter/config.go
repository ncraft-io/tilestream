package scatter

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

type Config struct {
	// the layer name
	Name string `json:"name"`

	KvType string `json:"kvType"`
}

func NewConfig(options core.Options) *Config {
	conf := &Config{}
	if v, ok := options["name"]; ok {
		if conf.Name, ok = v.(string); ok {
		}
	}

	return conf
}
