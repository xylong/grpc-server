// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.3
// source: enum.proto

package pb

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

type EmployeeStatus int32

const (
	EmployeeStatus_NORMAL     EmployeeStatus = 0
	EmployeeStatus_ONVACATION EmployeeStatus = 1
	EmployeeStatus_RESIGNED   EmployeeStatus = 2
	EmployeeStatus_RETIRED    EmployeeStatus = 3
)

// Enum value maps for EmployeeStatus.
var (
	EmployeeStatus_name = map[int32]string{
		0: "NORMAL",
		1: "ONVACATION",
		2: "RESIGNED",
		3: "RETIRED",
	}
	EmployeeStatus_value = map[string]int32{
		"NORMAL":     0,
		"ONVACATION": 1,
		"RESIGNED":   2,
		"RETIRED":    3,
	}
)

func (x EmployeeStatus) Enum() *EmployeeStatus {
	p := new(EmployeeStatus)
	*p = x
	return p
}

func (x EmployeeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmployeeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[0].Descriptor()
}

func (EmployeeStatus) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[0]
}

func (x EmployeeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmployeeStatus.Descriptor instead.
func (EmployeeStatus) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{0}
}

var File_enum_proto protoreflect.FileDescriptor

var file_enum_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x47, 0x0a, 0x0e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x4e,
	0x56, 0x41, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x54, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x03, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enum_proto_rawDescOnce sync.Once
	file_enum_proto_rawDescData = file_enum_proto_rawDesc
)

func file_enum_proto_rawDescGZIP() []byte {
	file_enum_proto_rawDescOnce.Do(func() {
		file_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_proto_rawDescData)
	})
	return file_enum_proto_rawDescData
}

var file_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enum_proto_goTypes = []interface{}{
	(EmployeeStatus)(0), // 0: EmployeeStatus
}
var file_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enum_proto_init() }
func file_enum_proto_init() {
	if File_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_proto_goTypes,
		DependencyIndexes: file_enum_proto_depIdxs,
		EnumInfos:         file_enum_proto_enumTypes,
	}.Build()
	File_enum_proto = out.File
	file_enum_proto_rawDesc = nil
	file_enum_proto_goTypes = nil
	file_enum_proto_depIdxs = nil
}
