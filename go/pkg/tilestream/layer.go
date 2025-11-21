package tilestream

import "github.com/ncraft-io/tilestream/go/pkg/tilestream/providers/postgis"

func (x *Layer) SetPostgisFilter(filter string) {
	if x != nil && len(filter) > 0 && x.Config.Object != nil && x.Type == "postgis" {
		if cfg, err := postgis.NewConfigFrom(x.Config.Object); err != nil {

		} else {
			cfg.SetFilter(filter)
			x.Config.Object = cfg.ToObject()
		}
	}
}
