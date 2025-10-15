package tilestream

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

func (x *LayerConfig) ToOptions() core.Options {
	if x != nil {
		options := make(core.Options)

		data, _ := jsoniter.Marshal(x)
		_ = jsoniter.Unmarshal(data, options)
		return options
	}
	return nil
}
