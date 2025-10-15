package handlers

import (
	"context"
	"github.com/ncraft-io/tilestream/service-go/pkg/model"
	"github.com/segmentio/ksuid"
	"sync"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"

	ts "github.com/ncraft-io/tilestream/service-go/pkg/tilestream"

	_ "github.com/ncraft-io/tilestream/service-go/pkg/tilestream/postgis"

	// this service api
	pb "github.com/ncraft-io/tilestream/go/pkg/tilestream/v1"
)

var (
	_ = tilestream.Tile{}
	_ = core.Null{}
	_ = tilestream.TileInfo{}
	_ = tilestream.Layer{}
)

type tilestreamServer struct {
	pb.UnimplementedTilestreamServer

	lock        *sync.Mutex
	TileStreams map[string]ts.TileStream
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.TilestreamServer {
	return tilestreamServer{
		lock:        &sync.Mutex{},
		TileStreams: make(map[string]ts.TileStream),
	}
}

// CreateTile implements Interface.
func (s tilestreamServer) CreateTile(ctx context.Context, in *pb.CreateTileRequest) (*tilestream.Tile, error) {
	resp := &tilestream.Tile{
		// X:
		// Y:
		// Level:
		// Format:
		// Content:
	}
	return resp, nil
}

// BatchCreateTiles implements Interface.
func (s tilestreamServer) BatchCreateTiles(ctx context.Context, in *pb.BatchCreateTilesRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// CreateTiles implements Interface.
func (s tilestreamServer) CreateTiles(ctx context.Context, in *pb.CreateTilesRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// GetTile implements Interface.
func (s tilestreamServer) GetTile(ctx context.Context, in *pb.GetTileRequest) (*tilestream.Tile, error) {
	var stream ts.TileStream
	{
		s.lock.Lock()
		defer s.lock.Unlock()

		if strm, ok := s.TileStreams[in.Layer]; ok {
			stream = strm
		} else {
			stream = ts.NewTileStream(in.Layer)
			s.TileStreams[in.Layer] = stream
		}
	}

	if stream == nil {
		return nil, core.NewInvalidArgumentError("The layer %s is exist", in.GetLayer())
	}

	if stream != nil {
		tile, options, err := stream.Tile(ctx, in.GetX(), in.GetY(), in.GetLevel())
		if err != nil {
			return nil, err
		}
		t := &tilestream.Tile{
			X:       in.GetX(),
			Y:       in.GetY(),
			Level:   in.GetLevel(),
			Content: tile,
		}

		format := options.GetString("Format")
		if len(format) > 0 {
			t.Format = format
		} else {

		}

		encoding := options.GetString("Content-Encoding")
		if len(encoding) > 0 {
			t.Encoding = encoding
		}

		return t, nil
	}

	return nil, core.NewNotFoundError("failed to found the tile %s(%d,%d,%d)", in.GetLevel(), in.GetX(), in.GetY(), in.GetLevel())
}

// GetTileInfo implements Interface.
func (s tilestreamServer) GetTileInfo(ctx context.Context, in *pb.GetTileInfoRequest) (*tilestream.TileInfo, error) {
	resp := &tilestream.TileInfo{
		// Id:
		// Name:
		// Description:
		// Version:
		// Legend:
		// Scheme:
		// MinZoom:
		// MaxZoom:
		// Bounds:
		// Center:
	}
	return resp, nil
}

// UpdateTile implements Interface.
func (s tilestreamServer) UpdateTile(ctx context.Context, in *pb.UpdateTileRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// UpdateTileInfo implements Interface.
func (s tilestreamServer) UpdateTileInfo(ctx context.Context, in *pb.UpdateTileInfoRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// CreateLayer implements Interface.
func (s tilestreamServer) CreateLayer(ctx context.Context, in *pb.CreateLayerRequest) (*tilestream.Layer, error) {
	if in.Layer == nil {
		return nil, core.NewInvalidArgumentError("the layer is empty")
	}
	if len(in.Layer.Name) == 0 {
		return nil, core.NewInvalidArgumentError("the layer's name is empty")
	}
	if len(in.Layer.Type) == 0 {
		return nil, core.NewInvalidArgumentError("the layer's type is empty")
	}
	if in.Layer.Config == nil {
		return nil, core.NewInvalidArgumentError("the layer's config is empty")
	}

	options := in.Layer.Config.ToOptions()
	if options == nil {
		return nil, core.NewInvalidArgumentError("the layer's config is invalid")
	}
	if tile := ts.LoadTileStream(in.Layer.Type, options); tile == nil {
		return nil, core.NewInvalidArgumentError("the layer's config is invalid")
	}

	if len(in.Layer.Id) == 0 {
		in.Layer.Id = ksuid.New().String()
	}
	if in.Layer.CreateTime == nil {
		in.Layer.CreateTime = core.Now()
	}
	if in.Layer.UpdateTime == nil {
		in.Layer.UpdateTime = core.Now()
	}

	if _, err := model.GetLayer().Create(ctx, in.Layer); err != nil {
		return nil, core.NewInternalError("failed to save the layer to database, error: %s", err.Error())
	}

	resp := &tilestream.Layer{
		Id: in.Layer.Id,
		// Name:
		// Type:
		// Templated:
		// Config:
		// Description:
		// CreateTime:
		// UpdateTime:
	}
	return resp, nil
}

// UpdateLayer implements Interface.
func (s tilestreamServer) UpdateLayer(ctx context.Context, in *pb.UpdateLayerRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// DeleteLayer implements Interface.
func (s tilestreamServer) DeleteLayer(ctx context.Context, in *pb.DeleteLayerRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// GetLayer implements Interface.
func (s tilestreamServer) GetLayer(ctx context.Context, in *pb.GetLayerRequest) (*tilestream.Layer, error) {
	resp := &tilestream.Layer{
		// Id:
		// Name:
		// Type:
		// Templated:
		// Config:
		// Description:
		// CreateTime:
		// UpdateTime:
	}
	return resp, nil
}

// ListLayers implements Interface.
func (s tilestreamServer) ListLayers(ctx context.Context, in *pb.ListLayersRequest) (*pb.ListLayersResponse, error) {
	resp := &pb.ListLayersResponse{
		// Layers:
	}
	return resp, nil
}
