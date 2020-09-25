# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [tc/zone.proto](#tc/zone.proto)
    - [Zone](#flamefatex.tc.Zone)
  
    - [ZoneType](#flamefatex.tc.ZoneType)
  
- [tcapi/v1/zone.proto](#tcapi/v1/zone.proto)
    - [Zone](#flamefatex.tcapi.v1.Zone)
    - [ZoneAllRequest](#flamefatex.tcapi.v1.ZoneAllRequest)
    - [ZoneAllResponse](#flamefatex.tcapi.v1.ZoneAllResponse)
  
    - [ZoneService](#flamefatex.tcapi.v1.ZoneService)
  
- [tcapi/v1/external/zone.proto](#tcapi/v1/external/zone.proto)
    - [ZoneListRequest](#flamefatex.tcapi.v1.external.ZoneListRequest)
    - [ZoneListResponse](#flamefatex.tcapi.v1.external.ZoneListResponse)
    - [ZoneListResponseData](#flamefatex.tcapi.v1.external.ZoneListResponseData)
  
    - [ZoneService](#flamefatex.tcapi.v1.external.ZoneService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="tc/zone.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tc/zone.proto



<a name="flamefatex.tc.Zone"></a>

### Zone



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| type | [ZoneType](#flamefatex.tc.ZoneType) |  |  |
| create_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| update_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 


<a name="flamefatex.tc.ZoneType"></a>

### ZoneType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ZONE_TYPE_UNSPECIFIED | 0 |  |
| ZONE_TYPE_ACCESS | 1 |  |
| ZONE_TYPE_PARSE | 2 |  |


 

 

 



<a name="tcapi/v1/zone.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tcapi/v1/zone.proto



<a name="flamefatex.tcapi.v1.Zone"></a>

### Zone



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| type | [flamefatex.tc.ZoneType](#flamefatex.tc.ZoneType) |  |  |
| description | [string](#string) |  |  |






<a name="flamefatex.tcapi.v1.ZoneAllRequest"></a>

### ZoneAllRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [flamefatex.tc.ZoneType](#flamefatex.tc.ZoneType) |  |  |






<a name="flamefatex.tcapi.v1.ZoneAllResponse"></a>

### ZoneAllResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Zone](#flamefatex.tcapi.v1.Zone) | repeated |  |





 

 

 


<a name="flamefatex.tcapi.v1.ZoneService"></a>

### ZoneService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| All | [ZoneAllRequest](#flamefatex.tcapi.v1.ZoneAllRequest) | [ZoneAllResponse](#flamefatex.tcapi.v1.ZoneAllResponse) |  |

 



<a name="tcapi/v1/external/zone.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tcapi/v1/external/zone.proto



<a name="flamefatex.tcapi.v1.external.ZoneListRequest"></a>

### ZoneListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [flamefatex.tc.ZoneType](#flamefatex.tc.ZoneType) |  |  |






<a name="flamefatex.tcapi.v1.external.ZoneListResponse"></a>

### ZoneListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [ZoneListResponseData](#flamefatex.tcapi.v1.external.ZoneListResponseData) | repeated |  |






<a name="flamefatex.tcapi.v1.external.ZoneListResponseData"></a>

### ZoneListResponseData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| type | [flamefatex.tc.ZoneType](#flamefatex.tc.ZoneType) |  |  |
| description | [string](#string) |  |  |





 

 

 


<a name="flamefatex.tcapi.v1.external.ZoneService"></a>

### ZoneService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ZoneListRequest](#flamefatex.tcapi.v1.external.ZoneListRequest) | [ZoneListResponse](#flamefatex.tcapi.v1.external.ZoneListResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

