// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/udpa/annotations/status.proto

package annotations

import (
	reflect "reflect"
	sync "sync"

	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PackageVersionStatus int32

const (
	// Unknown package version status.
	PackageVersionStatus_UNKNOWN PackageVersionStatus = 0
	// This version of the package is frozen.
	PackageVersionStatus_FROZEN PackageVersionStatus = 1
	// This version of the package is the active development version.
	PackageVersionStatus_ACTIVE PackageVersionStatus = 2
	// This version of the package is the candidate for the next major version. It
	// is typically machine generated from the active development version.
	PackageVersionStatus_NEXT_MAJOR_VERSION_CANDIDATE PackageVersionStatus = 3
)

// Enum value maps for PackageVersionStatus.
var (
	PackageVersionStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "FROZEN",
		2: "ACTIVE",
		3: "NEXT_MAJOR_VERSION_CANDIDATE",
	}
	PackageVersionStatus_value = map[string]int32{
		"UNKNOWN":                      0,
		"FROZEN":                       1,
		"ACTIVE":                       2,
		"NEXT_MAJOR_VERSION_CANDIDATE": 3,
	}
)

func (x PackageVersionStatus) Enum() *PackageVersionStatus {
	p := new(PackageVersionStatus)
	*p = x
	return p
}

func (x PackageVersionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackageVersionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_enumTypes[0].Descriptor()
}

func (PackageVersionStatus) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_enumTypes[0]
}

func (x PackageVersionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackageVersionStatus.Descriptor instead.
func (PackageVersionStatus) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDescGZIP(), []int{0}
}

type StatusAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The entity is work-in-progress and subject to breaking changes.
	WorkInProgress bool `protobuf:"varint,1,opt,name=work_in_progress,json=workInProgress,proto3" json:"work_in_progress,omitempty"`
	// The entity belongs to a package with the given version status.
	PackageVersionStatus PackageVersionStatus `protobuf:"varint,2,opt,name=package_version_status,json=packageVersionStatus,proto3,enum=solo.io.udpa.annotations.PackageVersionStatus" json:"package_version_status,omitempty"`
}

func (x *StatusAnnotation) Reset() {
	*x = StatusAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAnnotation) ProtoMessage() {}

func (x *StatusAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAnnotation.ProtoReflect.Descriptor instead.
func (*StatusAnnotation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDescGZIP(), []int{0}
}

func (x *StatusAnnotation) GetWorkInProgress() bool {
	if x != nil {
		return x.WorkInProgress
	}
	return false
}

func (x *StatusAnnotation) GetPackageVersionStatus() PackageVersionStatus {
	if x != nil {
		return x.PackageVersionStatus
	}
	return PackageVersionStatus_UNKNOWN
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*StatusAnnotation)(nil),
		Field:         254737244,
		Name:          "solo.io.udpa.annotations.file_status",
		Tag:           "bytes,254737244,opt,name=file_status",
		Filename:      "github.com/solo-io/gloo/projects/gloo/api/external/udpa/annotations/status.proto",
	},
}

// Extension fields to descriptor.FileOptions.
var (
	// optional solo.io.udpa.annotations.StatusAnnotation file_status = 254737244;
	E_FileStatus = &file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_extTypes[0]
)

var File_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDesc = []byte{
	0x0a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x18, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x75, 0x64, 0x70, 0x61,
	0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2,
	0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x64, 0x0a,
	0x16, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x14, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2a, 0x5d, 0x0a, 0x14, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x4f, 0x5a,
	0x45, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02,
	0x12, 0x20, 0x0a, 0x1c, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x4d, 0x41, 0x4a, 0x4f, 0x52, 0x5f, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45,
	0x10, 0x03, 0x3a, 0x6c, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xdc, 0xf6, 0xbb, 0x79, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_goTypes = []interface{}{
	(PackageVersionStatus)(0),      // 0: solo.io.udpa.annotations.PackageVersionStatus
	(*StatusAnnotation)(nil),       // 1: solo.io.udpa.annotations.StatusAnnotation
	(*descriptor.FileOptions)(nil), // 2: google.protobuf.FileOptions
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_depIdxs = []int32{
	0, // 0: solo.io.udpa.annotations.StatusAnnotation.package_version_status:type_name -> solo.io.udpa.annotations.PackageVersionStatus
	2, // 1: solo.io.udpa.annotations.file_status:extendee -> google.protobuf.FileOptions
	1, // 2: solo.io.udpa.annotations.file_status:type_name -> solo.io.udpa.annotations.StatusAnnotation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusAnnotation); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_msgTypes,
		ExtensionInfos:    file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_extTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_udpa_annotations_status_proto_depIdxs = nil
}
