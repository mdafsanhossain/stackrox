// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: api/v1/central_health_service.proto

package v1

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

type GetUpgradeStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpgradeStatus *CentralUpgradeStatus `protobuf:"bytes,1,opt,name=upgradeStatus,proto3" json:"upgradeStatus,omitempty"`
}

func (x *GetUpgradeStatusResponse) Reset() {
	*x = GetUpgradeStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_central_health_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpgradeStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpgradeStatusResponse) ProtoMessage() {}

func (x *GetUpgradeStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_central_health_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpgradeStatusResponse.ProtoReflect.Descriptor instead.
func (*GetUpgradeStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_central_health_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUpgradeStatusResponse) GetUpgradeStatus() *CentralUpgradeStatus {
	if x != nil {
		return x.UpgradeStatus
	}
	return nil
}

type CentralUpgradeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current Central Version
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The version of previous clone in Central. This is the version we can force rollback to.
	ForceRollbackTo string `protobuf:"bytes,2,opt,name=force_rollback_to,json=forceRollbackTo,proto3" json:"force_rollback_to,omitempty"`
	// If true, we can rollback to the current version if an upgrade failed.
	CanRollbackAfterUpgrade bool `protobuf:"varint,3,opt,name=can_rollback_after_upgrade,json=canRollbackAfterUpgrade,proto3" json:"can_rollback_after_upgrade,omitempty"`
	// Current disk space stats for upgrade
	SpaceRequiredForRollbackAfterUpgrade  int64 `protobuf:"varint,4,opt,name=space_required_for_rollback_after_upgrade,json=spaceRequiredForRollbackAfterUpgrade,proto3" json:"space_required_for_rollback_after_upgrade,omitempty"`
	SpaceAvailableForRollbackAfterUpgrade int64 `protobuf:"varint,5,opt,name=space_available_for_rollback_after_upgrade,json=spaceAvailableForRollbackAfterUpgrade,proto3" json:"space_available_for_rollback_after_upgrade,omitempty"`
}

func (x *CentralUpgradeStatus) Reset() {
	*x = CentralUpgradeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_central_health_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CentralUpgradeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CentralUpgradeStatus) ProtoMessage() {}

func (x *CentralUpgradeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_central_health_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CentralUpgradeStatus.ProtoReflect.Descriptor instead.
func (*CentralUpgradeStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_central_health_service_proto_rawDescGZIP(), []int{1}
}

func (x *CentralUpgradeStatus) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CentralUpgradeStatus) GetForceRollbackTo() string {
	if x != nil {
		return x.ForceRollbackTo
	}
	return ""
}

func (x *CentralUpgradeStatus) GetCanRollbackAfterUpgrade() bool {
	if x != nil {
		return x.CanRollbackAfterUpgrade
	}
	return false
}

func (x *CentralUpgradeStatus) GetSpaceRequiredForRollbackAfterUpgrade() int64 {
	if x != nil {
		return x.SpaceRequiredForRollbackAfterUpgrade
	}
	return 0
}

func (x *CentralUpgradeStatus) GetSpaceAvailableForRollbackAfterUpgrade() int64 {
	if x != nil {
		return x.SpaceAvailableForRollbackAfterUpgrade
	}
	return 0
}

var File_api_v1_central_health_service_proto protoreflect.FileDescriptor

var file_api_v1_central_health_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x75, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcd, 0x02, 0x0a, 0x14, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x54, 0x6f, 0x12, 0x3b, 0x0a, 0x1a, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x6f,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x63, 0x61, 0x6e, 0x52,
	0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x66, 0x74, 0x65, 0x72, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x29, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x24, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x59, 0x0a, 0x2a,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x66, 0x6f, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x25, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x46, 0x6f, 0x72, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x32, 0x7c, 0x0a, 0x14, 0x43, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x64, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x1e, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x5a, 0x02, 0x76, 0x31, 0x58, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_central_health_service_proto_rawDescOnce sync.Once
	file_api_v1_central_health_service_proto_rawDescData = file_api_v1_central_health_service_proto_rawDesc
)

func file_api_v1_central_health_service_proto_rawDescGZIP() []byte {
	file_api_v1_central_health_service_proto_rawDescOnce.Do(func() {
		file_api_v1_central_health_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_central_health_service_proto_rawDescData)
	})
	return file_api_v1_central_health_service_proto_rawDescData
}

var file_api_v1_central_health_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_central_health_service_proto_goTypes = []interface{}{
	(*GetUpgradeStatusResponse)(nil), // 0: v1.GetUpgradeStatusResponse
	(*CentralUpgradeStatus)(nil),     // 1: v1.CentralUpgradeStatus
	(*Empty)(nil),                    // 2: v1.Empty
}
var file_api_v1_central_health_service_proto_depIdxs = []int32{
	1, // 0: v1.GetUpgradeStatusResponse.upgradeStatus:type_name -> v1.CentralUpgradeStatus
	2, // 1: v1.CentralHealthService.GetUpgradeStatus:input_type -> v1.Empty
	0, // 2: v1.CentralHealthService.GetUpgradeStatus:output_type -> v1.GetUpgradeStatusResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_central_health_service_proto_init() }
func file_api_v1_central_health_service_proto_init() {
	if File_api_v1_central_health_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_central_health_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpgradeStatusResponse); i {
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
		file_api_v1_central_health_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CentralUpgradeStatus); i {
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
			RawDescriptor: file_api_v1_central_health_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_central_health_service_proto_goTypes,
		DependencyIndexes: file_api_v1_central_health_service_proto_depIdxs,
		MessageInfos:      file_api_v1_central_health_service_proto_msgTypes,
	}.Build()
	File_api_v1_central_health_service_proto = out.File
	file_api_v1_central_health_service_proto_rawDesc = nil
	file_api_v1_central_health_service_proto_goTypes = nil
	file_api_v1_central_health_service_proto_depIdxs = nil
}
