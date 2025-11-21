package tilestream

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
)

func EncodeLayer(name string, feats []*geom.Feature, tid *TileId, extent uint32) *VectorTile_Layer {
	keyset := make(map[string]int)
	valset := make(map[string]*hashVal)
	tfeats := make([]*VectorTile_Feature, 0, len(feats))
	for _, f := range feats {
		tfeat := EncodeFeature(f, tid, float64(extent), keyset, valset)
		if tfeat != nil {
			tfeats = append(tfeats, tfeat)
		}
	}

	keys := make([]string, len(keyset))
	vals := make([]*VectorTile_Value, len(valset))
	for k, i := range keyset {
		keys[i] = k
	}
	for _, hv := range valset {
		vals[hv.seq] = hv.val
	}

	return &VectorTile_Layer{
		Name:     name,
		Version:  DefaultVersion,
		Keys:     keys,
		Values:   vals,
		Features: tfeats,
		Extent:   extent,
	}
}

func (x *VectorTile_Layer) Decode(tid *TileId) []*geom.Feature {
	if x != nil {
		feats := make([]*geom.Feature, 0, len(x.Features))
		extent := DefaultExtent
		if x.Extent != 0 {
			extent = x.Extent
		}

		for _, feat := range x.Features {
			f := feat.Decode(tid, float64(extent), x.Keys, x.Values)
			f.Properties[DefaultLayerTag] = core.NewStringValue(x.Name)
			feats = append(feats, f)
		}

		return feats
	}
	return nil
}
