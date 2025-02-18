// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: configcenter.proto

package pb

import (
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// GetGroups
type GetGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID string `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	Keyword  string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	PageNum  int64  `protobuf:"varint,3,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int64  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *GetGroupRequest) Reset() {
	*x = GetGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configcenter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupRequest) ProtoMessage() {}

func (x *GetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configcenter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return file_configcenter_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *GetGroupRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *GetGroupRequest) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *GetGroupRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Groups `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetGroupResponse) Reset() {
	*x = GetGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configcenter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupResponse) ProtoMessage() {}

func (x *GetGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_configcenter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupResponse.ProtoReflect.Descriptor instead.
func (*GetGroupResponse) Descriptor() ([]byte, []int) {
	return file_configcenter_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupResponse) GetData() *Groups {
	if x != nil {
		return x.Data
	}
	return nil
}

type Groups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []string `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Groups) Reset() {
	*x = Groups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configcenter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Groups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Groups) ProtoMessage() {}

func (x *Groups) ProtoReflect() protoreflect.Message {
	mi := &file_configcenter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Groups.ProtoReflect.Descriptor instead.
func (*Groups) Descriptor() ([]byte, []int) {
	return file_configcenter_proto_rawDescGZIP(), []int{2}
}

func (x *Groups) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Groups) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

// GetGroupProperties
type GetGroupPropertiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID string `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	GroupID  string `protobuf:"bytes,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
}

func (x *GetGroupPropertiesRequest) Reset() {
	*x = GetGroupPropertiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configcenter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupPropertiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupPropertiesRequest) ProtoMessage() {}

func (x *GetGroupPropertiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configcenter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupPropertiesRequest.ProtoReflect.Descriptor instead.
func (*GetGroupPropertiesRequest) Descriptor() ([]byte, []int) {
	return file_configcenter_proto_rawDescGZIP(), []int{3}
}

func (x *GetGroupPropertiesRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *GetGroupPropertiesRequest) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

type GetGroupPropertiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*GroupProperties `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetGroupPropertiesResponse) Reset() {
	*x = GetGroupPropertiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configcenter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupPropertiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupPropertiesResponse) ProtoMessage() {}

func (x *GetGroupPropertiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_configcenter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupPropertiesResponse.ProtoReflect.Descriptor instead.
func (*GetGroupPropertiesResponse) Descriptor() ([]byte, []int) {
	return file_configcenter_proto_rawDescGZIP(), []int{4}
}

func (x *GetGroupPropertiesResponse) GetData() []*GroupProperties {
	if x != nil {
		return x.Data
	}
	return nil
}

// SaveGroupProperties
type SaveGroupPropertiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID   string      `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	GroupID    string      `protobuf:"bytes,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
	Properties []*Property `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *SaveGroupPropertiesRequest) Reset() {
	*x = SaveGroupPropertiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configcenter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveGroupPropertiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveGroupPropertiesRequest) ProtoMessage() {}

func (x *SaveGroupPropertiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configcenter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveGroupPropertiesRequest.ProtoReflect.Descriptor instead.
func (*SaveGroupPropertiesRequest) Descriptor() ([]byte, []int) {
	return file_configcenter_proto_rawDescGZIP(), []int{5}
}

func (x *SaveGroupPropertiesRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *SaveGroupPropertiesRequest) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *SaveGroupPropertiesRequest) GetProperties() []*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

type SaveGroupPropertiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SaveGroupPropertiesResponse) Reset() {
	*x = SaveGroupPropertiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configcenter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveGroupPropertiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveGroupPropertiesResponse) ProtoMessage() {}

func (x *SaveGroupPropertiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_configcenter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveGroupPropertiesResponse.ProtoReflect.Descriptor instead.
func (*SaveGroupPropertiesResponse) Descriptor() ([]byte, []int) {
	return file_configcenter_proto_rawDescGZIP(), []int{6}
}

func (x *SaveGroupPropertiesResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type GroupProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group      string      `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Properties []*Property `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *GroupProperties) Reset() {
	*x = GroupProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configcenter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupProperties) ProtoMessage() {}

func (x *GroupProperties) ProtoReflect() protoreflect.Message {
	mi := &file_configcenter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupProperties.ProtoReflect.Descriptor instead.
func (*GroupProperties) Descriptor() ([]byte, []int) {
	return file_configcenter_proto_rawDescGZIP(), []int{7}
}

func (x *GroupProperties) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GroupProperties) GetProperties() []*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configcenter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_configcenter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_configcenter_proto_rawDescGZIP(), []int{8}
}

func (x *Property) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Property) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Property) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_configcenter_proto protoreflect.FileDescriptor

var file_configcenter_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x45, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x32, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22, 0x58, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xa3, 0x01, 0x0a, 0x1a, 0x53, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x72, 0x64,
	0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x1b, 0x53, 0x61, 0x76, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x0f, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x3f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x32, 0xb9, 0x04, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x26, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x12,
	0x3a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x3f, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x6f, 0x3d, 0x7b, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x7d, 0x12, 0xb6, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x30, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12,
	0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x44, 0x7d, 0x12, 0xc5, 0x01, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x65,
	0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x22, 0x33, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x7d, 0x2f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x7d,
	0x3a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_configcenter_proto_rawDescOnce sync.Once
	file_configcenter_proto_rawDescData = file_configcenter_proto_rawDesc
)

func file_configcenter_proto_rawDescGZIP() []byte {
	file_configcenter_proto_rawDescOnce.Do(func() {
		file_configcenter_proto_rawDescData = protoimpl.X.CompressGZIP(file_configcenter_proto_rawDescData)
	})
	return file_configcenter_proto_rawDescData
}

var file_configcenter_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_configcenter_proto_goTypes = []interface{}{
	(*GetGroupRequest)(nil),             // 0: erda.msp.configcenter.GetGroupRequest
	(*GetGroupResponse)(nil),            // 1: erda.msp.configcenter.GetGroupResponse
	(*Groups)(nil),                      // 2: erda.msp.configcenter.Groups
	(*GetGroupPropertiesRequest)(nil),   // 3: erda.msp.configcenter.GetGroupPropertiesRequest
	(*GetGroupPropertiesResponse)(nil),  // 4: erda.msp.configcenter.GetGroupPropertiesResponse
	(*SaveGroupPropertiesRequest)(nil),  // 5: erda.msp.configcenter.SaveGroupPropertiesRequest
	(*SaveGroupPropertiesResponse)(nil), // 6: erda.msp.configcenter.SaveGroupPropertiesResponse
	(*GroupProperties)(nil),             // 7: erda.msp.configcenter.GroupProperties
	(*Property)(nil),                    // 8: erda.msp.configcenter.Property
}
var file_configcenter_proto_depIdxs = []int32{
	2, // 0: erda.msp.configcenter.GetGroupResponse.data:type_name -> erda.msp.configcenter.Groups
	7, // 1: erda.msp.configcenter.GetGroupPropertiesResponse.data:type_name -> erda.msp.configcenter.GroupProperties
	8, // 2: erda.msp.configcenter.SaveGroupPropertiesRequest.properties:type_name -> erda.msp.configcenter.Property
	8, // 3: erda.msp.configcenter.GroupProperties.properties:type_name -> erda.msp.configcenter.Property
	0, // 4: erda.msp.configcenter.ConfigCenterService.GetGroups:input_type -> erda.msp.configcenter.GetGroupRequest
	3, // 5: erda.msp.configcenter.ConfigCenterService.GetGroupProperties:input_type -> erda.msp.configcenter.GetGroupPropertiesRequest
	5, // 6: erda.msp.configcenter.ConfigCenterService.SaveGroupProperties:input_type -> erda.msp.configcenter.SaveGroupPropertiesRequest
	1, // 7: erda.msp.configcenter.ConfigCenterService.GetGroups:output_type -> erda.msp.configcenter.GetGroupResponse
	4, // 8: erda.msp.configcenter.ConfigCenterService.GetGroupProperties:output_type -> erda.msp.configcenter.GetGroupPropertiesResponse
	6, // 9: erda.msp.configcenter.ConfigCenterService.SaveGroupProperties:output_type -> erda.msp.configcenter.SaveGroupPropertiesResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_configcenter_proto_init() }
func file_configcenter_proto_init() {
	if File_configcenter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configcenter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupRequest); i {
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
		file_configcenter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupResponse); i {
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
		file_configcenter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Groups); i {
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
		file_configcenter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupPropertiesRequest); i {
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
		file_configcenter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupPropertiesResponse); i {
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
		file_configcenter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveGroupPropertiesRequest); i {
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
		file_configcenter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveGroupPropertiesResponse); i {
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
		file_configcenter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupProperties); i {
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
		file_configcenter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
			RawDescriptor: file_configcenter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_configcenter_proto_goTypes,
		DependencyIndexes: file_configcenter_proto_depIdxs,
		MessageInfos:      file_configcenter_proto_msgTypes,
	}.Build()
	File_configcenter_proto = out.File
	file_configcenter_proto_rawDesc = nil
	file_configcenter_proto_goTypes = nil
	file_configcenter_proto_depIdxs = nil
}
