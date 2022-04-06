// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: protos/worker.proto

package nemon

import (
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

type GetAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key int64 `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetAppsRequest) Reset() {
	*x = GetAppsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppsRequest) ProtoMessage() {}

func (x *GetAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppsRequest.ProtoReflect.Descriptor instead.
func (*GetAppsRequest) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{0}
}

func (x *GetAppsRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

type GetAppsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications []*GetAppsResponse_ApplicationInfo `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
	Username     string                             `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetAppsResponse) Reset() {
	*x = GetAppsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppsResponse) ProtoMessage() {}

func (x *GetAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppsResponse.ProtoReflect.Descriptor instead.
func (*GetAppsResponse) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{1}
}

func (x *GetAppsResponse) GetApplications() []*GetAppsResponse_ApplicationInfo {
	if x != nil {
		return x.Applications
	}
	return nil
}

func (x *GetAppsResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type DeleteAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key  int64  `protobuf:"varint,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteAppsRequest) Reset() {
	*x = DeleteAppsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppsRequest) ProtoMessage() {}

func (x *DeleteAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppsRequest.ProtoReflect.Descriptor instead.
func (*DeleteAppsRequest) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAppsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteAppsRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

type DeleteAppsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *DeleteAppsResponse) Reset() {
	*x = DeleteAppsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppsResponse) ProtoMessage() {}

func (x *DeleteAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppsResponse.ProtoReflect.Descriptor instead.
func (*DeleteAppsResponse) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAppsResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type GetAppsResponse_ApplicationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *GetAppsResponse_ApplicationInfo) Reset() {
	*x = GetAppsResponse_ApplicationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppsResponse_ApplicationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppsResponse_ApplicationInfo) ProtoMessage() {}

func (x *GetAppsResponse_ApplicationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppsResponse_ApplicationInfo.ProtoReflect.Descriptor instead.
func (*GetAppsResponse_ApplicationInfo) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetAppsResponse_ApplicationInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAppsResponse_ApplicationInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_protos_worker_proto protoreflect.FileDescriptor

var file_protos_worker_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xb6, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a,
	0x41, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x24, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x32, 0x6c, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x6d, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_worker_proto_rawDescOnce sync.Once
	file_protos_worker_proto_rawDescData = file_protos_worker_proto_rawDesc
)

func file_protos_worker_proto_rawDescGZIP() []byte {
	file_protos_worker_proto_rawDescOnce.Do(func() {
		file_protos_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_worker_proto_rawDescData)
	})
	return file_protos_worker_proto_rawDescData
}

var file_protos_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_worker_proto_goTypes = []interface{}{
	(*GetAppsRequest)(nil),                  // 0: GetAppsRequest
	(*GetAppsResponse)(nil),                 // 1: GetAppsResponse
	(*DeleteAppsRequest)(nil),               // 2: DeleteAppsRequest
	(*DeleteAppsResponse)(nil),              // 3: DeleteAppsResponse
	(*GetAppsResponse_ApplicationInfo)(nil), // 4: GetAppsResponse.ApplicationInfo
}
var file_protos_worker_proto_depIdxs = []int32{
	4, // 0: GetAppsResponse.applications:type_name -> GetAppsResponse.ApplicationInfo
	0, // 1: Worker.GetApps:input_type -> GetAppsRequest
	2, // 2: Worker.DeleteApp:input_type -> DeleteAppsRequest
	1, // 3: Worker.GetApps:output_type -> GetAppsResponse
	3, // 4: Worker.DeleteApp:output_type -> DeleteAppsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_worker_proto_init() }
func file_protos_worker_proto_init() {
	if File_protos_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppsRequest); i {
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
		file_protos_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppsResponse); i {
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
		file_protos_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppsRequest); i {
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
		file_protos_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppsResponse); i {
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
		file_protos_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppsResponse_ApplicationInfo); i {
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
			RawDescriptor: file_protos_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_worker_proto_goTypes,
		DependencyIndexes: file_protos_worker_proto_depIdxs,
		MessageInfos:      file_protos_worker_proto_msgTypes,
	}.Build()
	File_protos_worker_proto = out.File
	file_protos_worker_proto_rawDesc = nil
	file_protos_worker_proto_goTypes = nil
	file_protos_worker_proto_depIdxs = nil
}
