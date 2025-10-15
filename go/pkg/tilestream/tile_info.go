package tilestream

func (x *TileInfo) Merge(info *TileInfo) *TileInfo {
	if x != nil {
		if info.MinZoom < x.MinZoom {
			x.MinZoom = info.MinZoom
		}
		if info.MaxZoom > x.MaxZoom {
			x.MaxZoom = info.MaxZoom
		}
		x.Bounds.Extend(info.Bounds)
	}
	return x
}
