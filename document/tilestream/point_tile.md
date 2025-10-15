| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `version` | `integer` | `Int32` | N |  | the version of the point tile specification |
| `id` | `string` |  | N |  | tile quad-key |
| `type` | `string` |  | N |  | the point data typethe user defined point data type |
| `keys` | `Array<string>` |  | N |  | Dictionary encoding for keys |
| `values` | `Array<tilestream.VectorTile.Value>` |  | N |  | Dictionary encoding for values |
| `rawValues` | `Array<integer>` |  | N |  | attribute index for which value is raw integer |
| `ids` | `Array<integer>` |  | N |  | for string ids, using keys or md5(string id) |
| `tags` | `Array<integer>` |  | N |  | Tags of this feature are encoded as repeated pairs ofintegers. |
| `xs` | `Array<integer>` |  | N |  | the coordinates for the point, using E7 of the latitude & longitude |
| `ys` | `Array<integer>` |  | N |  |
