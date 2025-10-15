package mbtiles

type Tiles struct {
	ZoomLevel  int32 `gorm:"index:tile_index"`
	TileColumn int32 `gorm:"index:tile_index"`
	TileRow    int32 `gorm:"index:tile_index"`
	TileData   []byte
}
