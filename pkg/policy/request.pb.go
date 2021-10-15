// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: pkg/policy/pb/request.proto

package policy

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

type CreatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamespaceId string     `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id" bson:"namespace_id" validate:"lte=120"`
	Account     string     `protobuf:"bytes,2,opt,name=account,proto3" json:"account" bson:"account" validate:"required,lte=120"`
	RoleId      string     `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id" bson:"role_id" validate:"required,lte=40"`
	Scope       string     `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope" bson:"scope"`
	ExpiredTime int64      `protobuf:"varint,5,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time" bson:"expired_time"`
	Type        PolicyType `protobuf:"varint,6,opt,name=type,proto3,enum=auth.policy.PolicyType" json:"type" bson:"type"`
}

func (x *CreatePolicyRequest) Reset() {
	*x = CreatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyRequest) ProtoMessage() {}

func (x *CreatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_policy_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolicyRequest) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *CreatePolicyRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CreatePolicyRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *CreatePolicyRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *CreatePolicyRequest) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

func (x *CreatePolicyRequest) GetType() PolicyType {
	if x != nil {
		return x.Type
	}
	return PolicyType_NULL
}

type DescribePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribePolicyRequest) Reset() {
	*x = DescribePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePolicyRequest) ProtoMessage() {}

func (x *DescribePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePolicyRequest.ProtoReflect.Descriptor instead.
func (*DescribePolicyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_policy_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *DescribePolicyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Account     string     `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	RoleId      string     `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	NamespaceId string     `protobuf:"bytes,4,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	Type        PolicyType `protobuf:"varint,5,opt,name=type,proto3,enum=auth.policy.PolicyType" json:"type,omitempty"`
	Domain      string     `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain"`
}

func (x *DeletePolicyRequest) Reset() {
	*x = DeletePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyRequest) ProtoMessage() {}

func (x *DeletePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyRequest.ProtoReflect.Descriptor instead.
func (*DeletePolicyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_policy_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *DeletePolicyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeletePolicyRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *DeletePolicyRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *DeletePolicyRequest) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *DeletePolicyRequest) GetType() PolicyType {
	if x != nil {
		return x.Type
	}
	return PolicyType_NULL
}

func (x *DeletePolicyRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

var File_pkg_policy_pb_request_proto protoreflect.FileDescriptor

var file_pkg_policy_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c, 0x61, 0x62,
	0x2f, 0x6d, 0x69, 0x6e, 0x67, 0x6f, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83,
	0x04, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x40, 0xc2, 0xde,
	0x1f, 0x3c, 0x0a, 0x3a, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x6c, 0x74, 0x65, 0x3d, 0x31, 0x32, 0x30, 0x22, 0x52, 0x0b,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x59, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0xc2, 0xde,
	0x1f, 0x3b, 0x0a, 0x39, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x31, 0x32, 0x30, 0x22, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x57, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3e, 0xc2, 0xde, 0x1f, 0x3a, 0x0a, 0x38, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c,
	0x6c, 0x74, 0x65, 0x3d, 0x34, 0x30, 0x22, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x35, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f,
	0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x2d, 0xc2, 0xde,
	0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x52, 0x0b, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x38, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a,
	0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0xe8,
	0x02, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xc2, 0xde, 0x1f, 0x1a, 0x0a, 0x18,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2c, 0x6f, 0x6d,
	0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x37, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1e, 0xc2, 0xde, 0x1f, 0x1a, 0x0a, 0x18, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x0c, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x48, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1b, 0xc2, 0xde, 0x1f, 0x17, 0x0a,
	0x15, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xc2, 0xde,
	0x1f, 0x0f, 0x0a, 0x0d, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x22, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c, 0x61,
	0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_policy_pb_request_proto_rawDescOnce sync.Once
	file_pkg_policy_pb_request_proto_rawDescData = file_pkg_policy_pb_request_proto_rawDesc
)

func file_pkg_policy_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_policy_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_policy_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_policy_pb_request_proto_rawDescData)
	})
	return file_pkg_policy_pb_request_proto_rawDescData
}

var file_pkg_policy_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_policy_pb_request_proto_goTypes = []interface{}{
	(*CreatePolicyRequest)(nil),   // 0: auth.policy.CreatePolicyRequest
	(*DescribePolicyRequest)(nil), // 1: auth.policy.DescribePolicyRequest
	(*DeletePolicyRequest)(nil),   // 2: auth.policy.DeletePolicyRequest
	(PolicyType)(0),               // 3: auth.policy.PolicyType
}
var file_pkg_policy_pb_request_proto_depIdxs = []int32{
	3, // 0: auth.policy.CreatePolicyRequest.type:type_name -> auth.policy.PolicyType
	3, // 1: auth.policy.DeletePolicyRequest.type:type_name -> auth.policy.PolicyType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_policy_pb_request_proto_init() }
func file_pkg_policy_pb_request_proto_init() {
	if File_pkg_policy_pb_request_proto != nil {
		return
	}
	file_pkg_policy_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_policy_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyRequest); i {
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
		file_pkg_policy_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePolicyRequest); i {
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
		file_pkg_policy_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyRequest); i {
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
			RawDescriptor: file_pkg_policy_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_policy_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_policy_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_policy_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_policy_pb_request_proto = out.File
	file_pkg_policy_pb_request_proto_rawDesc = nil
	file_pkg_policy_pb_request_proto_goTypes = nil
	file_pkg_policy_pb_request_proto_depIdxs = nil
}
