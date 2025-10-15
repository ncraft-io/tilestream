package grpc

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

type Grpc struct {
	svcName   string
	mode      string
	srcLayer  string
	instances []string
	//client    *Client
}

//func init() {
//	createReader := func(config tilestream.Config) tilestream.TileReader {
//		return New(config)
//	}
//	createWriter := func(config tilestream.Config) tilestream.TileWriter {
//		return New(config)
//	}
//
//	tilestream.RegisterReader("grpc", createReader)
//	tilestream.RegisterWriter("grpc", createWriter)
//}

func New(options *core.Options) *Grpc {
	/*
		// config looks like:
		config:
		  mode: etcd # etcd or direct, if direct, instances must be given
		  serviceName: maps.tilestream.v1.Tile # default 'maps.tilestream.v1.Tile'
		  layer: basemap #target layer, essential
		  instances:
		    - 52.82.35.145:17051
	*/
	//var (
	//	svcName, mode string
	//	instances     []string
	//	srcLayer      string
	//)
	//if v, ok := config["serviceName"]; !ok {
	//	svcName = tile_client.FullServiceName
	//} else if name, ok := v.(string); !ok {
	//	svcName = tile_client.FullServiceName
	//} else {
	//	svcName = name
	//}
	//
	//if v, ok := config["layer"]; !ok {
	//	panic("source layer not given")
	//} else if l, ok := v.(string); !ok {
	//	panic("could not parse source layer, string type wanted")
	//} else {
	//	srcLayer = l
	//}
	//
	//if v, ok := config["mode"]; !ok {
	//	mode = "etcd"
	//} else if m, ok := v.(string); !ok {
	//	mode = "etcd"
	//} else {
	//	mode = m
	//}
	//
	//var client *tile_client.Client
	//tracer, _ := tracing.New("maps.tilestream.v1.Tile")
	//logger := zap.Logger()
	//if mode == "direct" {
	//	if v, ok := config["instances"]; !ok {
	//		panic("instances not given for direct grpc connetcion")
	//	} else if ifaces, ok := v.([]interface{}); !ok {
	//		panic("failed to parse instances field")
	//	} else if len(ifaces) == 0 {
	//		panic("got zero instances")
	//	} else {
	//		instances = make([]string, 0, len(ifaces))
	//		for _, iface := range ifaces {
	//			instances = append(instances, iface.(string))
	//		}
	//	}
	//
	//	sdCfg := sd.Config{
	//		Mode: "direct",
	//		Retry: &retry.Config{
	//			Enable:  true,
	//			Timeout: 1000,
	//			Max:     3,
	//		},
	//		Direct: map[string]*direct.Config{
	//			svcName: &direct.Config{
	//				Urls: instances,
	//				Name: svcName,
	//			},
	//		},
	//	}
	//	clientCfg := &wzclient.Config{sdCfg}
	//	tileSd := direct.New(sdCfg.Direct)
	//	client = tile_client.New(clientCfg, tileSd.Instancer(svcName), tracer, logger)
	//
	//} else {
	//	sdCfg := sd.NewConfig("sd")
	//	sdClient := sd.New(sdCfg, logger)
	//	client = tile_client.New(wzclient.NewConfig("sd"), sdClient.Instancer(svcName), tracer, logger)
	//}
	//
	return &Grpc{
		//svcName:   svcName,
		//srcLayer:  srcLayer,
		//mode:      mode,
		//instances: instances,
		//client:    client,
	}
}

//func (g *Grpc) Tile(ctx context.Context, x, y, level int32) ([]byte, *tilestream.Options, error) {
//	req := &maps_tilestream_v1.GetTileRequest{
//		Layer: g.srcLayer,
//		Level: level,
//		X:     x,
//		Y:     y,
//	}
//	resp, err := g.client.GetTile(ctx, req)
//	if err != nil {
//		return nil, nil, err
//	}
//	return resp.Content, nil, nil
//}
//
//func (g Grpc) Info(ctx context.Context) (*tilestream.TileInfo, error) {
//	return nil, nil
//}
//
//func (g *Grpc) StartWriting(ctx context.Context) error {
//	return nil
//}
//
//func (g *Grpc) StopWriting(ctx context.Context) error {
//	return nil
//}
//
//func (g *Grpc) WriteTile(ctx context.Context, x, y, level int32, tile []byte) error {
//	req := &maps_tilestream_v1.UpdateTileRequest{
//		Layer: g.srcLayer,
//		Tile: &maps_tilestream_tile.Tile{
//			X:       x,
//			Y:       y,
//			Level:   level,
//			Content: tile,
//		},
//	}
//	_, err := g.client.UpdateTile(ctx, req)
//	return err
//}
//
//func (g *Grpc) WriteInfo(ctx context.Context, info tilestream.TileInfo) error {
//	return nil
//}
