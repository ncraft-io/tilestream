
type Layer {
    id: String @1
    name: String @2
    type: String @3
    templated: Bool @4

    config: Object @5 @db.json
    description: String @6

    original_id: String @10 //< copied layer from original layer id

    create_time: Timestamp @100
    update_time: Timestamp @101
}