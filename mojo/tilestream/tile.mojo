
type Tile {
    x: Int32 @1
    y: Int32 @2
    level: Int32 @3

    format: String @10
    encoding: String @11
    content: Bytes @15
}
