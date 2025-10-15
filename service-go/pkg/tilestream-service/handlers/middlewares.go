package handlers

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/ncraft-io/ncraft-gokit/pkg/middleware"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/ncraft-io/tilestream/go/pkg/tilestream"

	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream-service/svc"

	// this service api
	pb "github.com/ncraft-io/tilestream/go/pkg/tilestream/v1"
)

var (
	_ = tilestream.Tile{}
	_ = core.Null{}
	_ = tilestream.TileInfo{}
	_ = tilestream.Layer{}
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints, options map[string]interface{}) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/ncraft-io//_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	var tracer stdopentracing.Tracer
	if value, ok := options["tracer"]; ok && value != nil {
		tracer = value.(stdopentracing.Tracer)
	}
	var count *kitprometheus.Counter
	if value, ok := options["count"]; ok && value != nil {
		count = value.(*kitprometheus.Counter)
	}
	var latency *kitprometheus.Histogram
	if value, ok := options["latency"]; ok && value != nil {
		latency = value.(*kitprometheus.Histogram)
	}
	//var validator *middleware.Validator
	//if value, ok := options["validator"]; ok && value != nil {
	//	validator = value.(*middleware.Validator)
	//}

	{ // create_tile
		if tracer != nil {
			in.CreateTileEndpoint = opentracing.TraceServer(tracer, "create_tile")(in.CreateTileEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateTileEndpoint = middleware.Instrumenting(latency.With("method", "create_tile"), count.With("method", "create_tile"))(in.CreateTileEndpoint)
		}
		//if validator != nil {
		//	in.CreateTileEndpoint = validator.Validate()(in.CreateTileEndpoint)
		//}
	}
	{ // batch_create_tiles
		if tracer != nil {
			in.BatchCreateTilesEndpoint = opentracing.TraceServer(tracer, "batch_create_tiles")(in.BatchCreateTilesEndpoint)
		}
		if count != nil && latency != nil {
			in.BatchCreateTilesEndpoint = middleware.Instrumenting(latency.With("method", "batch_create_tiles"), count.With("method", "batch_create_tiles"))(in.BatchCreateTilesEndpoint)
		}
		//if validator != nil {
		//	in.BatchCreateTilesEndpoint = validator.Validate()(in.BatchCreateTilesEndpoint)
		//}
	}
	{ // create_tiles
		if tracer != nil {
			in.CreateTilesEndpoint = opentracing.TraceServer(tracer, "create_tiles")(in.CreateTilesEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateTilesEndpoint = middleware.Instrumenting(latency.With("method", "create_tiles"), count.With("method", "create_tiles"))(in.CreateTilesEndpoint)
		}
		//if validator != nil {
		//	in.CreateTilesEndpoint = validator.Validate()(in.CreateTilesEndpoint)
		//}
	}
	{ // get_tile
		if tracer != nil {
			in.GetTileEndpoint = opentracing.TraceServer(tracer, "get_tile")(in.GetTileEndpoint)
		}
		if count != nil && latency != nil {
			in.GetTileEndpoint = middleware.Instrumenting(latency.With("method", "get_tile"), count.With("method", "get_tile"))(in.GetTileEndpoint)
		}
		//if validator != nil {
		//	in.GetTileEndpoint = validator.Validate()(in.GetTileEndpoint)
		//}
	}
	{ // get_tile_info
		if tracer != nil {
			in.GetTileInfoEndpoint = opentracing.TraceServer(tracer, "get_tile_info")(in.GetTileInfoEndpoint)
		}
		if count != nil && latency != nil {
			in.GetTileInfoEndpoint = middleware.Instrumenting(latency.With("method", "get_tile_info"), count.With("method", "get_tile_info"))(in.GetTileInfoEndpoint)
		}
		//if validator != nil {
		//	in.GetTileInfoEndpoint = validator.Validate()(in.GetTileInfoEndpoint)
		//}
	}
	{ // update_tile
		if tracer != nil {
			in.UpdateTileEndpoint = opentracing.TraceServer(tracer, "update_tile")(in.UpdateTileEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateTileEndpoint = middleware.Instrumenting(latency.With("method", "update_tile"), count.With("method", "update_tile"))(in.UpdateTileEndpoint)
		}
		//if validator != nil {
		//	in.UpdateTileEndpoint = validator.Validate()(in.UpdateTileEndpoint)
		//}
	}
	{ // update_tile_info
		if tracer != nil {
			in.UpdateTileInfoEndpoint = opentracing.TraceServer(tracer, "update_tile_info")(in.UpdateTileInfoEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateTileInfoEndpoint = middleware.Instrumenting(latency.With("method", "update_tile_info"), count.With("method", "update_tile_info"))(in.UpdateTileInfoEndpoint)
		}
		//if validator != nil {
		//	in.UpdateTileInfoEndpoint = validator.Validate()(in.UpdateTileInfoEndpoint)
		//}
	}
	{ // create_layer
		if tracer != nil {
			in.CreateLayerEndpoint = opentracing.TraceServer(tracer, "create_layer")(in.CreateLayerEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateLayerEndpoint = middleware.Instrumenting(latency.With("method", "create_layer"), count.With("method", "create_layer"))(in.CreateLayerEndpoint)
		}
		//if validator != nil {
		//	in.CreateLayerEndpoint = validator.Validate()(in.CreateLayerEndpoint)
		//}
	}
	{ // update_layer
		if tracer != nil {
			in.UpdateLayerEndpoint = opentracing.TraceServer(tracer, "update_layer")(in.UpdateLayerEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateLayerEndpoint = middleware.Instrumenting(latency.With("method", "update_layer"), count.With("method", "update_layer"))(in.UpdateLayerEndpoint)
		}
		//if validator != nil {
		//	in.UpdateLayerEndpoint = validator.Validate()(in.UpdateLayerEndpoint)
		//}
	}
	{ // delete_layer
		if tracer != nil {
			in.DeleteLayerEndpoint = opentracing.TraceServer(tracer, "delete_layer")(in.DeleteLayerEndpoint)
		}
		if count != nil && latency != nil {
			in.DeleteLayerEndpoint = middleware.Instrumenting(latency.With("method", "delete_layer"), count.With("method", "delete_layer"))(in.DeleteLayerEndpoint)
		}
		//if validator != nil {
		//	in.DeleteLayerEndpoint = validator.Validate()(in.DeleteLayerEndpoint)
		//}
	}
	{ // get_layer
		if tracer != nil {
			in.GetLayerEndpoint = opentracing.TraceServer(tracer, "get_layer")(in.GetLayerEndpoint)
		}
		if count != nil && latency != nil {
			in.GetLayerEndpoint = middleware.Instrumenting(latency.With("method", "get_layer"), count.With("method", "get_layer"))(in.GetLayerEndpoint)
		}
		//if validator != nil {
		//	in.GetLayerEndpoint = validator.Validate()(in.GetLayerEndpoint)
		//}
	}
	{ // list_layers
		if tracer != nil {
			in.ListLayersEndpoint = opentracing.TraceServer(tracer, "list_layers")(in.ListLayersEndpoint)
		}
		if count != nil && latency != nil {
			in.ListLayersEndpoint = middleware.Instrumenting(latency.With("method", "list_layers"), count.With("method", "list_layers"))(in.ListLayersEndpoint)
		}
		//if validator != nil {
		//	in.ListLayersEndpoint = validator.Validate()(in.ListLayersEndpoint)
		//}
	}

	return in
}

func WrapService(in pb.TilestreamServer, options map[string]interface{}) pb.TilestreamServer {
	_ = options
	return in
}
