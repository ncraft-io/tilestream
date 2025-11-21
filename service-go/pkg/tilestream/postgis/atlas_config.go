package postgis

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/tegola/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream/providers/postgis"
	"text/template"
)

type AtlasConfig struct {
	Debug     bool   `json:"debug"`
	MapName   string `json:"mapName"`
	LayerName string `json:"layerName"`

	MinZoom int         `json:"minZoom"`
	MaxZoom int         `json:"maxZoom"`
	Bounds  geom.Extent `json:"bounds"`

	Atlas    config.Config          `json:"-"`
	AtlasCfg map[string]interface{} `json:"atlas"` // should change to toml
}

func NewAtlasConfigFrom(json string) (*AtlasConfig, error) {
	cfg := &AtlasConfig{}

	if err := (jsoniter.Config{UseNumber: true}).Froze().UnmarshalFromString(json, cfg); err != nil {
		return nil, err
	}
	conf, err := toml.Marshal(cfg.AtlasCfg)
	if err != nil {
		return nil, err
	}
	logs.Debugw("the atlas config", "atlas", string(conf))
	cfg.Atlas, err = config.Parse(bytes.NewReader(conf), "")
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func NewAtlasConfig(cfg *postgis.Config) (*AtlasConfig, error) {
	tmp, _ := template.New("config").Parse(`{
      "debug": true,
      "mapName": "{{.Name}}",
      "minZoom": {{.MinZoom}},
      "maxZoom": {{.MaxZoom}},
      "bounds": [{{.Bounds.LeftBottom.Longitude}}, {{.Bounds.LeftBottom.Latitude}}, {{.Bounds.RightTop.Longitude}}, {{.Bounds.RightTop.Latitude}}],
      "atlas": {
        "providers": [{
          "name": "{{.Name}}",
          "type": "postgis",
          "uri": "{{.Provider.Uri}}",
          "srid": 4326,
          "layers":[{
              "name": "{{.Name}}",
              "geometry_fieldname": "{{.Provider.GeometryField}}",
              "id_fieldname": "{{.Provider.IdField}}",
              "sql": "{{.GetSql}}"
			}]
		}],
        "maps": [{
          "name": "{{.Name}}",
          "layers": [{
              "provider_layer": "{{.Name}}.{{.Name}}",
              "min_zoom": {{.MinZoom}},
              "max_zoom": {{.MaxZoom}},
			  "dont_clean": true
          }]
		}]
		}
	}`)
	buf := bytes.NewBuffer(nil)
	if err := tmp.Execute(buf, cfg); err != nil {
		return nil, err
	}
	return NewAtlasConfigFrom(buf.String())
}
