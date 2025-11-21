package handlers

import (
	"context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"
	"github.com/ncraft-io/tilestream/service-go/pkg/model"
	"github.com/segmentio/ksuid"
	"net/url"

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
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.TilestreamServer {
	return tilestreamServer{}
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

func getHashKey(ctx context.Context) string {
	if values, ok := ctx.Value("http-request-query").(url.Values); ok {
		if vs, ok := values["hash_key"]; ok && len(vs) > 0 {
			return vs[0]
		}
	}
	return ""
}

// GetTile implements Interface.
func (s tilestreamServer) GetTile(ctx context.Context, in *pb.GetTileRequest) (*tilestream.Tile, error) {
	instance := GetTileStream()

	ccCtx := ctx
	if instance.Cache != nil {
		if hashKey := getHashKey(ctx); len(hashKey) > 0 {
			ccCtx = context.WithValue(ccCtx, "hash_key", hashKey)
		}
		ccCtx = context.WithValue(ccCtx, "layer", in.Layer)

		if tile, _, err := instance.Cache.Tile(ccCtx, in.GetX(), in.GetY(), in.GetLevel()); err == nil && len(tile) > 0 {
			return &tilestream.Tile{
				X:       in.GetX(),
				Y:       in.GetY(),
				Level:   in.GetLevel(),
				Content: tile,
				Format:  in.Format,
			}, nil
		}
	}

	st := instance.Get(in.Layer)
	if st == nil {
		return nil, core.NewInvalidArgumentError("The layer %s is exist", in.GetLayer())
	}
	tile, options, err := st.Tile(ctx, in.GetX(), in.GetY(), in.GetLevel())
	if err != nil {
		return nil, err
	}

	if instance.Cache != nil {
		_ = instance.Cache.WriteTile(ccCtx, in.GetX(), in.GetY(), in.GetLevel(), tile)
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

	//return nil, core.NewNotFoundError("failed to found the tile %s(%d,%d,%d)", in.GetLevel(), in.GetX(), in.GetY(), in.GetLevel())
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
	if in.Layer.Config.Object == nil {
		return nil, core.NewInvalidArgumentError("the layer's config is empty")
	}

	options := in.Layer.Config.ToOptions()
	if options == nil {
		return nil, core.NewInvalidArgumentError("the layer's config is invalid")
	}
	if tile := ts.LoadTileStream(in.Layer.Type, options); tile == nil {
		return nil, core.NewInvalidArgumentError("the layer's config is invalid")
	}

	if len(in.Layer.OriginalId) == 0 {
		if layer, _ := model.NewLayer().Get(ctx, in.Layer.Name); layer != nil {
			in.Layer.Id = layer.Id
			in.Layer.CreateTime = layer.CreateTime
		}
	}

	if len(in.Layer.Id) == 0 {
		in.Layer.Id = ksuid.New().String()
	}
	if in.Layer.CreateTime == nil {
		in.Layer.CreateTime = core.Now()
	}

	in.Layer.UpdateTime = core.Now()

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
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("the layer id is empty")
	}
	if in.Layer == nil {
		return nil, core.NewInvalidArgumentError("the layer is empty")
	}
	if len(in.Layer.Id) == 0 {
		in.Layer.Id = in.Id
	}

	var dbLayer *tilestream.Layer
	if l, err := model.NewLayer().Get(ctx, in.Layer.Id); err != nil {
		return nil, core.NewInvalidArgumentError("the layer id (%s) is not found", in.Layer.Id)
	} else {
		dbLayer = l
	}

	if in.Layer.Config.Object != nil {
		options := in.Layer.Config.ToOptions()
		if options == nil {
			return nil, core.NewInvalidArgumentError("the layer's config is invalid")
		}
		if tile := ts.LoadTileStream(in.Layer.Type, options); tile == nil {
			return nil, core.NewInvalidArgumentError("the layer's config is invalid")
		}
		in.Layer.UpdateTime = core.Now()
		if _, err := model.GetLayer().Update(ctx, in.Layer); err != nil {
			return nil, core.NewInternalError("failed to save the layer to database, error: %s", err.Error())
		}

		{
			instance := GetTileStream()
			instance.Delete(in.Layer.Id)
			if len(dbLayer.OriginalId) == 0 {
				instance.Delete(in.Layer.Name)
			}

		}
	}
	return &core.Null{}, nil
}

// DeleteLayer implements Interface.
func (s tilestreamServer) DeleteLayer(ctx context.Context, in *pb.DeleteLayerRequest) (*core.Null, error) {
	if len(in.Layer) == 0 {
		return nil, core.NewInvalidArgumentError("the layer id is empty")
	}

	{
		instance := GetTileStream()
		instance.Delete(in.Layer)
	}

	if _, err := model.NewLayer().Delete(ctx, in.Layer); err != nil {
		return nil, core.NewInvalidArgumentError("failed to delete the layer (%s)", in.Layer)
	}

	return &core.Null{}, nil
}

// GetLayer implements Interface.
func (s tilestreamServer) GetLayer(ctx context.Context, in *pb.GetLayerRequest) (*tilestream.Layer, error) {
	if len(in.Layer) == 0 {
		return nil, core.NewInvalidArgumentError("the layer id is empty")
	}

	if layer, err := model.NewLayer().Get(ctx, in.Layer); err != nil {
		return nil, core.NewInvalidArgumentError("the layer id (%s) is not found", in.Layer)
	} else {
		return layer, nil
	}
}

// ListLayers implements Interface.
func (s tilestreamServer) ListLayers(ctx context.Context, in *pb.ListLayersRequest) (*pb.ListLayersResponse, error) {
	resp := &pb.ListLayersResponse{
		// Layers:
	}
	return resp, nil
}

// BatchUpdateLayer implements Interface.
func (s tilestreamServer) BatchUpdateLayer(ctx context.Context, in *pb.BatchUpdateLayerRequest) (*core.Null, error) {
	if len(in.Layers) == 0 {
		return nil, core.NewInvalidArgumentError("the layers is empty")
	}

	var ids []string
	for i, layer := range in.Layers {
		if len(layer.Id) == 0 {
			return nil, core.NewInvalidArgumentError("the No.%d layer's id is empty, name is %s", i, layer.Name)
		}
		layer.UpdateTime = core.Now()
		ids = append(ids, layer.Id+":"+layer.Name)
	}

	{
		instance := GetTileStream()

		for _, layer := range in.Layers {
			instance.Delete(layer.Id)
			if len(layer.OriginalId) == 0 {
				instance.Delete(layer.Name)
			}
		}
	}

	for _, layer := range in.Layers {
		_, err := model.GetLayer().Update(ctx, layer)
		if err != nil {
			return nil, core.NewInternalError("failed to update layers (%v) err: %s", ids, err.Error())
		}
	}

	return &core.Null{}, nil
}

// BatchGetLayers implements Interface.
func (s tilestreamServer) BatchGetLayers(ctx context.Context, in *pb.BatchGetLayersRequest) (*pb.BatchGetLayersResponse, error) {
	if len(in.Layers) == 0 {
		return nil, core.NewInvalidArgumentError("the layers is empty")
	}

	layers, err := model.GetLayer().BatchGet(ctx, in.Layers)
	if err != nil {
		return nil, core.NewInternalError("failed to get layers (%v) err: %s", in.Layers, err.Error())
	}

	resp := &pb.BatchGetLayersResponse{
		Layers: layers,
	}
	return resp, nil
}

// BatchCreateLayer implements Interface.
func (s tilestreamServer) BatchCreateLayer(ctx context.Context, in *pb.BatchCreateLayerRequest) (*pb.BatchCreateLayerResponse, error) {
	if len(in.Layers) == 0 {
		return nil, core.NewInvalidArgumentError("the layers is empty")
	}
	for i, layer := range in.Layers {
		if layer == nil {
			return nil, core.NewInvalidArgumentError("the No.%d layer is empty", i)
		}
		if len(layer.Name) == 0 {
			return nil, core.NewInvalidArgumentError("the No.%d layer's name is empty", i)
		}
		if len(layer.Type) == 0 {
			return nil, core.NewInvalidArgumentError("the No.%d layer's type is empty", i)
		}
		if layer.Config.Object == nil {
			return nil, core.NewInvalidArgumentError("the No.%d layer's config is empty", i)
		}

		options := layer.Config.ToOptions()
		if options == nil {
			return nil, core.NewInvalidArgumentError("the No.%d layer's config is invalid", i)
		}
		if tile := ts.LoadTileStream(layer.Type, options); tile == nil {
			return nil, core.NewInvalidArgumentError("the No.%d layer's config is invalid", i)
		}

		if len(layer.OriginalId) == 0 {
			if l, _ := model.NewLayer().Get(ctx, layer.Name); l != nil {
				layer.Id = l.Id
				layer.CreateTime = l.CreateTime
			}
		}

		if len(layer.Id) == 0 {
			layer.Id = ksuid.New().String()
		}
		if layer.CreateTime == nil {
			layer.CreateTime = core.Now()
		}

		layer.UpdateTime = core.Now()
	}

	if _, err := model.GetLayer().Create(ctx, in.Layers...); err != nil {
		return nil, core.NewInternalError("failed to create the layer to database, error: %s", err.Error())
	}

	resp := &pb.BatchCreateLayerResponse{}
	for _, l := range in.Layers {
		resp.Layers = append(resp.Layers, &tilestream.Layer{
			Id: l.Id,
		})
	}
	return resp, nil
}
