
# Tilestream service
[TOC]

## 整体说明
1.	字符串都为utf8格式;
1.	HTTP Headers:
	1.	Content-Type设置为：application/json
1.	DataTime格式参考RFC3339标准

## 错误处理
错误的具体信息将在error字段中返回。

### 错误码示例
```json
{
    "code": "400",
    "message": "Param Error"
}
```


### 状态码列表
| 状态码 | 说明 |
|---|---|
| 200 | 返回正常 |
| 400 | 参数错误 |
| 401 | 无access<br> key或key无效 |
| 500 | 服务器内部错误 |


## 

### 请求路径
```http
GET /tilestream/v1/layers
```


### 请求参数

#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `prefix` | `string` |  | 否 |  | prefix of layer name |
| `show_config` | `boolean` |  | 否 |  | show config of layer, hide layer config in most cases |


### 返回值

#### 返回对象
| type | description |
|---|---|
| `Array<tilestream.Layer>` |  |


#### `tilestream.Layer`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `type` | `string` |  | N |  |
| `templated` | `boolean` |  | N |  |
| `config` | `mojo.core.Object` |  | N |  | Object type |
| `description` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


## 

### 请求路径
```http
POST /tilestream/v1/layers
```


### 请求参数

#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `type` | `string` |  | N |  |
| `templated` | `boolean` |  | N |  |
| `config` | `mojo.core.Object` |  | N |  | Object type |
| `description` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `type` | `string` |  | N |  |
| `templated` | `boolean` |  | N |  |
| `config` | `mojo.core.Object` |  | N |  | Object type |
| `description` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


## 

### 请求路径
```http
PUT /tilestream/v1/layers/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `type` | `string` |  | N |  |
| `templated` | `boolean` |  | N |  |
| `config` | `mojo.core.Object` |  | N |  | Object type |
| `description` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


### 返回值

#### 返回对象
对象为空

## 

### 请求路径
```http
GET /tilestream/v1/layers/{layer}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `type` | `string` |  | N |  |
| `templated` | `boolean` |  | N |  |
| `config` | `mojo.core.Object` |  | N |  | Object type |
| `description` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


## 

### 请求路径
```http
DELETE /tilestream/v1/layers/{layer}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |


### 返回值

#### 返回对象
对象为空

## 

### 请求路径
```http
GET /tilestream/v1/layers/{layer}/tile_info
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `type` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `format` | `string` |  | N |  |
| `version` | `string` |  | N |  |
| `attribution` | `string` |  | N |  |
| `scheme` | `string` |  | N |  | tms, xyz |
| `minZoom` | `integer` | `Int32` | N |  |
| `maxZoom` | `integer` | `Int32` | N |  |
| `bounds` | `mojo.geom.BoundingBox` |  | N |  |  |
| `center` | `mojo.geom.LngLat` |  | N |  | 经纬度 |


#### `mojo.geom.BoundingBox`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `leftBottom` | `mojo.geom.LngLat` |  | N |  | 经纬度 |
| `rightTop` | `mojo.geom.LngLat` |  | N |  | 经纬度 |


#### `mojo.geom.LngLat`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `longitude` | `number` | `Float64` | Y |  | the longitude of the `LngLat`经度 |
| `latitude` | `number` | `Float64` | Y |  | the latitude of the `LngLat`维度 |
| `altitude` | `number` | `Float64` | N |  | the altitude of the `LngLat` in meters.高度 |


## 

### 请求路径
```http
PUT /tilestream/v1/layers/{layer}/tile_info
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `type` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `format` | `string` |  | N |  |
| `version` | `string` |  | N |  |
| `attribution` | `string` |  | N |  |
| `scheme` | `string` |  | N |  | tms, xyz |
| `minZoom` | `integer` | `Int32` | N |  |
| `maxZoom` | `integer` | `Int32` | N |  |
| `bounds` | `mojo.geom.BoundingBox` |  | N |  |  |
| `center` | `mojo.geom.LngLat` |  | N |  | 经纬度 |


#### `mojo.geom.BoundingBox`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `leftBottom` | `mojo.geom.LngLat` |  | N |  | 经纬度 |
| `rightTop` | `mojo.geom.LngLat` |  | N |  | 经纬度 |


#### `mojo.geom.LngLat`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `longitude` | `number` | `Float64` | Y |  | the longitude of the `LngLat`经度 |
| `latitude` | `number` | `Float64` | Y |  | the latitude of the `LngLat`维度 |
| `altitude` | `number` | `Float64` | N |  | the altitude of the `LngLat` in meters.高度 |


### 返回值

#### 返回对象
对象为空

## 

### 请求路径
```http
POST /tilestream/v1/layers/{layer}/tiles
```


### 请求参数

### 返回值

#### 返回对象
对象为空

## 

### 请求路径
```http
GET /tilestream/v1/layers/{layer}/tiles/{level}/{x}/{y}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |
| `level` | `integer` | `Int32` |  |
| `x` | `integer` | `Int32` |  |
| `y` | `integer` | `Int32` |  |
| `format` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `x` | `integer` | `Int32` | N |  |
| `y` | `integer` | `Int32` | N |  |
| `level` | `integer` | `Int32` | N |  |
| `format` | `string` |  | N |  |
| `encoding` | `string` |  | N |  |
| `content` | `string` | `Bytes` | N |  |


## 

### 请求路径
```http
PUT /tilestream/v1/layers/{layer}/tiles/{level}/{x}/{y}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |
| `level` | `integer` | `Int32` |  |
| `x` | `integer` | `Int32` |  |
| `y` | `integer` | `Int32` |  |
| `format` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `x` | `integer` | `Int32` | N |  |
| `y` | `integer` | `Int32` | N |  |
| `level` | `integer` | `Int32` | N |  |
| `format` | `string` |  | N |  |
| `encoding` | `string` |  | N |  |
| `content` | `string` | `Bytes` | N |  |


### 返回值

#### 返回对象
对象为空

## 

### 请求路径
```http
POST /tilestream/v1/layers/{layer}/tiles/{level}/{x}/{y}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |
| `level` | `integer` | `Int32` |  |
| `x` | `integer` | `Int32` |  |
| `y` | `integer` | `Int32` |  |
| `format` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `x` | `integer` | `Int32` | N |  |
| `y` | `integer` | `Int32` | N |  |
| `level` | `integer` | `Int32` | N |  |
| `format` | `string` |  | N |  |
| `encoding` | `string` |  | N |  |
| `content` | `string` | `Bytes` | N |  |


## 

### 请求路径
```http
GET /tilestream/v1/layers/{layer}/tiles/{level}/{x}/{y}.{format}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |
| `level` | `integer` | `Int32` |  |
| `x` | `integer` | `Int32` |  |
| `y` | `integer` | `Int32` |  |
| `format` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `x` | `integer` | `Int32` | N |  |
| `y` | `integer` | `Int32` | N |  |
| `level` | `integer` | `Int32` | N |  |
| `format` | `string` |  | N |  |
| `encoding` | `string` |  | N |  |
| `content` | `string` | `Bytes` | N |  |


## 

### 请求路径
```http
PUT /tilestream/v1/layers/{layer}/tiles/{level}/{x}/{y}.{format}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |
| `level` | `integer` | `Int32` |  |
| `x` | `integer` | `Int32` |  |
| `y` | `integer` | `Int32` |  |
| `format` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `x` | `integer` | `Int32` | N |  |
| `y` | `integer` | `Int32` | N |  |
| `level` | `integer` | `Int32` | N |  |
| `format` | `string` |  | N |  |
| `encoding` | `string` |  | N |  |
| `content` | `string` | `Bytes` | N |  |


### 返回值

#### 返回对象
对象为空

## 

### 请求路径
```http
POST /tilestream/v1/layers/{layer}/tiles/{level}/{x}/{y}.{format}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `layer` | `string` |  |  |
| `level` | `integer` | `Int32` |  |
| `x` | `integer` | `Int32` |  |
| `y` | `integer` | `Int32` |  |
| `format` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `x` | `integer` | `Int32` | N |  |
| `y` | `integer` | `Int32` | N |  |
| `level` | `integer` | `Int32` | N |  |
| `format` | `string` |  | N |  |
| `encoding` | `string` |  | N |  |
| `content` | `string` | `Bytes` | N |  |
