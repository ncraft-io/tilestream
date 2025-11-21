package tilestream

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

func NewLayerConfigFrom(json string) (*LayerConfig, error) {
	cfg := &LayerConfig{}
	cfg.Init()
	if err := jsoniter.UnmarshalFromString(json, cfg.Object); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (x *LayerConfig) ToOptions() core.Options {
	if x != nil {
		options := make(core.Options)

		data, _ := jsoniter.Marshal(x.Object)
		_ = jsoniter.Unmarshal(data, &options)
		return options
	}
	return nil
}
