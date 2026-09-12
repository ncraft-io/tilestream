# MBTiles 命令行工具

在 `service-go` 目录运行：

```sh
go run ./cmd/mbtiles merge --src /data/input.mbtiles --dest /data/output.mbtiles
go run ./cmd/mbtiles merge --src /data/sources --dest /data/output.mbtiles
```

`merge` 可简写为 `m`，`--src` 和 `--dest` 可简写为 `-s` 和 `-d`，两者均为必填参数。目标文件的父目录需要已存在。

- 目录输入会递归查找 `.mbtiles` 文件，按规范化后的路径排序；目标文件及其文件别名会被排除。单文件输入不能与目标指向同一文件。
- 保留目标库中已有瓦片；同坐标的瓦片由后面的源文件整块覆盖，不合并瓦片内部的矢量要素。
- 直接复制 `tiles` 中的 TMS 坐标和二进制数据，包括空数据；不依赖 `bounds`、`minzoom`、`maxzoom` 枚举瓦片。源文件可以使用符合 MBTiles 表接口的视图。
- 保留原始元数据。普通同名键由后面的源文件覆盖；`bounds` 取并集，缩放范围从目标中实际存储的瓦片计算。若任一源文件缺少 `bounds`，则不输出可能遗漏覆盖区域的局部 bounds。
- 所有源文件必须包含 `name` 和 `format`，格式必须一致。PBF 还必须包含 `json.vector_layers`；同名图层的字段和缩放范围会合并，字段类型冲突时报错。
- 源库以只读方式打开。目标表和索引会自动初始化，整个合并使用一个事务；失败时回滚已有目标，新建的失败输出会被删除。

`Mbtiles.Tile` 和 `Mbtiles.WriteTile` 接收 XYZ 坐标，并在数据库边界转换为 TMS 行号。底层 `TilesModel` 和合并操作直接使用数据库中的 TMS 坐标。
