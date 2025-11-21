package io.ncraft.tilestream.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.63.0)",
    comments = "Source: tilestream/v1/tilestream.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class TilestreamGrpc {

  private TilestreamGrpc() {}

  public static final java.lang.String SERVICE_NAME = "tilestream.v1.Tilestream";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.CreateTileRequest,
      io.ncraft.tilestream.Tile> getCreateTileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "create_tile",
      requestType = io.ncraft.tilestream.v1.CreateTileRequest.class,
      responseType = io.ncraft.tilestream.Tile.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.CreateTileRequest,
      io.ncraft.tilestream.Tile> getCreateTileMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.CreateTileRequest, io.ncraft.tilestream.Tile> getCreateTileMethod;
    if ((getCreateTileMethod = TilestreamGrpc.getCreateTileMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getCreateTileMethod = TilestreamGrpc.getCreateTileMethod) == null) {
          TilestreamGrpc.getCreateTileMethod = getCreateTileMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.CreateTileRequest, io.ncraft.tilestream.Tile>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "create_tile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.CreateTileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.Tile.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("create_tile"))
              .build();
        }
      }
    }
    return getCreateTileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchCreateTilesRequest,
      org.mojolang.mojo.core.Null> getBatchCreateTilesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_create_tiles",
      requestType = io.ncraft.tilestream.v1.BatchCreateTilesRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchCreateTilesRequest,
      org.mojolang.mojo.core.Null> getBatchCreateTilesMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchCreateTilesRequest, org.mojolang.mojo.core.Null> getBatchCreateTilesMethod;
    if ((getBatchCreateTilesMethod = TilestreamGrpc.getBatchCreateTilesMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getBatchCreateTilesMethod = TilestreamGrpc.getBatchCreateTilesMethod) == null) {
          TilestreamGrpc.getBatchCreateTilesMethod = getBatchCreateTilesMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.BatchCreateTilesRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_create_tiles"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.BatchCreateTilesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("batch_create_tiles"))
              .build();
        }
      }
    }
    return getBatchCreateTilesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.CreateTilesRequest,
      org.mojolang.mojo.core.Null> getCreateTilesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "create_tiles",
      requestType = io.ncraft.tilestream.v1.CreateTilesRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.CreateTilesRequest,
      org.mojolang.mojo.core.Null> getCreateTilesMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.CreateTilesRequest, org.mojolang.mojo.core.Null> getCreateTilesMethod;
    if ((getCreateTilesMethod = TilestreamGrpc.getCreateTilesMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getCreateTilesMethod = TilestreamGrpc.getCreateTilesMethod) == null) {
          TilestreamGrpc.getCreateTilesMethod = getCreateTilesMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.CreateTilesRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "create_tiles"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.CreateTilesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("create_tiles"))
              .build();
        }
      }
    }
    return getCreateTilesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.GetTileRequest,
      io.ncraft.tilestream.Tile> getGetTileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "get_tile",
      requestType = io.ncraft.tilestream.v1.GetTileRequest.class,
      responseType = io.ncraft.tilestream.Tile.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.GetTileRequest,
      io.ncraft.tilestream.Tile> getGetTileMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.GetTileRequest, io.ncraft.tilestream.Tile> getGetTileMethod;
    if ((getGetTileMethod = TilestreamGrpc.getGetTileMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getGetTileMethod = TilestreamGrpc.getGetTileMethod) == null) {
          TilestreamGrpc.getGetTileMethod = getGetTileMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.GetTileRequest, io.ncraft.tilestream.Tile>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "get_tile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.GetTileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.Tile.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("get_tile"))
              .build();
        }
      }
    }
    return getGetTileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.GetTileInfoRequest,
      io.ncraft.tilestream.TileInfo> getGetTileInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "get_tile_info",
      requestType = io.ncraft.tilestream.v1.GetTileInfoRequest.class,
      responseType = io.ncraft.tilestream.TileInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.GetTileInfoRequest,
      io.ncraft.tilestream.TileInfo> getGetTileInfoMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.GetTileInfoRequest, io.ncraft.tilestream.TileInfo> getGetTileInfoMethod;
    if ((getGetTileInfoMethod = TilestreamGrpc.getGetTileInfoMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getGetTileInfoMethod = TilestreamGrpc.getGetTileInfoMethod) == null) {
          TilestreamGrpc.getGetTileInfoMethod = getGetTileInfoMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.GetTileInfoRequest, io.ncraft.tilestream.TileInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "get_tile_info"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.GetTileInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.TileInfo.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("get_tile_info"))
              .build();
        }
      }
    }
    return getGetTileInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.UpdateTileRequest,
      org.mojolang.mojo.core.Null> getUpdateTileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "update_tile",
      requestType = io.ncraft.tilestream.v1.UpdateTileRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.UpdateTileRequest,
      org.mojolang.mojo.core.Null> getUpdateTileMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.UpdateTileRequest, org.mojolang.mojo.core.Null> getUpdateTileMethod;
    if ((getUpdateTileMethod = TilestreamGrpc.getUpdateTileMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getUpdateTileMethod = TilestreamGrpc.getUpdateTileMethod) == null) {
          TilestreamGrpc.getUpdateTileMethod = getUpdateTileMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.UpdateTileRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "update_tile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.UpdateTileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("update_tile"))
              .build();
        }
      }
    }
    return getUpdateTileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.UpdateTileInfoRequest,
      org.mojolang.mojo.core.Null> getUpdateTileInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "update_tile_info",
      requestType = io.ncraft.tilestream.v1.UpdateTileInfoRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.UpdateTileInfoRequest,
      org.mojolang.mojo.core.Null> getUpdateTileInfoMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.UpdateTileInfoRequest, org.mojolang.mojo.core.Null> getUpdateTileInfoMethod;
    if ((getUpdateTileInfoMethod = TilestreamGrpc.getUpdateTileInfoMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getUpdateTileInfoMethod = TilestreamGrpc.getUpdateTileInfoMethod) == null) {
          TilestreamGrpc.getUpdateTileInfoMethod = getUpdateTileInfoMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.UpdateTileInfoRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "update_tile_info"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.UpdateTileInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("update_tile_info"))
              .build();
        }
      }
    }
    return getUpdateTileInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.CreateLayerRequest,
      io.ncraft.tilestream.Layer> getCreateLayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "create_layer",
      requestType = io.ncraft.tilestream.v1.CreateLayerRequest.class,
      responseType = io.ncraft.tilestream.Layer.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.CreateLayerRequest,
      io.ncraft.tilestream.Layer> getCreateLayerMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.CreateLayerRequest, io.ncraft.tilestream.Layer> getCreateLayerMethod;
    if ((getCreateLayerMethod = TilestreamGrpc.getCreateLayerMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getCreateLayerMethod = TilestreamGrpc.getCreateLayerMethod) == null) {
          TilestreamGrpc.getCreateLayerMethod = getCreateLayerMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.CreateLayerRequest, io.ncraft.tilestream.Layer>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "create_layer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.CreateLayerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.Layer.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("create_layer"))
              .build();
        }
      }
    }
    return getCreateLayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchCreateLayerRequest,
      io.ncraft.tilestream.v1.BatchCreateLayerResponse> getBatchCreateLayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_create_layer",
      requestType = io.ncraft.tilestream.v1.BatchCreateLayerRequest.class,
      responseType = io.ncraft.tilestream.v1.BatchCreateLayerResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchCreateLayerRequest,
      io.ncraft.tilestream.v1.BatchCreateLayerResponse> getBatchCreateLayerMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchCreateLayerRequest, io.ncraft.tilestream.v1.BatchCreateLayerResponse> getBatchCreateLayerMethod;
    if ((getBatchCreateLayerMethod = TilestreamGrpc.getBatchCreateLayerMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getBatchCreateLayerMethod = TilestreamGrpc.getBatchCreateLayerMethod) == null) {
          TilestreamGrpc.getBatchCreateLayerMethod = getBatchCreateLayerMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.BatchCreateLayerRequest, io.ncraft.tilestream.v1.BatchCreateLayerResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_create_layer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.BatchCreateLayerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.BatchCreateLayerResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("batch_create_layer"))
              .build();
        }
      }
    }
    return getBatchCreateLayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.UpdateLayerRequest,
      org.mojolang.mojo.core.Null> getUpdateLayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "update_layer",
      requestType = io.ncraft.tilestream.v1.UpdateLayerRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.UpdateLayerRequest,
      org.mojolang.mojo.core.Null> getUpdateLayerMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.UpdateLayerRequest, org.mojolang.mojo.core.Null> getUpdateLayerMethod;
    if ((getUpdateLayerMethod = TilestreamGrpc.getUpdateLayerMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getUpdateLayerMethod = TilestreamGrpc.getUpdateLayerMethod) == null) {
          TilestreamGrpc.getUpdateLayerMethod = getUpdateLayerMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.UpdateLayerRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "update_layer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.UpdateLayerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("update_layer"))
              .build();
        }
      }
    }
    return getUpdateLayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchUpdateLayerRequest,
      org.mojolang.mojo.core.Null> getBatchUpdateLayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_update_layer",
      requestType = io.ncraft.tilestream.v1.BatchUpdateLayerRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchUpdateLayerRequest,
      org.mojolang.mojo.core.Null> getBatchUpdateLayerMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchUpdateLayerRequest, org.mojolang.mojo.core.Null> getBatchUpdateLayerMethod;
    if ((getBatchUpdateLayerMethod = TilestreamGrpc.getBatchUpdateLayerMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getBatchUpdateLayerMethod = TilestreamGrpc.getBatchUpdateLayerMethod) == null) {
          TilestreamGrpc.getBatchUpdateLayerMethod = getBatchUpdateLayerMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.BatchUpdateLayerRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_update_layer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.BatchUpdateLayerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("batch_update_layer"))
              .build();
        }
      }
    }
    return getBatchUpdateLayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.DeleteLayerRequest,
      org.mojolang.mojo.core.Null> getDeleteLayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "delete_layer",
      requestType = io.ncraft.tilestream.v1.DeleteLayerRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.DeleteLayerRequest,
      org.mojolang.mojo.core.Null> getDeleteLayerMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.DeleteLayerRequest, org.mojolang.mojo.core.Null> getDeleteLayerMethod;
    if ((getDeleteLayerMethod = TilestreamGrpc.getDeleteLayerMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getDeleteLayerMethod = TilestreamGrpc.getDeleteLayerMethod) == null) {
          TilestreamGrpc.getDeleteLayerMethod = getDeleteLayerMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.DeleteLayerRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "delete_layer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.DeleteLayerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("delete_layer"))
              .build();
        }
      }
    }
    return getDeleteLayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.GetLayerRequest,
      io.ncraft.tilestream.Layer> getGetLayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "get_layer",
      requestType = io.ncraft.tilestream.v1.GetLayerRequest.class,
      responseType = io.ncraft.tilestream.Layer.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.GetLayerRequest,
      io.ncraft.tilestream.Layer> getGetLayerMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.GetLayerRequest, io.ncraft.tilestream.Layer> getGetLayerMethod;
    if ((getGetLayerMethod = TilestreamGrpc.getGetLayerMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getGetLayerMethod = TilestreamGrpc.getGetLayerMethod) == null) {
          TilestreamGrpc.getGetLayerMethod = getGetLayerMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.GetLayerRequest, io.ncraft.tilestream.Layer>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "get_layer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.GetLayerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.Layer.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("get_layer"))
              .build();
        }
      }
    }
    return getGetLayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchGetLayersRequest,
      io.ncraft.tilestream.v1.BatchGetLayersResponse> getBatchGetLayersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_get_layers",
      requestType = io.ncraft.tilestream.v1.BatchGetLayersRequest.class,
      responseType = io.ncraft.tilestream.v1.BatchGetLayersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchGetLayersRequest,
      io.ncraft.tilestream.v1.BatchGetLayersResponse> getBatchGetLayersMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.BatchGetLayersRequest, io.ncraft.tilestream.v1.BatchGetLayersResponse> getBatchGetLayersMethod;
    if ((getBatchGetLayersMethod = TilestreamGrpc.getBatchGetLayersMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getBatchGetLayersMethod = TilestreamGrpc.getBatchGetLayersMethod) == null) {
          TilestreamGrpc.getBatchGetLayersMethod = getBatchGetLayersMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.BatchGetLayersRequest, io.ncraft.tilestream.v1.BatchGetLayersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_get_layers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.BatchGetLayersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.BatchGetLayersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("batch_get_layers"))
              .build();
        }
      }
    }
    return getBatchGetLayersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.ListLayersRequest,
      io.ncraft.tilestream.v1.ListLayersResponse> getListLayersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "list_layers",
      requestType = io.ncraft.tilestream.v1.ListLayersRequest.class,
      responseType = io.ncraft.tilestream.v1.ListLayersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.ListLayersRequest,
      io.ncraft.tilestream.v1.ListLayersResponse> getListLayersMethod() {
    io.grpc.MethodDescriptor<io.ncraft.tilestream.v1.ListLayersRequest, io.ncraft.tilestream.v1.ListLayersResponse> getListLayersMethod;
    if ((getListLayersMethod = TilestreamGrpc.getListLayersMethod) == null) {
      synchronized (TilestreamGrpc.class) {
        if ((getListLayersMethod = TilestreamGrpc.getListLayersMethod) == null) {
          TilestreamGrpc.getListLayersMethod = getListLayersMethod =
              io.grpc.MethodDescriptor.<io.ncraft.tilestream.v1.ListLayersRequest, io.ncraft.tilestream.v1.ListLayersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "list_layers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.ListLayersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.tilestream.v1.ListLayersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TilestreamMethodDescriptorSupplier("list_layers"))
              .build();
        }
      }
    }
    return getListLayersMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static TilestreamStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TilestreamStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TilestreamStub>() {
        @java.lang.Override
        public TilestreamStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TilestreamStub(channel, callOptions);
        }
      };
    return TilestreamStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static TilestreamBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TilestreamBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TilestreamBlockingStub>() {
        @java.lang.Override
        public TilestreamBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TilestreamBlockingStub(channel, callOptions);
        }
      };
    return TilestreamBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static TilestreamFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TilestreamFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TilestreamFutureStub>() {
        @java.lang.Override
        public TilestreamFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TilestreamFutureStub(channel, callOptions);
        }
      };
    return TilestreamFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void createTile(io.ncraft.tilestream.v1.CreateTileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.Tile> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateTileMethod(), responseObserver);
    }

    /**
     */
    default void batchCreateTiles(io.ncraft.tilestream.v1.BatchCreateTilesRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchCreateTilesMethod(), responseObserver);
    }

    /**
     */
    default void createTiles(io.ncraft.tilestream.v1.CreateTilesRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateTilesMethod(), responseObserver);
    }

    /**
     */
    default void getTile(io.ncraft.tilestream.v1.GetTileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.Tile> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTileMethod(), responseObserver);
    }

    /**
     */
    default void getTileInfo(io.ncraft.tilestream.v1.GetTileInfoRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.TileInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTileInfoMethod(), responseObserver);
    }

    /**
     */
    default void updateTile(io.ncraft.tilestream.v1.UpdateTileRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateTileMethod(), responseObserver);
    }

    /**
     */
    default void updateTileInfo(io.ncraft.tilestream.v1.UpdateTileInfoRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateTileInfoMethod(), responseObserver);
    }

    /**
     */
    default void createLayer(io.ncraft.tilestream.v1.CreateLayerRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.Layer> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateLayerMethod(), responseObserver);
    }

    /**
     */
    default void batchCreateLayer(io.ncraft.tilestream.v1.BatchCreateLayerRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.v1.BatchCreateLayerResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchCreateLayerMethod(), responseObserver);
    }

    /**
     */
    default void updateLayer(io.ncraft.tilestream.v1.UpdateLayerRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateLayerMethod(), responseObserver);
    }

    /**
     */
    default void batchUpdateLayer(io.ncraft.tilestream.v1.BatchUpdateLayerRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchUpdateLayerMethod(), responseObserver);
    }

    /**
     */
    default void deleteLayer(io.ncraft.tilestream.v1.DeleteLayerRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteLayerMethod(), responseObserver);
    }

    /**
     */
    default void getLayer(io.ncraft.tilestream.v1.GetLayerRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.Layer> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetLayerMethod(), responseObserver);
    }

    /**
     */
    default void batchGetLayers(io.ncraft.tilestream.v1.BatchGetLayersRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.v1.BatchGetLayersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchGetLayersMethod(), responseObserver);
    }

    /**
     */
    default void listLayers(io.ncraft.tilestream.v1.ListLayersRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.v1.ListLayersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListLayersMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service Tilestream.
   */
  public static abstract class TilestreamImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return TilestreamGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service Tilestream.
   */
  public static final class TilestreamStub
      extends io.grpc.stub.AbstractAsyncStub<TilestreamStub> {
    private TilestreamStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TilestreamStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TilestreamStub(channel, callOptions);
    }

    /**
     */
    public void createTile(io.ncraft.tilestream.v1.CreateTileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.Tile> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateTileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchCreateTiles(io.ncraft.tilestream.v1.BatchCreateTilesRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchCreateTilesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createTiles(io.ncraft.tilestream.v1.CreateTilesRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateTilesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTile(io.ncraft.tilestream.v1.GetTileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.Tile> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTileInfo(io.ncraft.tilestream.v1.GetTileInfoRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.TileInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTileInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateTile(io.ncraft.tilestream.v1.UpdateTileRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateTileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateTileInfo(io.ncraft.tilestream.v1.UpdateTileInfoRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateTileInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createLayer(io.ncraft.tilestream.v1.CreateLayerRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.Layer> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateLayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchCreateLayer(io.ncraft.tilestream.v1.BatchCreateLayerRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.v1.BatchCreateLayerResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchCreateLayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateLayer(io.ncraft.tilestream.v1.UpdateLayerRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateLayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchUpdateLayer(io.ncraft.tilestream.v1.BatchUpdateLayerRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchUpdateLayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteLayer(io.ncraft.tilestream.v1.DeleteLayerRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteLayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getLayer(io.ncraft.tilestream.v1.GetLayerRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.Layer> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetLayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchGetLayers(io.ncraft.tilestream.v1.BatchGetLayersRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.v1.BatchGetLayersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchGetLayersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listLayers(io.ncraft.tilestream.v1.ListLayersRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.tilestream.v1.ListLayersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListLayersMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service Tilestream.
   */
  public static final class TilestreamBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<TilestreamBlockingStub> {
    private TilestreamBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TilestreamBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TilestreamBlockingStub(channel, callOptions);
    }

    /**
     */
    public io.ncraft.tilestream.Tile createTile(io.ncraft.tilestream.v1.CreateTileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateTileMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null batchCreateTiles(io.ncraft.tilestream.v1.BatchCreateTilesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchCreateTilesMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null createTiles(io.ncraft.tilestream.v1.CreateTilesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateTilesMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.tilestream.Tile getTile(io.ncraft.tilestream.v1.GetTileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTileMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.tilestream.TileInfo getTileInfo(io.ncraft.tilestream.v1.GetTileInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTileInfoMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null updateTile(io.ncraft.tilestream.v1.UpdateTileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateTileMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null updateTileInfo(io.ncraft.tilestream.v1.UpdateTileInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateTileInfoMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.tilestream.Layer createLayer(io.ncraft.tilestream.v1.CreateLayerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateLayerMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.tilestream.v1.BatchCreateLayerResponse batchCreateLayer(io.ncraft.tilestream.v1.BatchCreateLayerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchCreateLayerMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null updateLayer(io.ncraft.tilestream.v1.UpdateLayerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateLayerMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null batchUpdateLayer(io.ncraft.tilestream.v1.BatchUpdateLayerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchUpdateLayerMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null deleteLayer(io.ncraft.tilestream.v1.DeleteLayerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteLayerMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.tilestream.Layer getLayer(io.ncraft.tilestream.v1.GetLayerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetLayerMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.tilestream.v1.BatchGetLayersResponse batchGetLayers(io.ncraft.tilestream.v1.BatchGetLayersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchGetLayersMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.tilestream.v1.ListLayersResponse listLayers(io.ncraft.tilestream.v1.ListLayersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListLayersMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service Tilestream.
   */
  public static final class TilestreamFutureStub
      extends io.grpc.stub.AbstractFutureStub<TilestreamFutureStub> {
    private TilestreamFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TilestreamFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TilestreamFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.tilestream.Tile> createTile(
        io.ncraft.tilestream.v1.CreateTileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateTileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> batchCreateTiles(
        io.ncraft.tilestream.v1.BatchCreateTilesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchCreateTilesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> createTiles(
        io.ncraft.tilestream.v1.CreateTilesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateTilesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.tilestream.Tile> getTile(
        io.ncraft.tilestream.v1.GetTileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.tilestream.TileInfo> getTileInfo(
        io.ncraft.tilestream.v1.GetTileInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTileInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> updateTile(
        io.ncraft.tilestream.v1.UpdateTileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateTileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> updateTileInfo(
        io.ncraft.tilestream.v1.UpdateTileInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateTileInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.tilestream.Layer> createLayer(
        io.ncraft.tilestream.v1.CreateLayerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateLayerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.tilestream.v1.BatchCreateLayerResponse> batchCreateLayer(
        io.ncraft.tilestream.v1.BatchCreateLayerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchCreateLayerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> updateLayer(
        io.ncraft.tilestream.v1.UpdateLayerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateLayerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> batchUpdateLayer(
        io.ncraft.tilestream.v1.BatchUpdateLayerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchUpdateLayerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> deleteLayer(
        io.ncraft.tilestream.v1.DeleteLayerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteLayerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.tilestream.Layer> getLayer(
        io.ncraft.tilestream.v1.GetLayerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetLayerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.tilestream.v1.BatchGetLayersResponse> batchGetLayers(
        io.ncraft.tilestream.v1.BatchGetLayersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchGetLayersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.tilestream.v1.ListLayersResponse> listLayers(
        io.ncraft.tilestream.v1.ListLayersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListLayersMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_TILE = 0;
  private static final int METHODID_BATCH_CREATE_TILES = 1;
  private static final int METHODID_CREATE_TILES = 2;
  private static final int METHODID_GET_TILE = 3;
  private static final int METHODID_GET_TILE_INFO = 4;
  private static final int METHODID_UPDATE_TILE = 5;
  private static final int METHODID_UPDATE_TILE_INFO = 6;
  private static final int METHODID_CREATE_LAYER = 7;
  private static final int METHODID_BATCH_CREATE_LAYER = 8;
  private static final int METHODID_UPDATE_LAYER = 9;
  private static final int METHODID_BATCH_UPDATE_LAYER = 10;
  private static final int METHODID_DELETE_LAYER = 11;
  private static final int METHODID_GET_LAYER = 12;
  private static final int METHODID_BATCH_GET_LAYERS = 13;
  private static final int METHODID_LIST_LAYERS = 14;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CREATE_TILE:
          serviceImpl.createTile((io.ncraft.tilestream.v1.CreateTileRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.tilestream.Tile>) responseObserver);
          break;
        case METHODID_BATCH_CREATE_TILES:
          serviceImpl.batchCreateTiles((io.ncraft.tilestream.v1.BatchCreateTilesRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_CREATE_TILES:
          serviceImpl.createTiles((io.ncraft.tilestream.v1.CreateTilesRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_GET_TILE:
          serviceImpl.getTile((io.ncraft.tilestream.v1.GetTileRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.tilestream.Tile>) responseObserver);
          break;
        case METHODID_GET_TILE_INFO:
          serviceImpl.getTileInfo((io.ncraft.tilestream.v1.GetTileInfoRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.tilestream.TileInfo>) responseObserver);
          break;
        case METHODID_UPDATE_TILE:
          serviceImpl.updateTile((io.ncraft.tilestream.v1.UpdateTileRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_UPDATE_TILE_INFO:
          serviceImpl.updateTileInfo((io.ncraft.tilestream.v1.UpdateTileInfoRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_CREATE_LAYER:
          serviceImpl.createLayer((io.ncraft.tilestream.v1.CreateLayerRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.tilestream.Layer>) responseObserver);
          break;
        case METHODID_BATCH_CREATE_LAYER:
          serviceImpl.batchCreateLayer((io.ncraft.tilestream.v1.BatchCreateLayerRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.tilestream.v1.BatchCreateLayerResponse>) responseObserver);
          break;
        case METHODID_UPDATE_LAYER:
          serviceImpl.updateLayer((io.ncraft.tilestream.v1.UpdateLayerRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_BATCH_UPDATE_LAYER:
          serviceImpl.batchUpdateLayer((io.ncraft.tilestream.v1.BatchUpdateLayerRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_DELETE_LAYER:
          serviceImpl.deleteLayer((io.ncraft.tilestream.v1.DeleteLayerRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_GET_LAYER:
          serviceImpl.getLayer((io.ncraft.tilestream.v1.GetLayerRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.tilestream.Layer>) responseObserver);
          break;
        case METHODID_BATCH_GET_LAYERS:
          serviceImpl.batchGetLayers((io.ncraft.tilestream.v1.BatchGetLayersRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.tilestream.v1.BatchGetLayersResponse>) responseObserver);
          break;
        case METHODID_LIST_LAYERS:
          serviceImpl.listLayers((io.ncraft.tilestream.v1.ListLayersRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.tilestream.v1.ListLayersResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getCreateTileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.CreateTileRequest,
              io.ncraft.tilestream.Tile>(
                service, METHODID_CREATE_TILE)))
        .addMethod(
          getBatchCreateTilesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.BatchCreateTilesRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_BATCH_CREATE_TILES)))
        .addMethod(
          getCreateTilesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.CreateTilesRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_CREATE_TILES)))
        .addMethod(
          getGetTileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.GetTileRequest,
              io.ncraft.tilestream.Tile>(
                service, METHODID_GET_TILE)))
        .addMethod(
          getGetTileInfoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.GetTileInfoRequest,
              io.ncraft.tilestream.TileInfo>(
                service, METHODID_GET_TILE_INFO)))
        .addMethod(
          getUpdateTileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.UpdateTileRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_UPDATE_TILE)))
        .addMethod(
          getUpdateTileInfoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.UpdateTileInfoRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_UPDATE_TILE_INFO)))
        .addMethod(
          getCreateLayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.CreateLayerRequest,
              io.ncraft.tilestream.Layer>(
                service, METHODID_CREATE_LAYER)))
        .addMethod(
          getBatchCreateLayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.BatchCreateLayerRequest,
              io.ncraft.tilestream.v1.BatchCreateLayerResponse>(
                service, METHODID_BATCH_CREATE_LAYER)))
        .addMethod(
          getUpdateLayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.UpdateLayerRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_UPDATE_LAYER)))
        .addMethod(
          getBatchUpdateLayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.BatchUpdateLayerRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_BATCH_UPDATE_LAYER)))
        .addMethod(
          getDeleteLayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.DeleteLayerRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_DELETE_LAYER)))
        .addMethod(
          getGetLayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.GetLayerRequest,
              io.ncraft.tilestream.Layer>(
                service, METHODID_GET_LAYER)))
        .addMethod(
          getBatchGetLayersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.BatchGetLayersRequest,
              io.ncraft.tilestream.v1.BatchGetLayersResponse>(
                service, METHODID_BATCH_GET_LAYERS)))
        .addMethod(
          getListLayersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.tilestream.v1.ListLayersRequest,
              io.ncraft.tilestream.v1.ListLayersResponse>(
                service, METHODID_LIST_LAYERS)))
        .build();
  }

  private static abstract class TilestreamBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    TilestreamBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.ncraft.tilestream.v1.TilestreamProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Tilestream");
    }
  }

  private static final class TilestreamFileDescriptorSupplier
      extends TilestreamBaseDescriptorSupplier {
    TilestreamFileDescriptorSupplier() {}
  }

  private static final class TilestreamMethodDescriptorSupplier
      extends TilestreamBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    TilestreamMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (TilestreamGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new TilestreamFileDescriptorSupplier())
              .addMethod(getCreateTileMethod())
              .addMethod(getBatchCreateTilesMethod())
              .addMethod(getCreateTilesMethod())
              .addMethod(getGetTileMethod())
              .addMethod(getGetTileInfoMethod())
              .addMethod(getUpdateTileMethod())
              .addMethod(getUpdateTileInfoMethod())
              .addMethod(getCreateLayerMethod())
              .addMethod(getBatchCreateLayerMethod())
              .addMethod(getUpdateLayerMethod())
              .addMethod(getBatchUpdateLayerMethod())
              .addMethod(getDeleteLayerMethod())
              .addMethod(getGetLayerMethod())
              .addMethod(getBatchGetLayersMethod())
              .addMethod(getListLayersMethod())
              .build();
        }
      }
    }
    return result;
  }
}
