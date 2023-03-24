// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: app/jobs/proto3/loginjobs.proto

package protomodel

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

type LoginJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	TenantId int64  `protobuf:"varint,3,opt,name=TenantId,proto3" json:"TenantId,omitempty"`
	UserType int32  `protobuf:"varint,4,opt,name=UserType,proto3" json:"UserType,omitempty"`
}

func (x *LoginJob) Reset() {
	*x = LoginJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_jobs_proto3_loginjobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginJob) ProtoMessage() {}

func (x *LoginJob) ProtoReflect() protoreflect.Message {
	mi := &file_app_jobs_proto3_loginjobs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginJob.ProtoReflect.Descriptor instead.
func (*LoginJob) Descriptor() ([]byte, []int) {
	return file_app_jobs_proto3_loginjobs_proto_rawDescGZIP(), []int{0}
}

func (x *LoginJob) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginJob) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LoginJob) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *LoginJob) GetUserType() int32 {
	if x != nil {
		return x.UserType
	}
	return 0
}

var File_app_jobs_proto3_loginjobs_proto protoreflect.FileDescriptor

var file_app_jobs_proto3_loginjobs_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x76, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_app_jobs_proto3_loginjobs_proto_rawDescOnce sync.Once
	file_app_jobs_proto3_loginjobs_proto_rawDescData = file_app_jobs_proto3_loginjobs_proto_rawDesc
)

func file_app_jobs_proto3_loginjobs_proto_rawDescGZIP() []byte {
	file_app_jobs_proto3_loginjobs_proto_rawDescOnce.Do(func() {
		file_app_jobs_proto3_loginjobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_jobs_proto3_loginjobs_proto_rawDescData)
	})
	return file_app_jobs_proto3_loginjobs_proto_rawDescData
}

var file_app_jobs_proto3_loginjobs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_jobs_proto3_loginjobs_proto_goTypes = []interface{}{
	(*LoginJob)(nil), // 0: LoginJob
}
var file_app_jobs_proto3_loginjobs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_jobs_proto3_loginjobs_proto_init() }
func file_app_jobs_proto3_loginjobs_proto_init() {
	if File_app_jobs_proto3_loginjobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_jobs_proto3_loginjobs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginJob); i {
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
			RawDescriptor: file_app_jobs_proto3_loginjobs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_jobs_proto3_loginjobs_proto_goTypes,
		DependencyIndexes: file_app_jobs_proto3_loginjobs_proto_depIdxs,
		MessageInfos:      file_app_jobs_proto3_loginjobs_proto_msgTypes,
	}.Build()
	File_app_jobs_proto3_loginjobs_proto = out.File
	file_app_jobs_proto3_loginjobs_proto_rawDesc = nil
	file_app_jobs_proto3_loginjobs_proto_goTypes = nil
	file_app_jobs_proto3_loginjobs_proto_depIdxs = nil
}
