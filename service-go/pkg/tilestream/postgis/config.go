package postgis

import (
	"fmt"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/tegola/config"
	"strings"
)

type Config struct {
	Debug     bool   `json:"debug"`
	MapName   string `json:"mapName"`
	LayerName string `json:"layerName"`

	MinZoom int         `json:"minZoom"`
	MaxZoom int         `json:"maxZoom"`
	Bounds  geom.Extent `json:"bounds"`

	Atlas config.Config `json:"atlas"`
}

func resetAtlasOptions(conf map[string]interface{}) (string, error) {

	confArr := make([]string, 0)

	if conf["providers"] != nil {
		confArr = append(confArr, "[[providers]]")
		atlas := conf["providers"].(map[string]interface{})
		var layers []interface{}
		for k, v := range atlas {
			if k == "layers" {
				layers = v.([]interface{})
			} else {
				if k != "srid" {
					confArr = append(confArr, fmt.Sprintf("%s=\"%v\"", k, v))
				} else {
					confArr = append(confArr, fmt.Sprintf("%s=%v", k, v))
				}
			}
		}

		if layers != nil {
			for _, layer := range layers {
				confArr = append(confArr, "  [[providers.layers]]")
				layerconf := layer.(map[string]interface{})
				for k, v := range layerconf {
					confArr = append(confArr, fmt.Sprintf("    %s=\"%v\"", k, v))
				}
			}
		}
	}

	if conf["maps"] != nil {
		confArr = append(confArr, "[[maps]]")
		maps := conf["maps"].(map[string]interface{})
		var layers []interface{}
		for k, v := range maps {
			if k == "layers" {
				layers = v.([]interface{})
			} else {
				confArr = append(confArr, fmt.Sprintf("%s=\"%v\"", k, v))
			}
		}
		if layers != nil {
			for _, layer := range layers {
				confArr = append(confArr, "  [[maps.layers]]")
				layerConf := layer.(map[string]interface{})
				for k, v := range layerConf {
					confArr = append(confArr, fmt.Sprintf("    %s=\"%v\"", k, v))
				}
			}
		}
	}
	confStr := strings.Join(confArr, "\n")
	return confStr, nil
}
