// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: internalapi/sensor/signal_iservice.proto

package sensor

import (
	v1 "github.com/stackrox/rox/generated/api/v1"
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

// A single message in the event stream between Collector and Sensor.
type SignalStreamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*SignalStreamMessage_CollectorRegisterRequest
	//	*SignalStreamMessage_Signal
	Msg isSignalStreamMessage_Msg `protobuf_oneof:"msg"`
}

func (x *SignalStreamMessage) Reset() {
	*x = SignalStreamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_sensor_signal_iservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalStreamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalStreamMessage) ProtoMessage() {}

func (x *SignalStreamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_signal_iservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalStreamMessage.ProtoReflect.Descriptor instead.
func (*SignalStreamMessage) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_signal_iservice_proto_rawDescGZIP(), []int{0}
}

func (m *SignalStreamMessage) GetMsg() isSignalStreamMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *SignalStreamMessage) GetCollectorRegisterRequest() *CollectorRegisterRequest {
	if x, ok := x.GetMsg().(*SignalStreamMessage_CollectorRegisterRequest); ok {
		return x.CollectorRegisterRequest
	}
	return nil
}

func (x *SignalStreamMessage) GetSignal() *v1.Signal {
	if x, ok := x.GetMsg().(*SignalStreamMessage_Signal); ok {
		return x.Signal
	}
	return nil
}

type isSignalStreamMessage_Msg interface {
	isSignalStreamMessage_Msg()
}

type SignalStreamMessage_CollectorRegisterRequest struct {
	// The first message in every stream that registers Collector with Sensor.
	CollectorRegisterRequest *CollectorRegisterRequest `protobuf:"bytes,1,opt,name=collector_register_request,json=collectorRegisterRequest,proto3,oneof"`
}

type SignalStreamMessage_Signal struct {
	// A signal event observed by Collector.
	Signal *v1.Signal `protobuf:"bytes,2,opt,name=signal,proto3,oneof"`
}

func (*SignalStreamMessage_CollectorRegisterRequest) isSignalStreamMessage_Msg() {}

func (*SignalStreamMessage_Signal) isSignalStreamMessage_Msg() {}

var File_internalapi_sensor_signal_iservice_proto protoreflect.FileDescriptor

var file_internalapi_sensor_signal_iservice_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x1a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa4, 0x01, 0x0a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x60, 0x0a, 0x1a, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x18, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x42,
	0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x4a, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x0b, 0x5a, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0xf8, 0x01, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_sensor_signal_iservice_proto_rawDescOnce sync.Once
	file_internalapi_sensor_signal_iservice_proto_rawDescData = file_internalapi_sensor_signal_iservice_proto_rawDesc
)

func file_internalapi_sensor_signal_iservice_proto_rawDescGZIP() []byte {
	file_internalapi_sensor_signal_iservice_proto_rawDescOnce.Do(func() {
		file_internalapi_sensor_signal_iservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_sensor_signal_iservice_proto_rawDescData)
	})
	return file_internalapi_sensor_signal_iservice_proto_rawDescData
}

var file_internalapi_sensor_signal_iservice_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internalapi_sensor_signal_iservice_proto_goTypes = []interface{}{
	(*SignalStreamMessage)(nil),      // 0: sensor.SignalStreamMessage
	(*CollectorRegisterRequest)(nil), // 1: sensor.CollectorRegisterRequest
	(*v1.Signal)(nil),                // 2: v1.Signal
	(*v1.Empty)(nil),                 // 3: v1.Empty
}
var file_internalapi_sensor_signal_iservice_proto_depIdxs = []int32{
	1, // 0: sensor.SignalStreamMessage.collector_register_request:type_name -> sensor.CollectorRegisterRequest
	2, // 1: sensor.SignalStreamMessage.signal:type_name -> v1.Signal
	0, // 2: sensor.SignalService.PushSignals:input_type -> sensor.SignalStreamMessage
	3, // 3: sensor.SignalService.PushSignals:output_type -> v1.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internalapi_sensor_signal_iservice_proto_init() }
func file_internalapi_sensor_signal_iservice_proto_init() {
	if File_internalapi_sensor_signal_iservice_proto != nil {
		return
	}
	file_internalapi_sensor_collector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internalapi_sensor_signal_iservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalStreamMessage); i {
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
	file_internalapi_sensor_signal_iservice_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SignalStreamMessage_CollectorRegisterRequest)(nil),
		(*SignalStreamMessage_Signal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_sensor_signal_iservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_sensor_signal_iservice_proto_goTypes,
		DependencyIndexes: file_internalapi_sensor_signal_iservice_proto_depIdxs,
		MessageInfos:      file_internalapi_sensor_signal_iservice_proto_msgTypes,
	}.Build()
	File_internalapi_sensor_signal_iservice_proto = out.File
	file_internalapi_sensor_signal_iservice_proto_rawDesc = nil
	file_internalapi_sensor_signal_iservice_proto_goTypes = nil
	file_internalapi_sensor_signal_iservice_proto_depIdxs = nil
}
