//go:generate  mingo enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: pkg/domain/pb/service.proto

package domain

import (
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

var File_pkg_domain_pb_service_proto protoreflect.FileDescriptor

var file_pkg_domain_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1a, 0x70, 0x6b, 0x67, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xaa, 0x02, 0x0a, 0x0d, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x49, 0x0a, 0x0e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x22,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c, 0x61, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_domain_pb_service_proto_goTypes = []interface{}{
	(*CreateDomainRequest)(nil),   // 0: auth.domain.CreateDomainRequest
	(*DescribeDomainRequest)(nil), // 1: auth.domain.DescribeDomainRequest
	(*QueryDomainRequest)(nil),    // 2: auth.domain.QueryDomainRequest
	(*DeleteDomainRequest)(nil),   // 3: auth.domain.DeleteDomainRequest
	(*Domain)(nil),                // 4: auth.domain.Domain
	(*Set)(nil),                   // 5: auth.domain.Set
}
var file_pkg_domain_pb_service_proto_depIdxs = []int32{
	0, // 0: auth.domain.DomainService.CreateDomain:input_type -> auth.domain.CreateDomainRequest
	1, // 1: auth.domain.DomainService.DescribeDomain:input_type -> auth.domain.DescribeDomainRequest
	2, // 2: auth.domain.DomainService.QueryDomain:input_type -> auth.domain.QueryDomainRequest
	3, // 3: auth.domain.DomainService.DeleteDomain:input_type -> auth.domain.DeleteDomainRequest
	4, // 4: auth.domain.DomainService.CreateDomain:output_type -> auth.domain.Domain
	4, // 5: auth.domain.DomainService.DescribeDomain:output_type -> auth.domain.Domain
	5, // 6: auth.domain.DomainService.QueryDomain:output_type -> auth.domain.Set
	4, // 7: auth.domain.DomainService.DeleteDomain:output_type -> auth.domain.Domain
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_domain_pb_service_proto_init() }
func file_pkg_domain_pb_service_proto_init() {
	if File_pkg_domain_pb_service_proto != nil {
		return
	}
	file_pkg_domain_pb_domain_proto_init()
	file_pkg_domain_pb_request_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_domain_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_domain_pb_service_proto_goTypes,
		DependencyIndexes: file_pkg_domain_pb_service_proto_depIdxs,
	}.Build()
	File_pkg_domain_pb_service_proto = out.File
	file_pkg_domain_pb_service_proto_rawDesc = nil
	file_pkg_domain_pb_service_proto_goTypes = nil
	file_pkg_domain_pb_service_proto_depIdxs = nil
}
