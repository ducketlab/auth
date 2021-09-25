//go:generate  mingo enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: pkg/user/pb/types.proto

package user

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

type UserType int32

const (
	UserType_SUB          UserType = 0
	UserType_PRIMARY      UserType = 1
	UserType_SUPPER       UserType = 2
	UserType_INTERNAL     UserType = 3
	UserType_DOMAIN_ADMIN UserType = 4
	UserType_ORG_ADMIN    UserType = 5
	UserType_PERM_ADMIN   UserType = 6
	UserType_AUDIT_ADMIN  UserType = 7
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "SUB",
		1: "PRIMARY",
		2: "SUPPER",
		3: "INTERNAL",
		4: "DOMAIN_ADMIN",
		5: "ORG_ADMIN",
		6: "PERM_ADMIN",
		7: "AUDIT_ADMIN",
	}
	UserType_value = map[string]int32{
		"SUB":          0,
		"PRIMARY":      1,
		"SUPPER":       2,
		"INTERNAL":     3,
		"DOMAIN_ADMIN": 4,
		"ORG_ADMIN":    5,
		"PERM_ADMIN":   6,
		"AUDIT_ADMIN":  7,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_user_pb_types_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_pkg_user_pb_types_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_user_pb_types_proto_rawDescGZIP(), []int{0}
}

var File_pkg_user_pb_types_proto protoreflect.FileDescriptor

var file_pkg_user_pb_types_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2a, 0x7c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x53, 0x55, 0x42, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49,
	0x4d, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x50, 0x50, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x03,
	0x12, 0x10, 0x0a, 0x0c, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x52, 0x47, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10,
	0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x45, 0x52, 0x4d, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10,
	0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x10, 0x07, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c, 0x61, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_user_pb_types_proto_rawDescOnce sync.Once
	file_pkg_user_pb_types_proto_rawDescData = file_pkg_user_pb_types_proto_rawDesc
)

func file_pkg_user_pb_types_proto_rawDescGZIP() []byte {
	file_pkg_user_pb_types_proto_rawDescOnce.Do(func() {
		file_pkg_user_pb_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_user_pb_types_proto_rawDescData)
	})
	return file_pkg_user_pb_types_proto_rawDescData
}

var file_pkg_user_pb_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_user_pb_types_proto_goTypes = []interface{}{
	(UserType)(0), // 0: auth.user.UserType
}
var file_pkg_user_pb_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_user_pb_types_proto_init() }
func file_pkg_user_pb_types_proto_init() {
	if File_pkg_user_pb_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_user_pb_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_user_pb_types_proto_goTypes,
		DependencyIndexes: file_pkg_user_pb_types_proto_depIdxs,
		EnumInfos:         file_pkg_user_pb_types_proto_enumTypes,
	}.Build()
	File_pkg_user_pb_types_proto = out.File
	file_pkg_user_pb_types_proto_rawDesc = nil
	file_pkg_user_pb_types_proto_goTypes = nil
	file_pkg_user_pb_types_proto_depIdxs = nil
}
