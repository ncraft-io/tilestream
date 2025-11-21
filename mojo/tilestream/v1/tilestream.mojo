
/// Tilestream service
interface Tilestream {
    @http.post("/tilestream/v1/layers/{layer}/tiles/{level}/{x}/{y}")
    create_tile (layer: String @1
                 level: Int32 @2
                 x: Int32 @3
                 y: Int32 @4
                 format: String @5)
                 -> Tile

    @http.post("/tilestream/v1/layers/{layer}/tiles")
    batch_create_tiles()
    create_tiles (layer: String @1
                  tiles: [Tile] @2 @http.body)

    @http.get("/tilestream/v1/layers/{layer}/tiles/{level}/{x}/{y}")
    get_tile (layer: String @1
                 level: Int32 @2
                 x: Int32 @3
                 y: Int32 @4
                 format: String @5)
                 -> Tile

    @http.get("/tilestream/v1/layers/{layer}/tile_info")
    get_tile_info (layer: String @1) -> TileInfo

    @http.put("/tilestream/v1/layers/{layer}/tiles/{level}/{x}/{y}")
    update_tile (layer: String @1
                 level: Int32 @2
                 x: Int32 @3
                 y: Int32 @4
                 format: String @5
                 tile: Tile @6 @http.body)

    @http.put("/tilestream/v1/layers/{layer}/tile_info")
    update_tile_info (layer: String @1
                      info: TileInfo @2 @http.body)

    @http.post("/tilestream/v1/layers")
    create_layer (layer: Layer @1 @http.body) -> Layer

    @http.post("/tilestream/v1/layers:batch")
    batch_create_layer (layers: [Layer] @2 @http.body) -> [Layer]

    @http.put("/tilestream/v1/layers/{id}")
    update_layer (id: String @1, layer: Layer @2 @http.body)

    @http.put("/tilestream/v1/layers:batch")
    batch_update_layer (layers: [Layer] @2 @http.body)

    @http.delete("/tilestream/v1/layers/{layer}")
    delete_layer (layer: String @1)

    @http.get("/tilestream/v1/layers/{layer}")
    get_layer (layer: String @1) -> Layer

    @http.get("/tilestream/v1/layers:batch")
    batch_get_layers(layers: [String] @1) -> [Layer]

    @http.get("/tilestream/v1/layers")
    @query
    @pagination
    list_layers (prefix: String @1) //< prefix of layer name
                 -> [Layer]
}