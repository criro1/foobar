// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.3
// source: foo.proto

package foo

import (
	context "context"
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

type BarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to BarRequests:
	//	*BarRequest_First
	//	*BarRequest_Second
	//	*BarRequest_Third
	BarRequests isBarRequest_BarRequests `protobuf_oneof:"bar_requests"`
}

func (x *BarRequest) Reset() {
	*x = BarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarRequest) ProtoMessage() {}

func (x *BarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarRequest.ProtoReflect.Descriptor instead.
func (*BarRequest) Descriptor() ([]byte, []int) {
	return file_foo_proto_rawDescGZIP(), []int{0}
}

func (m *BarRequest) GetBarRequests() isBarRequest_BarRequests {
	if m != nil {
		return m.BarRequests
	}
	return nil
}

func (x *BarRequest) GetFirst() *BarRequestFirst {
	if x, ok := x.GetBarRequests().(*BarRequest_First); ok {
		return x.First
	}
	return nil
}

func (x *BarRequest) GetSecond() *BarRequestSecond {
	if x, ok := x.GetBarRequests().(*BarRequest_Second); ok {
		return x.Second
	}
	return nil
}

func (x *BarRequest) GetThird() string {
	if x, ok := x.GetBarRequests().(*BarRequest_Third); ok {
		return x.Third
	}
	return ""
}

type isBarRequest_BarRequests interface {
	isBarRequest_BarRequests()
}

type BarRequest_First struct {
	First *BarRequestFirst `protobuf:"bytes,1,opt,name=first,proto3,oneof"`
}

type BarRequest_Second struct {
	Second *BarRequestSecond `protobuf:"bytes,2,opt,name=second,proto3,oneof"`
}

type BarRequest_Third struct {
	Third string `protobuf:"bytes,3,opt,name=third,proto3,oneof"`
}

func (*BarRequest_First) isBarRequest_BarRequests() {}

func (*BarRequest_Second) isBarRequest_BarRequests() {}

func (*BarRequest_Third) isBarRequest_BarRequests() {}

type BarRequestFirst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *BarRequestFirst) Reset() {
	*x = BarRequestFirst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarRequestFirst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarRequestFirst) ProtoMessage() {}

func (x *BarRequestFirst) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarRequestFirst.ProtoReflect.Descriptor instead.
func (*BarRequestFirst) Descriptor() ([]byte, []int) {
	return file_foo_proto_rawDescGZIP(), []int{1}
}

func (x *BarRequestFirst) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

type BarRequestSecond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num []int64 `protobuf:"varint,1,rep,packed,name=num,proto3" json:"num,omitempty"`
}

func (x *BarRequestSecond) Reset() {
	*x = BarRequestSecond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarRequestSecond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarRequestSecond) ProtoMessage() {}

func (x *BarRequestSecond) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarRequestSecond.ProtoReflect.Descriptor instead.
func (*BarRequestSecond) Descriptor() ([]byte, []int) {
	return file_foo_proto_rawDescGZIP(), []int{2}
}

func (x *BarRequestSecond) GetNum() []int64 {
	if x != nil {
		return x.Num
	}
	return nil
}

type BarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *BarResponse) Reset() {
	*x = BarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarResponse) ProtoMessage() {}

func (x *BarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarResponse.ProtoReflect.Descriptor instead.
func (*BarResponse) Descriptor() ([]byte, []int) {
	return file_foo_proto_rawDescGZIP(), []int{3}
}

func (x *BarResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_foo_proto protoreflect.FileDescriptor

var file_foo_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x66, 0x6f, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93,
	0x01, 0x0a, 0x0a, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x66,
	0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x6f,
	0x6f, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x05,
	0x74, 0x68, 0x69, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x62, 0x61, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x10, 0x42, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x21, 0x0a,
	0x0b, 0x42, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x32, 0x40, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x39, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x0f,
	0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22, 0x04, 0x2f, 0x62, 0x61, 0x72, 0x3a,
	0x01, 0x2a, 0x42, 0x06, 0x5a, 0x04, 0x3b, 0x66, 0x6f, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_foo_proto_rawDescOnce sync.Once
	file_foo_proto_rawDescData = file_foo_proto_rawDesc
)

func file_foo_proto_rawDescGZIP() []byte {
	file_foo_proto_rawDescOnce.Do(func() {
		file_foo_proto_rawDescData = protoimpl.X.CompressGZIP(file_foo_proto_rawDescData)
	})
	return file_foo_proto_rawDescData
}

var file_foo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_foo_proto_goTypes = []interface{}{
	(*BarRequest)(nil),       // 0: foo.BarRequest
	(*BarRequestFirst)(nil),  // 1: foo.BarRequestFirst
	(*BarRequestSecond)(nil), // 2: foo.BarRequestSecond
	(*BarResponse)(nil),      // 3: foo.BarResponse
}
var file_foo_proto_depIdxs = []int32{
	1, // 0: foo.BarRequest.first:type_name -> foo.BarRequestFirst
	2, // 1: foo.BarRequest.second:type_name -> foo.BarRequestSecond
	0, // 2: foo.Foo.Bar:input_type -> foo.BarRequest
	3, // 3: foo.Foo.Bar:output_type -> foo.BarResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foo_proto_init() }
func file_foo_proto_init() {
	if File_foo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarRequest); i {
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
		file_foo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarRequestFirst); i {
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
		file_foo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarRequestSecond); i {
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
		file_foo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarResponse); i {
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
	file_foo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BarRequest_First)(nil),
		(*BarRequest_Second)(nil),
		(*BarRequest_Third)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_foo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foo_proto_goTypes,
		DependencyIndexes: file_foo_proto_depIdxs,
		MessageInfos:      file_foo_proto_msgTypes,
	}.Build()
	File_foo_proto = out.File
	file_foo_proto_rawDesc = nil
	file_foo_proto_goTypes = nil
	file_foo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FooClient is the client API for Foo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FooClient interface {
	Bar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error)
}

type fooClient struct {
	cc grpc.ClientConnInterface
}

func NewFooClient(cc grpc.ClientConnInterface) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) Bar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error) {
	out := new(BarResponse)
	err := c.cc.Invoke(ctx, "/foo.Foo/Bar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServer is the server API for Foo service.
type FooServer interface {
	Bar(context.Context, *BarRequest) (*BarResponse, error)
}

// UnimplementedFooServer can be embedded to have forward compatible implementations.
type UnimplementedFooServer struct {
}

func (*UnimplementedFooServer) Bar(context.Context, *BarRequest) (*BarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bar not implemented")
}

func RegisterFooServer(s *grpc.Server, srv FooServer) {
	s.RegisterService(&_Foo_serviceDesc, srv)
}

func _Foo_Bar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServer).Bar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.Foo/Bar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServer).Bar(ctx, req.(*BarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Foo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foo.Foo",
	HandlerType: (*FooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bar",
			Handler:    _Foo_Bar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foo.proto",
}