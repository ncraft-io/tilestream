package postgis

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/kvstore"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	ts "github.com/ncraft-io/tilestream/go/pkg/tilestream"
	"sync"
	"testing"
	"time"
)

var json = `
debug: true
mapName: emg_18q1_src
minZoom: 3
maxZoom: 16
bounds: [73.0,17.0,135.0,54.0]
atlas:
  providers:
    name: emg_18q1_src
    type: postgis        # the type of data provider. currently only supports postgis
    host: 172.2.0.21      # postgis database host
    port: 5434             # postgis database port
    database: emg_19q2_src       # postgis database name
    user: postgres         # postgis database user
    password: wayzpg@1234mid           # postgis database password
    srid: 4326             # The default srid for this provider. If not provided it will be WebMercator (3857)
    layers:
      - name: boundaries_china_province
        geometry_fieldname: wkb_geometry
        id_fieldname: ogc_fid
        sql: SELECT ST_AsBinary(wkb_geometry) AS wkb_geometry, namec, adapri, featcode, ogc_fid FROM admin WHERE wkb_geometry && !BBOX! and featcode='1010202'
      - name: boundaries_china_city
        geometry_fieldname: wkb_geometry
        id_fieldname: ogc_fid
        sql: SELECT ST_AsBinary(wkb_geometry) AS wkb_geometry, namec, adaid, featcode, ogc_fid FROM admin WHERE wkb_geometry && !BBOX! and featcode='1010203'
      - name: boundaries_china_district
        geometry_fieldname: wkb_geometry
        id_fieldname: ogc_fid
        sql: SELECT ST_AsBinary(wkb_geometry) AS wkb_geometry, namec, adaid, featcode, ogc_fid FROM admin WHERE wkb_geometry && !BBOX! and featcode='1010204'
  maps:
    name: emg_18q1_src
    layers:
      - provider_layer: emg_18q1_src.boundaries_china_province
        min_zoom: 3
        max_zoom: 16
      - provider_layer: emg_18q1_src.boundaries_china_city
        min_zoom: 9
        max_zoom: 16
      - provider_layer: emg_18q1_src.boundaries_china_district
        min_zoom: 12
        max_zoom: 16
`
var cWriteData = make(chan interface{}, 0)

type TileData struct {
	id  *ts.TileId
	val []byte
}

func TestNew(t *testing.T) {
	config := make(core.Options)
	_ = yaml.Unmarshal([]byte(json), &config)

	postgisPlus, err := New(config)
	if err != nil {
		panic(fmt.Errorf("new post gis plug failed"))
	}

	store := kvstore.NewStore("bblot", map[string]interface{}{
		"bucketName": "maps.boundaries.china",
		"path":       "boundaries-8.db",
	})

	if store == nil {
		panic(fmt.Errorf("new store failed"))
	}

	defer func() {
		if store != nil {
			store.Close()
		}
	}()

	var wgdb sync.WaitGroup
	wgdb.Add(1)
	go func() {
		for data := range cWriteData {
			dat := data.([]*TileData)
			for i := 0; i < len(dat); i++ {
				store.Put(context.Background(), []byte(dat[i].id.QuadKey()), dat[i].val)
			}
		}
		wgdb.Done()
	}()
	seedTiles(postgisPlus)
	wgdb.Wait()
	close(cWriteData)
}

func seedTiles(postgisPlus *Postgis) {
	for level := postgisPlus.Config.MinZoom; level < postgisPlus.Config.MaxZoom; level++ {
		minId := ts.GetTileId(postgisPlus.Config.Bounds.MinX(), postgisPlus.Config.Bounds.MaxY(), int32(level))
		maxId := ts.GetTileId(postgisPlus.Config.Bounds.MaxX(), postgisPlus.Config.Bounds.MinY(), int32(level))
		for x := minId.X; x < maxId.X; x++ {
			routineCount := 8
			var wg sync.WaitGroup
			wg.Add(routineCount)
			for i := 0; i < routineCount; i++ {
				go func(minY, maxY int32, index int) {
					var tileId ts.TileId
					tileAr := make([]*TileData, 0)
					for y := minY; y < maxY; y++ {
						y1 := int(y-minY) % routineCount
						if y1 != index {
							continue
						}
						tileId.X = x
						tileId.Y = y
						tileId.Level = int32(level)
						t := time.Now()
						if tileBytes, _, err := postgisPlus.Tile(context.Background(), tileId.X, tileId.Y, tileId.Level); err == nil {
							tileAr = append(tileAr, &TileData{&tileId, tileBytes})
							logs.Info(fmt.Sprintf("[INFO] cache seed tile %d/%d/%d takes %dms", tileId.Level, tileId.X, tileId.Y, time.Now().Sub(t).Nanoseconds()/1000000))
						} else {
							logs.Info(fmt.Sprintf("[ERROR] cache seed tile %d/%d/%d failed %v", tileId.Level, tileId.X, tileId.Y, err))
						}
					}
					if len(tileAr) > 0 {
						cWriteData <- tileAr
					}
					wg.Done()

				}(minId.Y, maxId.Y, i)
			}
			wg.Wait()
		}
	}
	logs.Infow("all done")
}
