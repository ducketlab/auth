//go:generate  mingo enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: pkg/micro/pb/service.proto

package micro

import (
	_ "github.com/ducketlab/mingo/cmd/protoc-gen-go-ext/extension/tag"
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

var File_pkg_micro_pb_service_proto protoreflect.FileDescriptor

var file_pkg_micro_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c, 0x61, 0x62, 0x2f, 0x6d,
	0x69, 0x6e, 0x67, 0x6f, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x70, 0x62,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x6b,
	0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcd, 0x03, 0x0a, 0x0c, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x18, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x12, 0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x12, 0x3e, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x12, 0x46, 0x0a, 0x0f, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x12, 0x42, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x12, 0x51, 0x0a, 0x1a, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c, 0x61, 0x62,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_micro_pb_service_proto_goTypes = []interface{}{
	(*ValidateClientCredentialRequest)(nil), // 0: auth.micro.ValidateClientCredentialRequest
	(*CreateMicroRequest)(nil),              // 1: auth.micro.CreateMicroRequest
	(*QueryMicroRequest)(nil),               // 2: auth.micro.QueryMicroRequest
	(*DescribeMicroRequest)(nil),            // 3: auth.micro.DescribeMicroRequest
	(*DeleteMicroRequest)(nil),              // 4: auth.micro.DeleteMicroRequest
	(*Micro)(nil),                           // 5: auth.micro.Micro
	(*Set)(nil),                             // 6: auth.micro.Set
}
var file_pkg_micro_pb_service_proto_depIdxs = []int32{
	0, // 0: auth.micro.MicroService.ValidateClientCredential:input_type -> auth.micro.ValidateClientCredentialRequest
	1, // 1: auth.micro.MicroService.CreateService:input_type -> auth.micro.CreateMicroRequest
	2, // 2: auth.micro.MicroService.QueryService:input_type -> auth.micro.QueryMicroRequest
	3, // 3: auth.micro.MicroService.DescribeService:input_type -> auth.micro.DescribeMicroRequest
	4, // 4: auth.micro.MicroService.DeleteService:input_type -> auth.micro.DeleteMicroRequest
	3, // 5: auth.micro.MicroService.RefreshServiceClientSecret:input_type -> auth.micro.DescribeMicroRequest
	5, // 6: auth.micro.MicroService.ValidateClientCredential:output_type -> auth.micro.Micro
	5, // 7: auth.micro.MicroService.CreateService:output_type -> auth.micro.Micro
	6, // 8: auth.micro.MicroService.QueryService:output_type -> auth.micro.Set
	5, // 9: auth.micro.MicroService.DescribeService:output_type -> auth.micro.Micro
	5, // 10: auth.micro.MicroService.DeleteService:output_type -> auth.micro.Micro
	5, // 11: auth.micro.MicroService.RefreshServiceClientSecret:output_type -> auth.micro.Micro
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_micro_pb_service_proto_init() }
func file_pkg_micro_pb_service_proto_init() {
	if File_pkg_micro_pb_service_proto != nil {
		return
	}
	file_pkg_micro_pb_micro_proto_init()
	file_pkg_micro_pb_request_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_micro_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_micro_pb_service_proto_goTypes,
		DependencyIndexes: file_pkg_micro_pb_service_proto_depIdxs,
	}.Build()
	File_pkg_micro_pb_service_proto = out.File
	file_pkg_micro_pb_service_proto_rawDesc = nil
	file_pkg_micro_pb_service_proto_goTypes = nil
	file_pkg_micro_pb_service_proto_depIdxs = nil
}
