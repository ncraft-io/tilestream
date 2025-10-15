package tilestream

import (
	"testing"
)

func TestZXYConv(t *testing.T) {
	lon := 120.888666
	lat := 30.666888
	z := int32(10)

	x := GetX(lon, z)
	y := GetY(lat, z)

	left := GetLon(float64(x), z)
	right := GetLon(float64(x+1), z)
	top := GetLat(float64(y), z)
	bottom := GetLat(float64(y+1), z)

	t.Logf("x %d, y %d", x, y)
	t.Logf("left %f, top %f", left, top)
	t.Logf("right %f, bottom %f", right, bottom)

	if lon > right || lon < left || lat > top || lat < bottom {
		t.Error("lon lat not in box")
	}
}

func TestZigZag(t *testing.T) {
	a := int32(100)
	b := int32(-1000)
	if a != zigzagDecode(zigzagEncode(a)) || b != zigzagDecode(zigzagEncode(b)) {
		t.Error("not equal")
	}
}

//func TestVectorTile(t *testing.T) {
//	var json = `{
//		"type":"FeatureCollection",
//		"features":[
//			{
//				"type": "Feature",
//				"geometry": {
//					"type": "MultiPoint",
//					"coordinates": [
//						[
//							120.249825,
//							32.443943
//						],
//						[
//							120.250825,
//							32.443823
//						]
//					]
//				},
//				"properties": {
//					"name": "大伦桥东工业园",
//					"featcode": 2010107,
//					"vt_layer": "poi_name_town"
//				},
//				"id": 33781411
//			},
//			{
//				"type": "Feature",
//				"geometry": {
//					"type": "Polygon",
//					"coordinates": [
//						[
//							[
//								120.228882,
//								32.551444
//							],
//							[
//								120.591431,
//								32.551444
//							],
//							[
//								120.591431,
//								32.245329
//							],
//							[
//								120.228882,
//								32.245329
//							],
//							[
//								120.228882,
//								32.551444
//							]
//						]
//					]
//				},
//				"properties": {
//					"cntry_name": "ocean",
//					"vt_layer": "world_base"
//				},
//				"id": 3752
//			}
//		]
//	}`
//
//	bs := []byte(json)
//	buf := &geom.GeoBuf{}
//	err := jsoniter.ConfigFastest.Unmarshal(bs, buf)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	id := tile.TileId{
//		Level: 10,
//		X:     854,
//		Y:     414,
//	}
//	tile := &Tile{}
//	tile.EncodeFrom(buf, id)
//
//	newBuf, err := tile.Decode(id)
//	if err != nil {
//		t.Error(err)
//	}
//
//	s, err := jsoniter.ConfigFastest.MarshalToString(newBuf)
//	if err != nil {
//		t.Error(err)
//	}
//
//	t.Log(s)
//}
