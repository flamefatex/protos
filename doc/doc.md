# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [flamefatex/common/paging.proto](#flamefatex/common/paging.proto)
    - [Paging](#flamefatex.common.Paging)
  
- [flamefatex/example/example.proto](#flamefatex/example/example.proto)
    - [ExampleType](#flamefatex.example.ExampleType)
  
- [flamefatex/example-api/v2/example.proto](#flamefatex/example-api/v2/example.proto)
    - [Example](#flamefatex.example_api.v2.Example)
    - [ExampleAllRequest](#flamefatex.example_api.v2.ExampleAllRequest)
    - [ExampleAllResponse](#flamefatex.example_api.v2.ExampleAllResponse)
    - [ExampleDeleteRequest](#flamefatex.example_api.v2.ExampleDeleteRequest)
    - [ExampleGetRequest](#flamefatex.example_api.v2.ExampleGetRequest)
    - [ExampleGetResponse](#flamefatex.example_api.v2.ExampleGetResponse)
    - [ExampleListRequest](#flamefatex.example_api.v2.ExampleListRequest)
    - [ExampleListResponse](#flamefatex.example_api.v2.ExampleListResponse)
  
    - [ExampleService](#flamefatex.example_api.v2.ExampleService)
  
- [flamefatex/example-api/v2/external/example.proto](#flamefatex/example-api/v2/external/example.proto)
    - [Example](#flamefatex.example_api.v2.external.Example)
    - [ExampleAllRequest](#flamefatex.example_api.v2.external.ExampleAllRequest)
    - [ExampleAllResponse](#flamefatex.example_api.v2.external.ExampleAllResponse)
  
    - [ExampleService](#flamefatex.example_api.v2.external.ExampleService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="flamefatex/common/paging.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## flamefatex/common/paging.proto



<a name="flamefatex.common.Paging"></a>

### Paging
分页


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| page_size | [int32](#int32) |  |  |
| total | [int32](#int32) |  |  |





 

 

 

 



<a name="flamefatex/example/example.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## flamefatex/example/example.proto


 


<a name="flamefatex.example.ExampleType"></a>

### ExampleType


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXAMPLE_TYPE_UNSPECIFIED | 0 | 未指定 |
| EXAMPLE_TYPE_ACCESS | 1 | 接入类型 |
| EXAMPLE_TYPE_PARSE | 2 | 解析类型 |


 

 

 



<a name="flamefatex/example-api/v2/example.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## flamefatex/example-api/v2/example.proto



<a name="flamefatex.example_api.v2.Example"></a>

### Example



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| type | [flamefatex.example.ExampleType](#flamefatex.example.ExampleType) |  |  |
| description | [string](#string) |  |  |
| create_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| update_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="flamefatex.example_api.v2.ExampleAllRequest"></a>

### ExampleAllRequest
All






<a name="flamefatex.example_api.v2.ExampleAllResponse"></a>

### ExampleAllResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| examples | [Example](#flamefatex.example_api.v2.Example) | repeated |  |






<a name="flamefatex.example_api.v2.ExampleDeleteRequest"></a>

### ExampleDeleteRequest
Delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="flamefatex.example_api.v2.ExampleGetRequest"></a>

### ExampleGetRequest
Get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="flamefatex.example_api.v2.ExampleGetResponse"></a>

### ExampleGetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| example | [Example](#flamefatex.example_api.v2.Example) |  |  |






<a name="flamefatex.example_api.v2.ExampleListRequest"></a>

### ExampleListRequest
List


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| page_size | [int32](#int32) |  |  |
| type | [flamefatex.example.ExampleType](#flamefatex.example.ExampleType) |  |  |






<a name="flamefatex.example_api.v2.ExampleListResponse"></a>

### ExampleListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paging | [flamefatex.common.Paging](#flamefatex.common.Paging) |  |  |
| examples | [Example](#flamefatex.example_api.v2.Example) | repeated |  |





 

 

 


<a name="flamefatex.example_api.v2.ExampleService"></a>

### ExampleService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| All | [ExampleAllRequest](#flamefatex.example_api.v2.ExampleAllRequest) | [ExampleAllResponse](#flamefatex.example_api.v2.ExampleAllResponse) |  |
| List | [ExampleListRequest](#flamefatex.example_api.v2.ExampleListRequest) | [ExampleListResponse](#flamefatex.example_api.v2.ExampleListResponse) |  |
| Get | [ExampleGetRequest](#flamefatex.example_api.v2.ExampleGetRequest) | [ExampleGetResponse](#flamefatex.example_api.v2.ExampleGetResponse) |  |
| Create | [Example](#flamefatex.example_api.v2.Example) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| Update | [Example](#flamefatex.example_api.v2.Example) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| Delete | [ExampleDeleteRequest](#flamefatex.example_api.v2.ExampleDeleteRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |

 



<a name="flamefatex/example-api/v2/external/example.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## flamefatex/example-api/v2/external/example.proto



<a name="flamefatex.example_api.v2.external.Example"></a>

### Example



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| type | [flamefatex.example.ExampleType](#flamefatex.example.ExampleType) |  |  |
| description | [string](#string) |  |  |
| create_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| update_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="flamefatex.example_api.v2.external.ExampleAllRequest"></a>

### ExampleAllRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [flamefatex.example.ExampleType](#flamefatex.example.ExampleType) |  |  |






<a name="flamefatex.example_api.v2.external.ExampleAllResponse"></a>

### ExampleAllResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| examples | [Example](#flamefatex.example_api.v2.external.Example) | repeated |  |





 

 

 


<a name="flamefatex.example_api.v2.external.ExampleService"></a>

### ExampleService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| All | [ExampleAllRequest](#flamefatex.example_api.v2.external.ExampleAllRequest) | [ExampleAllResponse](#flamefatex.example_api.v2.external.ExampleAllResponse) |  |

 



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

