package tilestream

import (
	"crypto/md5"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream/providers/postgis"
)

func (x *Layer) SetPostgisFilter(filter string) {
	if x != nil && x.Config.Object != nil && x.Type == "postgis" {
		if cfg, err := postgis.NewConfigFrom(x.Config.Object); err != nil {

		} else {
			cfg.SetFilter(filter)
			x.Config.Object = cfg.ToObject()
		}
	}
}

func (x *Layer) VTileHash() string {
	if x != nil {
		layer := &Layer{
			Name:   x.Name,
			Type:   x.Type,
			Config: x.Config,
		}
		bytes, _ := jsoniter.Marshal(layer)
		return fmt.Sprintf("%x", md5.Sum(bytes))
	}

	return ""
}
