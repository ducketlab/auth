// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: pkg/permission/pb/service.proto

package permission

import (
	role "github.com/ducketlab/auth/pkg/role"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pkg_permission_pb_service_proto protoreflect.FileDescriptor

var file_pkg_permission_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfc, 0x01, 0x0a, 0x11,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x54, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x3e, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x12, 0x51, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c,
	0x61, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_permission_pb_service_proto_goTypes = []interface{}{
	(*QueryPermissionRequest)(nil), // 0: auth.permission.QueryPermissionRequest
	(*QueryRoleRequest)(nil),       // 1: auth.permission.QueryRoleRequest
	(*CheckPermissionRequest)(nil), // 2: auth.permission.CheckPermissionRequest
	(*role.PermissionSet)(nil),     // 3: auth.role.PermissionSet
	(*role.Set)(nil),               // 4: auth.role.Set
	(*role.Permission)(nil),        // 5: auth.role.Permission
}
var file_pkg_permission_pb_service_proto_depIdxs = []int32{
	0, // 0: auth.permission.PermissionService.QueryPermission:input_type -> auth.permission.QueryPermissionRequest
	1, // 1: auth.permission.PermissionService.QueryRole:input_type -> auth.permission.QueryRoleRequest
	2, // 2: auth.permission.PermissionService.CheckPermission:input_type -> auth.permission.CheckPermissionRequest
	3, // 3: auth.permission.PermissionService.QueryPermission:output_type -> auth.role.PermissionSet
	4, // 4: auth.permission.PermissionService.QueryRole:output_type -> auth.role.Set
	5, // 5: auth.permission.PermissionService.CheckPermission:output_type -> auth.role.Permission
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_permission_pb_service_proto_init() }
func file_pkg_permission_pb_service_proto_init() {
	if File_pkg_permission_pb_service_proto != nil {
		return
	}
	file_pkg_permission_pb_request_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_permission_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_permission_pb_service_proto_goTypes,
		DependencyIndexes: file_pkg_permission_pb_service_proto_depIdxs,
	}.Build()
	File_pkg_permission_pb_service_proto = out.File
	file_pkg_permission_pb_service_proto_rawDesc = nil
	file_pkg_permission_pb_service_proto_goTypes = nil
	file_pkg_permission_pb_service_proto_depIdxs = nil
}
