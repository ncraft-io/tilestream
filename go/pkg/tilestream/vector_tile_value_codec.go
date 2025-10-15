package tilestream

import (
	"encoding/binary"
	"fmt"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"math"
)

const (
	stringTag = 10 // 1 010 = 1<<3|2, length-delimited
	floatTag  = 21 // 10 101 = 2<<3|5, 32-bit
	doubleTag = 25 // 11 001 = 3<<3|1, 64-bit
	intTag    = 32 // 100 000 = 4<<3|0, varint
	uintTag   = 40 // 101 000 = 5<<3|0, varint
	sintTag   = 48 // 110 000 = 6<<3|0, varint
	boolTag   = 56 // 111 000 = 7<<3|0, varint

	valHeader = 34 // 100 010 = 4<<3|2, length-delimited
)

type hashVal struct {
	hash string
	val  *VectorTile_Value
	seq  int
}

func encodeValue(in *core.Value) *hashVal {
	switch vt := in.Val.(type) {
	case *core.Value_StringVal:
		return &hashVal{
			hash: hashValue(vt.StringVal),
			val: &VectorTile_Value{
				StringValue: vt.StringVal,
			},
		}
	case *core.Value_BoolVal:
		return &hashVal{
			hash: hashValue(vt.BoolVal),
			val: &VectorTile_Value{
				BoolValue: vt.BoolVal,
			},
		}
	case *core.Value_DoubleVal:
		return &hashVal{
			hash: hashValue(vt.DoubleVal),
			val: &VectorTile_Value{
				DoubleValue: vt.DoubleVal,
			},
		}
	case *core.Value_NegativeVal:
		x := -int64(vt.NegativeVal)
		return &hashVal{
			val: &VectorTile_Value{
				SintValue: x,
			},
			hash: hashValue(x),
		}
	case *core.Value_PositiveVal:
		x := int64(vt.PositiveVal)
		return &hashVal{
			val: &VectorTile_Value{
				IntValue: x,
			},
			hash: hashValue(x),
		}
	}

	return nil
}

func hashValue(v interface{}) string {
	var vpb []byte
	switch vt := v.(type) {
	case string:
		vpb = append([]byte{stringTag}, appendString(nil, vt)...)
	case uint64:
		vpb = append([]byte{uintTag}, appendUvarint(nil, vt)...)
	case float32:
		vpb = appendFloat32([]byte{floatTag}, vt)
	case float64:
		vpb = appendFloat64([]byte{doubleTag}, vt)
	case int64:
		vpb = appendVarint([]byte{sintTag}, vt)
	case int:
		vpb = appendVarint([]byte{intTag}, int64(vt))
	case bool:
		if vt {
			vpb = []byte{boolTag, 1}
		} else {
			vpb = []byte{boolTag, 0}
		}
	case uint8:
		return hashValue(uint64(vt))
	case uint16:
		return hashValue(uint64(vt))
	case uint32:
		return hashValue(uint64(vt))
	case int8:
		return hashValue(int64(vt))
	case int16:
		return hashValue(int64(vt))
	case int32:
		return hashValue(int64(vt))
	case []byte:
		return hashValue(string(vt))
	default:
		return hashValue(fmt.Sprintf("%v", v))
	}

	var pb = []byte{valHeader}
	pb = appendUvarint(pb, uint64(len(vpb)))
	pb = append(pb, vpb...)
	return string(pb)
}

func appendFloat32(pb []byte, v float32) []byte {
	bf := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(bf, math.Float32bits(v))
	return append(pb, bf...)
}

func appendFloat64(pb []byte, v float64) []byte {
	bf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint64(bf, math.Float64bits(v))
	return append(pb, bf...)
}

func appendString(pb []byte, s string) []byte {
	pb = appendUvarint(pb, uint64(len(s)))
	return append(pb, s...)
}

func appendUvarint(pb []byte, n uint64) []byte {
	bf := make([]byte, 10)
	sz := binary.PutUvarint(bf, n)
	return append(pb, bf[:sz]...)
}

func appendVarint(pb []byte, n int64) []byte {
	bf := make([]byte, 10)
	sz := binary.PutVarint(bf, n)
	return append(pb, bf[:sz]...)
}

//func valueConv(in *core.Value) *Tile_Value {
//	v := &Tile_Value{}
//	switch vt := in.Val.(type) {
//	case *geom.Value_StringVal:
//		v.StringValue = &vt.StringVal
//	case *geom.Value_PosIntVal:
//		tmp := new(int64)
//		*tmp = int64(vt.PosIntVal)
//		v.IntValue = tmp
//	case *geom.Value_NegIntVal:
//		tmp := new(int64)
//		*tmp = -int64(vt.NegIntVal)
//		v.IntValue = tmp
//	case *geom.Value_DoubleVal:
//		v.DoubleValue = &vt.DoubleVal
//	case *geom.Value_BoolVal:
//		v.BoolValue = &vt.BoolVal
//	default:
//		return nil
//	}
//	return v
//}
