type TileInfo {
    id: String @1
	name: String  @2
	type: String @3
	description: String @4
	format: String @5
	version: String @6
	attribution: String @7
    scheme: String @8 //< tms, xyz

	min_zoom: Int32 @10
	max_zoom: Int32 @11

    bounds: geom.BoundingBox @14
	center: geom.LngLat @15
}
