// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: api/v1/sensor_upgrade_service.proto

package v1

import (
	storage "github.com/stackrox/rox/generated/storage"
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

type UpdateSensorUpgradeConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *storage.SensorUpgradeConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *UpdateSensorUpgradeConfigRequest) Reset() {
	*x = UpdateSensorUpgradeConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_sensor_upgrade_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSensorUpgradeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSensorUpgradeConfigRequest) ProtoMessage() {}

func (x *UpdateSensorUpgradeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_sensor_upgrade_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSensorUpgradeConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateSensorUpgradeConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_sensor_upgrade_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateSensorUpgradeConfigRequest) GetConfig() *storage.SensorUpgradeConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type GetSensorUpgradeConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *storage.SensorUpgradeConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetSensorUpgradeConfigResponse) Reset() {
	*x = GetSensorUpgradeConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_sensor_upgrade_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensorUpgradeConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensorUpgradeConfigResponse) ProtoMessage() {}

func (x *GetSensorUpgradeConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_sensor_upgrade_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensorUpgradeConfigResponse.ProtoReflect.Descriptor instead.
func (*GetSensorUpgradeConfigResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_sensor_upgrade_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSensorUpgradeConfigResponse) GetConfig() *storage.SensorUpgradeConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_api_v1_sensor_upgrade_service_proto protoreflect.FileDescriptor

var file_api_v1_sensor_upgrade_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f,
	0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58,
	0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x56, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x32, 0xc2, 0x03, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x72, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x14, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x79, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x19, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x43, 0x65, 0x72, 0x74, 0x52, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x22, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x63, 0x65, 0x72, 0x74, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x1e, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x5a, 0x02, 0x76, 0x31, 0x58, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_sensor_upgrade_service_proto_rawDescOnce sync.Once
	file_api_v1_sensor_upgrade_service_proto_rawDescData = file_api_v1_sensor_upgrade_service_proto_rawDesc
)

func file_api_v1_sensor_upgrade_service_proto_rawDescGZIP() []byte {
	file_api_v1_sensor_upgrade_service_proto_rawDescOnce.Do(func() {
		file_api_v1_sensor_upgrade_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_sensor_upgrade_service_proto_rawDescData)
	})
	return file_api_v1_sensor_upgrade_service_proto_rawDescData
}

var file_api_v1_sensor_upgrade_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_sensor_upgrade_service_proto_goTypes = []interface{}{
	(*UpdateSensorUpgradeConfigRequest)(nil), // 0: v1.UpdateSensorUpgradeConfigRequest
	(*GetSensorUpgradeConfigResponse)(nil),   // 1: v1.GetSensorUpgradeConfigResponse
	(*storage.SensorUpgradeConfig)(nil),      // 2: storage.SensorUpgradeConfig
	(*Empty)(nil),                            // 3: v1.Empty
	(*ResourceByID)(nil),                     // 4: v1.ResourceByID
}
var file_api_v1_sensor_upgrade_service_proto_depIdxs = []int32{
	2, // 0: v1.UpdateSensorUpgradeConfigRequest.config:type_name -> storage.SensorUpgradeConfig
	2, // 1: v1.GetSensorUpgradeConfigResponse.config:type_name -> storage.SensorUpgradeConfig
	3, // 2: v1.SensorUpgradeService.GetSensorUpgradeConfig:input_type -> v1.Empty
	0, // 3: v1.SensorUpgradeService.UpdateSensorUpgradeConfig:input_type -> v1.UpdateSensorUpgradeConfigRequest
	4, // 4: v1.SensorUpgradeService.TriggerSensorUpgrade:input_type -> v1.ResourceByID
	4, // 5: v1.SensorUpgradeService.TriggerSensorCertRotation:input_type -> v1.ResourceByID
	1, // 6: v1.SensorUpgradeService.GetSensorUpgradeConfig:output_type -> v1.GetSensorUpgradeConfigResponse
	3, // 7: v1.SensorUpgradeService.UpdateSensorUpgradeConfig:output_type -> v1.Empty
	3, // 8: v1.SensorUpgradeService.TriggerSensorUpgrade:output_type -> v1.Empty
	3, // 9: v1.SensorUpgradeService.TriggerSensorCertRotation:output_type -> v1.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_sensor_upgrade_service_proto_init() }
func file_api_v1_sensor_upgrade_service_proto_init() {
	if File_api_v1_sensor_upgrade_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	file_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_sensor_upgrade_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSensorUpgradeConfigRequest); i {
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
		file_api_v1_sensor_upgrade_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensorUpgradeConfigResponse); i {
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
			RawDescriptor: file_api_v1_sensor_upgrade_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_sensor_upgrade_service_proto_goTypes,
		DependencyIndexes: file_api_v1_sensor_upgrade_service_proto_depIdxs,
		MessageInfos:      file_api_v1_sensor_upgrade_service_proto_msgTypes,
	}.Build()
	File_api_v1_sensor_upgrade_service_proto = out.File
	file_api_v1_sensor_upgrade_service_proto_rawDesc = nil
	file_api_v1_sensor_upgrade_service_proto_goTypes = nil
	file_api_v1_sensor_upgrade_service_proto_depIdxs = nil
}
