package tilestream

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/geom/go/pkg/mojo/geom"
)

// EncodeFeature ...
func EncodeFeature(feat *geom.Feature, tid *TileId, extent float64, keyset map[string]int, valset map[string]*hashVal) *VectorTile_Feature {
	tags := encodeTags(feat.Properties, keyset, valset)
	g, t := EncodeGeometry(feat.Geometry, tid, extent)
	if t == VectorTile_GEOM_TYPE_UNKNOWN {
		return nil
	}

	return &VectorTile_Feature{
		Id:       feat.Id.GetUint64Val(),
		Tags:     tags,
		Type:     t,
		Geometry: g,
	}
}

func (x *VectorTile_Feature) Decode(tid *TileId, extent float64, keys []string, vals []*VectorTile_Value) *geom.Feature {
	g := DecodeGeometry(x.GetGeometry(), x.GetType(), tid, extent)
	properties := decodeTags(x.GetTags(), keys, vals)
	feat := &geom.Feature{
		Properties: properties,
		Geometry:   g,
	}

	id := x.Id
	if id != 0 {
		feat.Id.SetInt(id)
	}

	return feat
}

func encodeTags(props map[string]*core.Value, keyset map[string]int, valset map[string]*hashVal) []uint32 {
	n := len(props) / 2
	tags := make([]uint32, 0, 2*n+1)
	//for u := 0; u < n; u++ {
	//	w := 2 * u
	//	k := keys[props[w]]
	//	if k == DefaultLayerTag {
	//		continue
	//	}
	//	v := vals[props[w+1]]
	//	hv := encodeValue(v)
	//	if hv == nil {
	//		continue
	//	}
	//
	//	var m, n int
	//	if seq, ok := keyset[k]; !ok {
	//		m = len(keyset)
	//		keyset[k] = m
	//	} else {
	//		m = seq
	//	}
	//
	//	if _hv, ok := valset[hv.hash]; !ok {
	//		n = len(valset)
	//		valset[hv.hash] = &hashVal{
	//			hash: hv.hash,
	//			val:  hv.val,
	//			seq:  n,
	//		}
	//	} else {
	//		n = _hv.seq
	//	}
	//
	//	tags = append(tags, uint32(m), uint32(n))
	//}
	return tags
}

func decodeTags(tags []uint32, allkeys []string, allvals []*VectorTile_Value) map[string]*core.Value {
	//nprops := len(tags) / 2
	//indices := make([]uint32, 0, nprops*2)
	//keys := make([]string, 0, nprops)
	//veys := make([]*core.Value, 0, nprops)
	//for i := 0; i < nprops; i++ {
	//	m := 2 * i
	//	n := m + 1
	//	idxk := tags[m]
	//	idxv := tags[n]
	//	keys = append(keys, allkeys[idxk])
	//	v := allvals[idxv]
	//	if s := v.StringValue; s != nil {
	//		veys = append(veys, core.NewStringValue(*s))
	//	} else if f := v.FloatValue; f != nil {
	//		veys = append(veys, core.NewFloat64Value(float64(*f)))
	//	} else if d := v.DoubleValue; d != nil {
	//		veys = append(veys, core.NewFloat64Value(*d))
	//	} else if _i := v.IntValue; _i != nil {
	//		veys = append(veys, core.NewIntValue(*_i))
	//	} else if u := v.UintValue; u != nil {
	//		veys = append(veys, core.NewIntValue(int64(*u)))
	//	} else if _s := v.SintValue; _s != nil {
	//		veys = append(veys, core.NewIntValue(*_s))
	//	} else if b := v.BoolValue; b != nil {
	//		veys = append(veys, core.NewBoolValue(*b))
	//	}
	//	indices = append(indices, uint32(i), uint32(i))
	//}
	//
	//return indices, keys, veys
	return nil
}
