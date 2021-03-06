// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: pkg/role/pb/request.proto

package role

import (
	_ "github.com/ducketlab/mingo/cmd/protoc-gen-go-ext/extension/tag"
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

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        RoleType                   `protobuf:"varint,1,opt,name=type,proto3,enum=auth.role.RoleType" json:"type" bson:"type"`
	Name        string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name" validate:"required,lte=30"`
	Description string                     `protobuf:"bytes,3,opt,name=description,proto3" json:"description" bson:"description" validate:"lte=400"`
	Meta        map[string]string          `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"meta" validate:"lte=400"`
	Permissions []*CreatePermissionRequest `protobuf:"bytes,9,rep,name=permissions,proto3" json:"permissions" bson:"permissions"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoleRequest) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_NULL
}

func (x *CreateRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoleRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRoleRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CreateRoleRequest) GetPermissions() []*CreatePermissionRequest {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type CreatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Effect       EffectType `protobuf:"varint,1,opt,name=effect,proto3,enum=auth.role.EffectType" json:"effect" bson:"effect"`
	ServiceId    string     `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id" bson:"service_id"`
	ResourceName string     `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name" bson:"resource_name"`
	LabelKey     string     `protobuf:"bytes,4,opt,name=label_key,json=labelKey,proto3" json:"label_key" bson:"label_key"`
	MatchAll     bool       `protobuf:"varint,5,opt,name=match_all,json=matchAll,proto3" json:"match_all" bson:"match_all"`
	LabelValues  []string   `protobuf:"bytes,6,rep,name=label_values,json=labelValues,proto3" json:"label_values" bson:"label_values"`
}

func (x *CreatePermissionRequest) Reset() {
	*x = CreatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionRequest) ProtoMessage() {}

func (x *CreatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePermissionRequest.ProtoReflect.Descriptor instead.
func (*CreatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePermissionRequest) GetEffect() EffectType {
	if x != nil {
		return x.Effect
	}
	return EffectType_ALLOW
}

func (x *CreatePermissionRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *CreatePermissionRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CreatePermissionRequest) GetLabelKey() string {
	if x != nil {
		return x.LabelKey
	}
	return ""
}

func (x *CreatePermissionRequest) GetMatchAll() bool {
	if x != nil {
		return x.MatchAll
	}
	return false
}

func (x *CreatePermissionRequest) GetLabelValues() []string {
	if x != nil {
		return x.LabelValues
	}
	return nil
}

type DescribeRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name            string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" validate:"required,lte=64"`
	WithPermissions bool     `protobuf:"varint,3,opt,name=with_permissions,json=withPermissions,proto3" json:"with_permissions" bson:"with_permissions"`
	Type            RoleType `protobuf:"varint,4,opt,name=type,proto3,enum=auth.role.RoleType" json:"type" bson:"type"`
}

func (x *DescribeRoleRequest) Reset() {
	*x = DescribeRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRoleRequest) ProtoMessage() {}

func (x *DescribeRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRoleRequest.ProtoReflect.Descriptor instead.
func (*DescribeRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeRoleRequest) GetWithPermissions() bool {
	if x != nil {
		return x.WithPermissions
	}
	return false
}

func (x *DescribeRoleRequest) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_NULL
}

type DescribePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribePermissionRequest) Reset() {
	*x = DescribePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePermissionRequest) ProtoMessage() {}

func (x *DescribePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePermissionRequest.ProtoReflect.Descriptor instead.
func (*DescribePermissionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *DescribePermissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddPermissionToRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId      string                     `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id" validate:"required,lte=64"`
	Permissions []*CreatePermissionRequest `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions" validate:"required"`
}

func (x *AddPermissionToRoleRequest) Reset() {
	*x = AddPermissionToRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPermissionToRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermissionToRoleRequest) ProtoMessage() {}

func (x *AddPermissionToRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPermissionToRoleRequest.ProtoReflect.Descriptor instead.
func (*AddPermissionToRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *AddPermissionToRoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AddPermissionToRoleRequest) GetPermissions() []*CreatePermissionRequest {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type QueryPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId    string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	SkipItmes bool   `protobuf:"varint,2,opt,name=skip_itmes,json=skipItmes,proto3" json:"skip_itmes,omitempty"`
}

func (x *QueryPermissionRequest) Reset() {
	*x = QueryPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_role_pb_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPermissionRequest) ProtoMessage() {}

func (x *QueryPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_role_pb_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPermissionRequest.ProtoReflect.Descriptor instead.
func (*QueryPermissionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_role_pb_request_proto_rawDescGZIP(), []int{5}
}

func (x *QueryPermissionRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *QueryPermissionRequest) GetSkipItmes() bool {
	if x != nil {
		return x.SkipItmes
	}
	return false
}

var File_pkg_role_pb_request_proto protoreflect.FileDescriptor

var file_pkg_role_pb_request_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c, 0x61, 0x62, 0x2f, 0x6d, 0x69, 0x6e,
	0x67, 0x6f, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x04, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xc2, 0xde, 0x1f, 0x34, 0x0a, 0x32, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x33, 0x30, 0x22, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x60, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3e, 0xc2, 0xde, 0x1f, 0x3a, 0x0a,
	0x38, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x22, 0x6c, 0x74, 0x65, 0x3d, 0x34, 0x30, 0x30, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x30, 0xc2, 0xde,
	0x1f, 0x2c, 0x0a, 0x2a, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x6c, 0x74, 0x65, 0x3d, 0x34, 0x30, 0x30, 0x22, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x71, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x2b, 0xc2,
	0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xe9, 0x03, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x06,
	0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x22, 0x52, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x48,
	0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2f, 0xc2, 0xde, 0x1f, 0x2b, 0x0a, 0x29, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44,
	0x0a, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x08, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x4b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x6c,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x22,
	0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x12, 0x50, 0x0a, 0x0c, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x52,
	0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xac, 0x02, 0x0a,
	0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64,
	0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x36, 0xc2, 0xde, 0x1f, 0x32, 0x0a, 0x30, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x36, 0x34, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x60, 0x0a, 0x10, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x35, 0xc2, 0xde, 0x1f,
	0x31, 0x0a, 0x2f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3c, 0x0a, 0x19, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0xda, 0x01, 0x0a, 0x1a, 0x41, 0x64,
	0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xc2, 0xde, 0x1f, 0x2b, 0x0a,
	0x29, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x36, 0x34, 0x22, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x72, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x2c, 0xc2, 0xde, 0x1f,
	0x28, 0x0a, 0x26, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x48, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x2f, 0xc2, 0xde, 0x1f, 0x2b, 0x0a, 0x29, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c, 0x74, 0x65, 0x3d,
	0x36, 0x34, 0x22, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0a, 0x73,
	0x6b, 0x69, 0x70, 0x5f, 0x69, 0x74, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x6b, 0x69,
	0x70, 0x5f, 0x69, 0x74, 0x6d, 0x65, 0x73, 0x22, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x70, 0x49, 0x74,
	0x6d, 0x65, 0x73, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c, 0x61, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_role_pb_request_proto_rawDescOnce sync.Once
	file_pkg_role_pb_request_proto_rawDescData = file_pkg_role_pb_request_proto_rawDesc
)

func file_pkg_role_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_role_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_role_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_role_pb_request_proto_rawDescData)
	})
	return file_pkg_role_pb_request_proto_rawDescData
}

var file_pkg_role_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_role_pb_request_proto_goTypes = []interface{}{
	(*CreateRoleRequest)(nil),          // 0: auth.role.CreateRoleRequest
	(*CreatePermissionRequest)(nil),    // 1: auth.role.CreatePermissionRequest
	(*DescribeRoleRequest)(nil),        // 2: auth.role.DescribeRoleRequest
	(*DescribePermissionRequest)(nil),  // 3: auth.role.DescribePermissionRequest
	(*AddPermissionToRoleRequest)(nil), // 4: auth.role.AddPermissionToRoleRequest
	(*QueryPermissionRequest)(nil),     // 5: auth.role.QueryPermissionRequest
	nil,                                // 6: auth.role.CreateRoleRequest.MetaEntry
	(RoleType)(0),                      // 7: auth.role.RoleType
	(EffectType)(0),                    // 8: auth.role.EffectType
}
var file_pkg_role_pb_request_proto_depIdxs = []int32{
	7, // 0: auth.role.CreateRoleRequest.type:type_name -> auth.role.RoleType
	6, // 1: auth.role.CreateRoleRequest.meta:type_name -> auth.role.CreateRoleRequest.MetaEntry
	1, // 2: auth.role.CreateRoleRequest.permissions:type_name -> auth.role.CreatePermissionRequest
	8, // 3: auth.role.CreatePermissionRequest.effect:type_name -> auth.role.EffectType
	7, // 4: auth.role.DescribeRoleRequest.type:type_name -> auth.role.RoleType
	1, // 5: auth.role.AddPermissionToRoleRequest.permissions:type_name -> auth.role.CreatePermissionRequest
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_role_pb_request_proto_init() }
func file_pkg_role_pb_request_proto_init() {
	if File_pkg_role_pb_request_proto != nil {
		return
	}
	file_pkg_role_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_role_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_pkg_role_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePermissionRequest); i {
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
		file_pkg_role_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRoleRequest); i {
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
		file_pkg_role_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePermissionRequest); i {
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
		file_pkg_role_pb_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPermissionToRoleRequest); i {
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
		file_pkg_role_pb_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPermissionRequest); i {
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
			RawDescriptor: file_pkg_role_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_role_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_role_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_role_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_role_pb_request_proto = out.File
	file_pkg_role_pb_request_proto_rawDesc = nil
	file_pkg_role_pb_request_proto_goTypes = nil
	file_pkg_role_pb_request_proto_depIdxs = nil
}
