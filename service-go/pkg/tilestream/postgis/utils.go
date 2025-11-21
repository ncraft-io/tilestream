package postgis

import (
	"errors"
	"fmt"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/tegola"
	"github.com/go-spatial/tegola/atlas"
	"github.com/go-spatial/tegola/dict"
	"github.com/go-spatial/tegola/provider"
	"html"
	"strings"
)

func Providers(providers []dict.Dicter) (map[string]provider.Tiler, error) {
	// holder for registered providers
	registeredProviders := map[string]provider.Tiler{}

	// iterate providers
	for _, p := range providers {
		// lookup our proivder name
		pname, err := p.String("name", nil)
		if err != nil {
			switch err.(type) {
			case dict.ErrKeyRequired:
				return registeredProviders, nil
			case dict.ErrKeyType:
				return registeredProviders, nil
			default:
				return registeredProviders, err
			}
		}

		// check if a proivder with this name is alrady registered
		_, ok := registeredProviders[pname]
		if ok {
			return registeredProviders, nil
		}

		// lookup our provider type
		ptype, err := p.String("type", nil)
		if err != nil {
			switch err.(type) {
			case dict.ErrKeyRequired:
				return registeredProviders, nil
			case dict.ErrKeyType:
				return registeredProviders, nil
			default:
				return registeredProviders, err
			}
		}

		// register the provider
		prov, err := provider.For(ptype, p, nil)
		if err != nil {
			return registeredProviders, err
		}

		// add the provider to our map of registered providers
		registeredProviders[pname] = prov.Std
	}

	return registeredProviders, nil
}

func Maps(a *atlas.Atlas, maps []provider.Map, providers map[string]provider.Tiler) error {

	// iterate our maps
	for _, m := range maps {
		newMap := atlas.NewWebMercatorMap(string(m.Name))
		newMap.Attribution = html.EscapeString(string(m.Attribution))

		// convert from env package
		centerArr := [3]float64{}
		for i, v := range m.Center {
			centerArr[i] = float64(v)
		}

		newMap.Center = centerArr

		if len(m.Bounds) == 4 {
			newMap.Bounds = geom.NewExtent(
				[2]float64{float64(m.Bounds[0]), float64(m.Bounds[1])},
				[2]float64{float64(m.Bounds[2]), float64(m.Bounds[3])},
			)
		}

		if m.TileBuffer == nil {
			newMap.TileBuffer = tegola.DefaultTileBuffer
		} else {
			newMap.TileBuffer = uint64(*m.TileBuffer)
		}

		// iterate our layers
		for _, l := range m.Layers {
			// split our provider name (provider.layer) into [provider,layer]
			providerLayer := strings.Split(string(l.ProviderLayer), ".")

			// we're expecting two params in the provider layer definition
			if len(providerLayer) != 2 {
				return fmt.Errorf("layer invalid")
			}

			// lookup our proivder
			provider, ok := providers[providerLayer[0]]
			if !ok {
				return errors.New(`Provider Not Found ` + providerLayer[0])
			}

			// read the provider's layer names
			layerInfos, err := provider.Layers()
			if err != nil {
				return err
			}

			// confirm our providerLayer name is registered
			var found bool
			var layerGeomType tegola.Geometry
			for i := range layerInfos {
				if layerInfos[i].Name() == providerLayer[1] {
					found = true

					// read the layerGeomType
					layerGeomType = layerInfos[i].GeomType()
					break
				}
			}
			if !found {
				return fmt.Errorf("Provider layer not registered")
			}

			var defaultTags map[string]interface{}
			if l.DefaultTags != nil {
				defaultTags = l.DefaultTags
			}

			var minZoom uint
			if l.MinZoom != nil {
				minZoom = uint(*l.MinZoom)
			}

			var maxZoom uint
			if l.MaxZoom != nil {
				maxZoom = uint(*l.MaxZoom)
			}

			// add our layer to our layers slice
			newMap.Layers = append(newMap.Layers, atlas.Layer{
				Name:              string(l.Name),
				ProviderLayerName: providerLayer[1],
				MinZoom:           minZoom,
				MaxZoom:           maxZoom,
				Provider:          provider,
				DefaultTags:       defaultTags,
				GeomType:          layerGeomType,
				DontSimplify:      bool(l.DontSimplify),
				DontClip:          bool(l.DontClip),
				DontClean:         bool(l.DontClean),
			})
		}

		a.AddMap(newMap)
	}

	return nil
}
