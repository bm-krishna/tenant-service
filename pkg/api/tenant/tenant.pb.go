// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: api/tenant.proto

package tenant

import (
	proto "github.com/golang/protobuf/proto"
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

type TenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request []byte `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *TenantRequest) Reset() {
	*x = TenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantRequest) ProtoMessage() {}

func (x *TenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantRequest.ProtoReflect.Descriptor instead.
func (*TenantRequest) Descriptor() ([]byte, []int) {
	return file_api_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *TenantRequest) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

type TenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *TenantResponse) Reset() {
	*x = TenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResponse) ProtoMessage() {}

func (x *TenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResponse.ProtoReflect.Descriptor instead.
func (*TenantResponse) Descriptor() ([]byte, []int) {
	return file_api_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *TenantResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_api_tenant_proto protoreflect.FileDescriptor

var file_api_tenant_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x29, 0x0a, 0x0d, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x77, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x34, 0x0a,
	0x03, 0x41, 0x50, 0x49, 0x12, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6d, 0x2d, 0x6b, 0x72,
	0x69, 0x73, 0x68, 0x6e, 0x61, 0x2d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_tenant_proto_rawDescOnce sync.Once
	file_api_tenant_proto_rawDescData = file_api_tenant_proto_rawDesc
)

func file_api_tenant_proto_rawDescGZIP() []byte {
	file_api_tenant_proto_rawDescOnce.Do(func() {
		file_api_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_tenant_proto_rawDescData)
	})
	return file_api_tenant_proto_rawDescData
}

var file_api_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_tenant_proto_goTypes = []interface{}{
	(*TenantRequest)(nil),  // 0: tenant.TenantRequest
	(*TenantResponse)(nil), // 1: tenant.TenantResponse
}
var file_api_tenant_proto_depIdxs = []int32{
	0, // 0: tenant.Tenant.API:input_type -> tenant.TenantRequest
	0, // 1: tenant.Tenant.Create:input_type -> tenant.TenantRequest
	1, // 2: tenant.Tenant.API:output_type -> tenant.TenantResponse
	1, // 3: tenant.Tenant.Create:output_type -> tenant.TenantResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_tenant_proto_init() }
func file_api_tenant_proto_init() {
	if File_api_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantRequest); i {
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
		file_api_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResponse); i {
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
			RawDescriptor: file_api_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_tenant_proto_goTypes,
		DependencyIndexes: file_api_tenant_proto_depIdxs,
		MessageInfos:      file_api_tenant_proto_msgTypes,
	}.Build()
	File_api_tenant_proto = out.File
	file_api_tenant_proto_rawDesc = nil
	file_api_tenant_proto_goTypes = nil
	file_api_tenant_proto_depIdxs = nil
}
