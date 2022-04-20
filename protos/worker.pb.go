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

type ExchangeEphemeralPublicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A []byte `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
}

func (x *ExchangeEphemeralPublicRequest) Reset() {
	*x = ExchangeEphemeralPublicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeEphemeralPublicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeEphemeralPublicRequest) ProtoMessage() {}

func (x *ExchangeEphemeralPublicRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ExchangeEphemeralPublicRequest.ProtoReflect.Descriptor instead.
func (*ExchangeEphemeralPublicRequest) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeEphemeralPublicRequest) GetA() []byte {
	if x != nil {
		return x.A
	}
	return nil
}

type ExchangeEphemeralPublicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B []byte `protobuf:"bytes,1,opt,name=B,proto3" json:"B,omitempty"`
}

func (x *ExchangeEphemeralPublicResponse) Reset() {
	*x = ExchangeEphemeralPublicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeEphemeralPublicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeEphemeralPublicResponse) ProtoMessage() {}

func (x *ExchangeEphemeralPublicResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ExchangeEphemeralPublicResponse.ProtoReflect.Descriptor instead.
func (*ExchangeEphemeralPublicResponse) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeEphemeralPublicResponse) GetB() []byte {
	if x != nil {
		return x.B
	}
	return nil
}

type IsEnrolledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key int64 `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *IsEnrolledRequest) Reset() {
	*x = IsEnrolledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsEnrolledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsEnrolledRequest) ProtoMessage() {}

func (x *IsEnrolledRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IsEnrolledRequest.ProtoReflect.Descriptor instead.
func (*IsEnrolledRequest) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{2}
}

func (x *IsEnrolledRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

type IsEnrolledResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enrolled bool `protobuf:"varint,1,opt,name=enrolled,proto3" json:"enrolled,omitempty"` // check if is enrolled after above check
}

func (x *IsEnrolledResponse) Reset() {
	*x = IsEnrolledResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsEnrolledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsEnrolledResponse) ProtoMessage() {}

func (x *IsEnrolledResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IsEnrolledResponse.ProtoReflect.Descriptor instead.
func (*IsEnrolledResponse) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{3}
}

func (x *IsEnrolledResponse) GetEnrolled() bool {
	if x != nil {
		return x.Enrolled
	}
	return false
}

type SaveEnrollmentInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verifier []byte `protobuf:"bytes,1,opt,name=verifier,proto3" json:"verifier,omitempty"`
	Salt     []byte `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
	SRPGroup int64  `protobuf:"varint,3,opt,name=SRPGroup,proto3" json:"SRPGroup,omitempty"`
}

func (x *SaveEnrollmentInfoRequest) Reset() {
	*x = SaveEnrollmentInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveEnrollmentInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveEnrollmentInfoRequest) ProtoMessage() {}

func (x *SaveEnrollmentInfoRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SaveEnrollmentInfoRequest.ProtoReflect.Descriptor instead.
func (*SaveEnrollmentInfoRequest) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{4}
}

func (x *SaveEnrollmentInfoRequest) GetVerifier() []byte {
	if x != nil {
		return x.Verifier
	}
	return nil
}

func (x *SaveEnrollmentInfoRequest) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *SaveEnrollmentInfoRequest) GetSRPGroup() int64 {
	if x != nil {
		return x.SRPGroup
	}
	return 0
}

type SaveEnrollmentInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveEnrollmentInfoResponse) Reset() {
	*x = SaveEnrollmentInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveEnrollmentInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveEnrollmentInfoResponse) ProtoMessage() {}

func (x *SaveEnrollmentInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveEnrollmentInfoResponse.ProtoReflect.Descriptor instead.
func (*SaveEnrollmentInfoResponse) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{5}
}

type GetSaltAndSRPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSaltAndSRPRequest) Reset() {
	*x = GetSaltAndSRPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSaltAndSRPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSaltAndSRPRequest) ProtoMessage() {}

func (x *GetSaltAndSRPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSaltAndSRPRequest.ProtoReflect.Descriptor instead.
func (*GetSaltAndSRPRequest) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{6}
}

type GetSaltAndSRPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Salt     []byte `protobuf:"bytes,1,opt,name=salt,proto3" json:"salt,omitempty"`
	SRPGroup int64  `protobuf:"varint,2,opt,name=SRPGroup,proto3" json:"SRPGroup,omitempty"`
}

func (x *GetSaltAndSRPResponse) Reset() {
	*x = GetSaltAndSRPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSaltAndSRPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSaltAndSRPResponse) ProtoMessage() {}

func (x *GetSaltAndSRPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSaltAndSRPResponse.ProtoReflect.Descriptor instead.
func (*GetSaltAndSRPResponse) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{7}
}

func (x *GetSaltAndSRPResponse) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *GetSaltAndSRPResponse) GetSRPGroup() int64 {
	if x != nil {
		return x.SRPGroup
	}
	return 0
}

type GetSysInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        int64  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	PublicKeyN string `protobuf:"bytes,2,opt,name=publicKeyN,proto3" json:"publicKeyN,omitempty"`
	PublicKeyE string `protobuf:"bytes,3,opt,name=publicKeyE,proto3" json:"publicKeyE,omitempty"`
}

func (x *GetSysInfoRequest) Reset() {
	*x = GetSysInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysInfoRequest) ProtoMessage() {}

func (x *GetSysInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysInfoRequest.ProtoReflect.Descriptor instead.
func (*GetSysInfoRequest) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{8}
}

func (x *GetSysInfoRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *GetSysInfoRequest) GetPublicKeyN() string {
	if x != nil {
		return x.PublicKeyN
	}
	return ""
}

func (x *GetSysInfoRequest) GetPublicKeyE() string {
	if x != nil {
		return x.PublicKeyE
	}
	return ""
}

type GetSysInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerSysInfo *GetSysInfoResponse_SysInfo `protobuf:"bytes,1,opt,name=workerSysInfo,proto3" json:"workerSysInfo,omitempty"`
	AESKey        []byte                      `protobuf:"bytes,2,opt,name=AESKey,proto3" json:"AESKey,omitempty"`
}

func (x *GetSysInfoResponse) Reset() {
	*x = GetSysInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysInfoResponse) ProtoMessage() {}

func (x *GetSysInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysInfoResponse.ProtoReflect.Descriptor instead.
func (*GetSysInfoResponse) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{9}
}

func (x *GetSysInfoResponse) GetWorkerSysInfo() *GetSysInfoResponse_SysInfo {
	if x != nil {
		return x.WorkerSysInfo
	}
	return nil
}

func (x *GetSysInfoResponse) GetAESKey() []byte {
	if x != nil {
		return x.AESKey
	}
	return nil
}

type GetAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAppsRequest) Reset() {
	*x = GetAppsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppsRequest) ProtoMessage() {}

func (x *GetAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[10]
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
	return file_protos_worker_proto_rawDescGZIP(), []int{10}
}

type GetAppsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications []*GetAppsResponse_ApplicationInfo `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
	Username     []byte                             `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetAppsResponse) Reset() {
	*x = GetAppsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppsResponse) ProtoMessage() {}

func (x *GetAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[11]
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
	return file_protos_worker_proto_rawDescGZIP(), []int{11}
}

func (x *GetAppsResponse) GetApplications() []*GetAppsResponse_ApplicationInfo {
	if x != nil {
		return x.Applications
	}
	return nil
}

func (x *GetAppsResponse) GetUsername() []byte {
	if x != nil {
		return x.Username
	}
	return nil
}

type DeleteAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location []byte `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *DeleteAppsRequest) Reset() {
	*x = DeleteAppsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppsRequest) ProtoMessage() {}

func (x *DeleteAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[12]
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
	return file_protos_worker_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteAppsRequest) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *DeleteAppsRequest) GetLocation() []byte {
	if x != nil {
		return x.Location
	}
	return nil
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
		mi := &file_protos_worker_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppsResponse) ProtoMessage() {}

func (x *DeleteAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[13]
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
	return file_protos_worker_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteAppsResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type GetSysInfoResponse_SysInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username []byte `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Hostname []byte `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Os       []byte `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
}

func (x *GetSysInfoResponse_SysInfo) Reset() {
	*x = GetSysInfoResponse_SysInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysInfoResponse_SysInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysInfoResponse_SysInfo) ProtoMessage() {}

func (x *GetSysInfoResponse_SysInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysInfoResponse_SysInfo.ProtoReflect.Descriptor instead.
func (*GetSysInfoResponse_SysInfo) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{9, 0}
}

func (x *GetSysInfoResponse_SysInfo) GetUsername() []byte {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *GetSysInfoResponse_SysInfo) GetHostname() []byte {
	if x != nil {
		return x.Hostname
	}
	return nil
}

func (x *GetSysInfoResponse_SysInfo) GetOs() []byte {
	if x != nil {
		return x.Os
	}
	return nil
}

type GetAppsResponse_ApplicationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location []byte `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *GetAppsResponse_ApplicationInfo) Reset() {
	*x = GetAppsResponse_ApplicationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppsResponse_ApplicationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppsResponse_ApplicationInfo) ProtoMessage() {}

func (x *GetAppsResponse_ApplicationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[15]
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
	return file_protos_worker_proto_rawDescGZIP(), []int{11, 0}
}

func (x *GetAppsResponse_ApplicationInfo) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *GetAppsResponse_ApplicationInfo) GetLocation() []byte {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_protos_worker_proto protoreflect.FileDescriptor

var file_protos_worker_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x1e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x41, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x41, 0x22, 0x2f, 0x0a, 0x1f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x42, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x01, 0x42, 0x22, 0x25, 0x0a, 0x11, 0x49, 0x73, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x30, 0x0a,
	0x12, 0x49, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x22,
	0x67, 0x0a, 0x19, 0x53, 0x61, 0x76, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x52, 0x50, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x53, 0x52, 0x50, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x1c, 0x0a, 0x1a, 0x53, 0x61, 0x76, 0x65,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c,
	0x74, 0x41, 0x6e, 0x64, 0x53, 0x52, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x74, 0x41, 0x6e, 0x64, 0x53, 0x52, 0x50, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x52, 0x50, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53,
	0x52, 0x50, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x65, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x79,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x4e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x4e, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x45, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x45, 0x22, 0xc2,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53,
	0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x45, 0x53, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x41, 0x45, 0x53, 0x4b, 0x65, 0x79,
	0x1a, 0x51, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x6f, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x41, 0x0a, 0x0f, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x43,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xc7, 0x03, 0x0a, 0x06, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73, 0x12,
	0x0f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x0a, 0x49, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x12, 0x2e,
	0x49, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x49, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x45, 0x6e,
	0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x74,
	0x41, 0x6e, 0x64, 0x53, 0x52, 0x50, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x74,
	0x41, 0x6e, 0x64, 0x53, 0x52, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x74, 0x41, 0x6e, 0x64, 0x53, 0x52, 0x50, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x17, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x12, 0x1f, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x70, 0x68, 0x65, 0x6d,
	0x65, 0x72, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x70, 0x68, 0x65,
	0x6d, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x6d,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_protos_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_protos_worker_proto_goTypes = []interface{}{
	(*ExchangeEphemeralPublicRequest)(nil),  // 0: ExchangeEphemeralPublicRequest
	(*ExchangeEphemeralPublicResponse)(nil), // 1: ExchangeEphemeralPublicResponse
	(*IsEnrolledRequest)(nil),               // 2: IsEnrolledRequest
	(*IsEnrolledResponse)(nil),              // 3: IsEnrolledResponse
	(*SaveEnrollmentInfoRequest)(nil),       // 4: SaveEnrollmentInfoRequest
	(*SaveEnrollmentInfoResponse)(nil),      // 5: SaveEnrollmentInfoResponse
	(*GetSaltAndSRPRequest)(nil),            // 6: GetSaltAndSRPRequest
	(*GetSaltAndSRPResponse)(nil),           // 7: GetSaltAndSRPResponse
	(*GetSysInfoRequest)(nil),               // 8: GetSysInfoRequest
	(*GetSysInfoResponse)(nil),              // 9: GetSysInfoResponse
	(*GetAppsRequest)(nil),                  // 10: GetAppsRequest
	(*GetAppsResponse)(nil),                 // 11: GetAppsResponse
	(*DeleteAppsRequest)(nil),               // 12: DeleteAppsRequest
	(*DeleteAppsResponse)(nil),              // 13: DeleteAppsResponse
	(*GetSysInfoResponse_SysInfo)(nil),      // 14: GetSysInfoResponse.SysInfo
	(*GetAppsResponse_ApplicationInfo)(nil), // 15: GetAppsResponse.ApplicationInfo
}
var file_protos_worker_proto_depIdxs = []int32{
	14, // 0: GetSysInfoResponse.workerSysInfo:type_name -> GetSysInfoResponse.SysInfo
	15, // 1: GetAppsResponse.applications:type_name -> GetAppsResponse.ApplicationInfo
	10, // 2: Worker.GetApps:input_type -> GetAppsRequest
	8,  // 3: Worker.GetSysInfo:input_type -> GetSysInfoRequest
	12, // 4: Worker.DeleteApp:input_type -> DeleteAppsRequest
	2,  // 5: Worker.IsEnrolled:input_type -> IsEnrolledRequest
	4,  // 6: Worker.SaveEnrollmentInfo:input_type -> SaveEnrollmentInfoRequest
	6,  // 7: Worker.GetSaltAndSRP:input_type -> GetSaltAndSRPRequest
	0,  // 8: Worker.ExchangeEphemeralPublic:input_type -> ExchangeEphemeralPublicRequest
	11, // 9: Worker.GetApps:output_type -> GetAppsResponse
	9,  // 10: Worker.GetSysInfo:output_type -> GetSysInfoResponse
	13, // 11: Worker.DeleteApp:output_type -> DeleteAppsResponse
	3,  // 12: Worker.IsEnrolled:output_type -> IsEnrolledResponse
	5,  // 13: Worker.SaveEnrollmentInfo:output_type -> SaveEnrollmentInfoResponse
	7,  // 14: Worker.GetSaltAndSRP:output_type -> GetSaltAndSRPResponse
	1,  // 15: Worker.ExchangeEphemeralPublic:output_type -> ExchangeEphemeralPublicResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_protos_worker_proto_init() }
func file_protos_worker_proto_init() {
	if File_protos_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeEphemeralPublicRequest); i {
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
			switch v := v.(*ExchangeEphemeralPublicResponse); i {
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
			switch v := v.(*IsEnrolledRequest); i {
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
			switch v := v.(*IsEnrolledResponse); i {
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
			switch v := v.(*SaveEnrollmentInfoRequest); i {
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
		file_protos_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveEnrollmentInfoResponse); i {
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
		file_protos_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSaltAndSRPRequest); i {
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
		file_protos_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSaltAndSRPResponse); i {
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
		file_protos_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSysInfoRequest); i {
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
		file_protos_worker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSysInfoResponse); i {
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
		file_protos_worker_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_worker_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_worker_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_worker_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_worker_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSysInfoResponse_SysInfo); i {
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
		file_protos_worker_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   16,
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
