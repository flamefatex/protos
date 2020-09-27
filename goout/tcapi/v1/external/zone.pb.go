// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: tcapi/v1/external/zone.proto

package external

import (
	context "context"
	tc "github.com/flamefatex/protos/goout/tc"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ZoneListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type tc.ZoneType `protobuf:"varint,1,opt,name=type,proto3,enum=flamefatex.tc.ZoneType" json:"type,omitempty"`
}

func (x *ZoneListRequest) Reset() {
	*x = ZoneListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcapi_v1_external_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneListRequest) ProtoMessage() {}

func (x *ZoneListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tcapi_v1_external_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneListRequest.ProtoReflect.Descriptor instead.
func (*ZoneListRequest) Descriptor() ([]byte, []int) {
	return file_tcapi_v1_external_zone_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneListRequest) GetType() tc.ZoneType {
	if x != nil {
		return x.Type
	}
	return tc.ZoneType_ZONE_TYPE_UNSPECIFIED
}

type ZoneListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ZoneListResponseData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ZoneListResponse) Reset() {
	*x = ZoneListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcapi_v1_external_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneListResponse) ProtoMessage() {}

func (x *ZoneListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tcapi_v1_external_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneListResponse.ProtoReflect.Descriptor instead.
func (*ZoneListResponse) Descriptor() ([]byte, []int) {
	return file_tcapi_v1_external_zone_proto_rawDescGZIP(), []int{1}
}

func (x *ZoneListResponse) GetData() []*ZoneListResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ZoneListResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type        tc.ZoneType `protobuf:"varint,3,opt,name=type,proto3,enum=flamefatex.tc.ZoneType" json:"type,omitempty"`
	Description string      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ZoneListResponseData) Reset() {
	*x = ZoneListResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcapi_v1_external_zone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneListResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneListResponseData) ProtoMessage() {}

func (x *ZoneListResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_tcapi_v1_external_zone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneListResponseData.ProtoReflect.Descriptor instead.
func (*ZoneListResponseData) Descriptor() ([]byte, []int) {
	return file_tcapi_v1_external_zone_proto_rawDescGZIP(), []int{2}
}

func (x *ZoneListResponseData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ZoneListResponseData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ZoneListResponseData) GetType() tc.ZoneType {
	if x != nil {
		return x.Type
	}
	return tc.ZoneType_ZONE_TYPE_UNSPECIFIED
}

func (x *ZoneListResponseData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_tcapi_v1_external_zone_proto protoreflect.FileDescriptor

var file_tcapi_v1_external_zone_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x63, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x66, 0x6c, 0x61, 0x6d, 0x65, 0x66, 0x61, 0x74, 0x65, 0x78, 0x2e, 0x74, 0x63, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x0d, 0x74, 0x63,
	0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0f, 0x5a, 0x6f, 0x6e,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x66, 0x6c, 0x61,
	0x6d, 0x65, 0x66, 0x61, 0x74, 0x65, 0x78, 0x2e, 0x74, 0x63, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5a, 0x0a, 0x10, 0x5a, 0x6f, 0x6e,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x66, 0x6c,
	0x61, 0x6d, 0x65, 0x66, 0x61, 0x74, 0x65, 0x78, 0x2e, 0x74, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x89, 0x01, 0x0a, 0x14, 0x5a, 0x6f, 0x6e, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x66, 0x6c, 0x61, 0x6d, 0x65, 0x66, 0x61, 0x74, 0x65, 0x78, 0x2e, 0x74, 0x63,
	0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0x95, 0x01, 0x0a, 0x0b, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x85, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x66, 0x6c, 0x61,
	0x6d, 0x65, 0x66, 0x61, 0x74, 0x65, 0x78, 0x2e, 0x74, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x66, 0x6c, 0x61, 0x6d,
	0x65, 0x66, 0x61, 0x74, 0x65, 0x78, 0x2e, 0x74, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x61, 0x6d, 0x65, 0x66, 0x61, 0x74,
	0x65, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x75, 0x74, 0x2f,
	0x74, 0x63, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tcapi_v1_external_zone_proto_rawDescOnce sync.Once
	file_tcapi_v1_external_zone_proto_rawDescData = file_tcapi_v1_external_zone_proto_rawDesc
)

func file_tcapi_v1_external_zone_proto_rawDescGZIP() []byte {
	file_tcapi_v1_external_zone_proto_rawDescOnce.Do(func() {
		file_tcapi_v1_external_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_tcapi_v1_external_zone_proto_rawDescData)
	})
	return file_tcapi_v1_external_zone_proto_rawDescData
}

var file_tcapi_v1_external_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tcapi_v1_external_zone_proto_goTypes = []interface{}{
	(*ZoneListRequest)(nil),      // 0: flamefatex.tcapi.v1.external.ZoneListRequest
	(*ZoneListResponse)(nil),     // 1: flamefatex.tcapi.v1.external.ZoneListResponse
	(*ZoneListResponseData)(nil), // 2: flamefatex.tcapi.v1.external.ZoneListResponseData
	(tc.ZoneType)(0),             // 3: flamefatex.tc.ZoneType
}
var file_tcapi_v1_external_zone_proto_depIdxs = []int32{
	3, // 0: flamefatex.tcapi.v1.external.ZoneListRequest.type:type_name -> flamefatex.tc.ZoneType
	2, // 1: flamefatex.tcapi.v1.external.ZoneListResponse.data:type_name -> flamefatex.tcapi.v1.external.ZoneListResponseData
	3, // 2: flamefatex.tcapi.v1.external.ZoneListResponseData.type:type_name -> flamefatex.tc.ZoneType
	0, // 3: flamefatex.tcapi.v1.external.ZoneService.List:input_type -> flamefatex.tcapi.v1.external.ZoneListRequest
	1, // 4: flamefatex.tcapi.v1.external.ZoneService.List:output_type -> flamefatex.tcapi.v1.external.ZoneListResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tcapi_v1_external_zone_proto_init() }
func file_tcapi_v1_external_zone_proto_init() {
	if File_tcapi_v1_external_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tcapi_v1_external_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tcapi_v1_external_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tcapi_v1_external_zone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneListResponseData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tcapi_v1_external_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tcapi_v1_external_zone_proto_goTypes,
		DependencyIndexes: file_tcapi_v1_external_zone_proto_depIdxs,
		MessageInfos:      file_tcapi_v1_external_zone_proto_msgTypes,
	}.Build()
	File_tcapi_v1_external_zone_proto = out.File
	file_tcapi_v1_external_zone_proto_rawDesc = nil
	file_tcapi_v1_external_zone_proto_goTypes = nil
	file_tcapi_v1_external_zone_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ZoneServiceClient is the client API for ZoneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ZoneServiceClient interface {
	List(ctx context.Context, in *ZoneListRequest, opts ...grpc.CallOption) (*ZoneListResponse, error)
}

type zoneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZoneServiceClient(cc grpc.ClientConnInterface) ZoneServiceClient {
	return &zoneServiceClient{cc}
}

func (c *zoneServiceClient) List(ctx context.Context, in *ZoneListRequest, opts ...grpc.CallOption) (*ZoneListResponse, error) {
	out := new(ZoneListResponse)
	err := c.cc.Invoke(ctx, "/flamefatex.tcapi.v1.external.ZoneService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZoneServiceServer is the server API for ZoneService service.
type ZoneServiceServer interface {
	List(context.Context, *ZoneListRequest) (*ZoneListResponse, error)
}

// UnimplementedZoneServiceServer can be embedded to have forward compatible implementations.
type UnimplementedZoneServiceServer struct {
}

func (*UnimplementedZoneServiceServer) List(context.Context, *ZoneListRequest) (*ZoneListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterZoneServiceServer(s *grpc.Server, srv ZoneServiceServer) {
	s.RegisterService(&_ZoneService_serviceDesc, srv)
}

func _ZoneService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flamefatex.tcapi.v1.external.ZoneService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).List(ctx, req.(*ZoneListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ZoneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flamefatex.tcapi.v1.external.ZoneService",
	HandlerType: (*ZoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ZoneService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tcapi/v1/external/zone.proto",
}
