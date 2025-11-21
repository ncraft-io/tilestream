
type Config {
    type Sql {
        type Field {
            name: String @1
            alias: String @2
        }

        table: String @1
        fields: [Field] @2
        filter: String @3
    }

    type Provider {
        uri: String @1
        id_field: String @2
        geometry_field: String @3
        sql: Sql @4
    }

    debug: Bool @1
    name: String @2

    min_zoom: Int @4
    max_zoom: Int @5
    bounds: geom.BoundingBox @6

    tile_buffer: Int @7
    provider: Provider @10
}
