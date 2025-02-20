// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/sensor/admission_control_iservice.proto

package sensor

import (
	context "context"
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	central "github.com/stackrox/rox/generated/internalapi/central"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MsgFromAdmissionControl struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgFromAdmissionControl) Reset()         { *m = MsgFromAdmissionControl{} }
func (m *MsgFromAdmissionControl) String() string { return proto.CompactTextString(m) }
func (*MsgFromAdmissionControl) ProtoMessage()    {}
func (*MsgFromAdmissionControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_861e2dad9d4a82c2, []int{0}
}
func (m *MsgFromAdmissionControl) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFromAdmissionControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFromAdmissionControl.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFromAdmissionControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFromAdmissionControl.Merge(m, src)
}
func (m *MsgFromAdmissionControl) XXX_Size() int {
	return m.Size()
}
func (m *MsgFromAdmissionControl) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFromAdmissionControl.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFromAdmissionControl proto.InternalMessageInfo

func (m *MsgFromAdmissionControl) MessageClone() proto.Message {
	return m.Clone()
}
func (m *MsgFromAdmissionControl) Clone() *MsgFromAdmissionControl {
	if m == nil {
		return nil
	}
	cloned := new(MsgFromAdmissionControl)
	*cloned = *m

	return cloned
}

type MsgToAdmissionControl struct {
	// Types that are valid to be assigned to Msg:
	//	*MsgToAdmissionControl_SettingsPush
	//	*MsgToAdmissionControl_UpdateResourceRequest
	//	*MsgToAdmissionControl_TestNotifierRequest
	Msg                  isMsgToAdmissionControl_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MsgToAdmissionControl) Reset()         { *m = MsgToAdmissionControl{} }
func (m *MsgToAdmissionControl) String() string { return proto.CompactTextString(m) }
func (*MsgToAdmissionControl) ProtoMessage()    {}
func (*MsgToAdmissionControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_861e2dad9d4a82c2, []int{1}
}
func (m *MsgToAdmissionControl) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgToAdmissionControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgToAdmissionControl.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgToAdmissionControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgToAdmissionControl.Merge(m, src)
}
func (m *MsgToAdmissionControl) XXX_Size() int {
	return m.Size()
}
func (m *MsgToAdmissionControl) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgToAdmissionControl.DiscardUnknown(m)
}

var xxx_messageInfo_MsgToAdmissionControl proto.InternalMessageInfo

type isMsgToAdmissionControl_Msg interface {
	isMsgToAdmissionControl_Msg()
	MarshalTo([]byte) (int, error)
	Size() int
	Clone() isMsgToAdmissionControl_Msg
}

type MsgToAdmissionControl_SettingsPush struct {
	SettingsPush *AdmissionControlSettings `protobuf:"bytes,1,opt,name=settings_push,json=settingsPush,proto3,oneof" json:"settings_push,omitempty"`
}
type MsgToAdmissionControl_UpdateResourceRequest struct {
	UpdateResourceRequest *AdmCtrlUpdateResourceRequest `protobuf:"bytes,2,opt,name=update_resource_request,json=updateResourceRequest,proto3,oneof" json:"update_resource_request,omitempty"`
}
type MsgToAdmissionControl_TestNotifierRequest struct {
	TestNotifierRequest *central.TestNotifierRequest `protobuf:"bytes,3,opt,name=test_notifier_request,json=testNotifierRequest,proto3,oneof" json:"test_notifier_request,omitempty"`
}

func (*MsgToAdmissionControl_SettingsPush) isMsgToAdmissionControl_Msg() {}
func (m *MsgToAdmissionControl_SettingsPush) Clone() isMsgToAdmissionControl_Msg {
	if m == nil {
		return nil
	}
	cloned := new(MsgToAdmissionControl_SettingsPush)
	*cloned = *m

	cloned.SettingsPush = m.SettingsPush.Clone()
	return cloned
}
func (*MsgToAdmissionControl_UpdateResourceRequest) isMsgToAdmissionControl_Msg() {}
func (m *MsgToAdmissionControl_UpdateResourceRequest) Clone() isMsgToAdmissionControl_Msg {
	if m == nil {
		return nil
	}
	cloned := new(MsgToAdmissionControl_UpdateResourceRequest)
	*cloned = *m

	cloned.UpdateResourceRequest = m.UpdateResourceRequest.Clone()
	return cloned
}
func (*MsgToAdmissionControl_TestNotifierRequest) isMsgToAdmissionControl_Msg() {}
func (m *MsgToAdmissionControl_TestNotifierRequest) Clone() isMsgToAdmissionControl_Msg {
	if m == nil {
		return nil
	}
	cloned := new(MsgToAdmissionControl_TestNotifierRequest)
	*cloned = *m

	cloned.TestNotifierRequest = m.TestNotifierRequest.Clone()
	return cloned
}

func (m *MsgToAdmissionControl) GetMsg() isMsgToAdmissionControl_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *MsgToAdmissionControl) GetSettingsPush() *AdmissionControlSettings {
	if x, ok := m.GetMsg().(*MsgToAdmissionControl_SettingsPush); ok {
		return x.SettingsPush
	}
	return nil
}

func (m *MsgToAdmissionControl) GetUpdateResourceRequest() *AdmCtrlUpdateResourceRequest {
	if x, ok := m.GetMsg().(*MsgToAdmissionControl_UpdateResourceRequest); ok {
		return x.UpdateResourceRequest
	}
	return nil
}

func (m *MsgToAdmissionControl) GetTestNotifierRequest() *central.TestNotifierRequest {
	if x, ok := m.GetMsg().(*MsgToAdmissionControl_TestNotifierRequest); ok {
		return x.TestNotifierRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MsgToAdmissionControl) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MsgToAdmissionControl_SettingsPush)(nil),
		(*MsgToAdmissionControl_UpdateResourceRequest)(nil),
		(*MsgToAdmissionControl_TestNotifierRequest)(nil),
	}
}

func (m *MsgToAdmissionControl) MessageClone() proto.Message {
	return m.Clone()
}
func (m *MsgToAdmissionControl) Clone() *MsgToAdmissionControl {
	if m == nil {
		return nil
	}
	cloned := new(MsgToAdmissionControl)
	*cloned = *m

	if m.Msg != nil {
		cloned.Msg = m.Msg.Clone()
	}
	return cloned
}

func init() {
	proto.RegisterType((*MsgFromAdmissionControl)(nil), "sensor.MsgFromAdmissionControl")
	proto.RegisterType((*MsgToAdmissionControl)(nil), "sensor.MsgToAdmissionControl")
}

func init() {
	proto.RegisterFile("internalapi/sensor/admission_control_iservice.proto", fileDescriptor_861e2dad9d4a82c2)
}

var fileDescriptor_861e2dad9d4a82c2 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0xc7, 0xad, 0xe8, 0xc2, 0x2d, 0x1b, 0xa3, 0xa1, 0xc3, 0x00, 0xa1, 0x54, 0x2c, 0x2a,
	0x16, 0x0e, 0x9a, 0x9e, 0xa0, 0x1d, 0x51, 0xba, 0x19, 0x54, 0xd2, 0xb2, 0x61, 0x41, 0xe4, 0x7a,
	0x5e, 0x53, 0x4b, 0x8e, 0x1d, 0xfc, 0x9e, 0x91, 0x7a, 0x13, 0x24, 0xee, 0xc2, 0x9a, 0x25, 0x47,
	0x40, 0xc3, 0x45, 0x50, 0xe3, 0xa4, 0x8c, 0xa6, 0x54, 0xea, 0x2e, 0xc9, 0xfb, 0xfe, 0xdf, 0x79,
	0xff, 0x6f, 0xbe, 0x6f, 0x1c, 0x41, 0x70, 0xca, 0xaa, 0xc6, 0xe4, 0x08, 0x0e, 0x7d, 0xc8, 0xd5,
	0xbc, 0x36, 0x88, 0xc6, 0xbb, 0x52, 0x7b, 0x47, 0xc1, 0xdb, 0xd2, 0x20, 0x84, 0xaf, 0x46, 0x83,
	0x6c, 0x82, 0x27, 0x2f, 0x36, 0x12, 0x38, 0x7e, 0x7d, 0x1f, 0x71, 0xd2, 0x8c, 0x27, 0xcb, 0xac,
	0x06, 0x47, 0x41, 0xd9, 0x1c, 0x41, 0xc7, 0x00, 0xf3, 0x52, 0xdb, 0x88, 0x04, 0xa1, 0x74, 0x9e,
	0xcc, 0x85, 0x81, 0xd0, 0x69, 0x9e, 0x56, 0xde, 0x57, 0x16, 0xf2, 0xf6, 0xed, 0x3c, 0x5e, 0xe4,
	0x50, 0x37, 0x74, 0x95, 0x86, 0xbb, 0x4f, 0xf8, 0xf6, 0x0c, 0xab, 0xa3, 0xe0, 0xeb, 0x83, 0xfe,
	0xc8, 0x69, 0x3a, 0x71, 0xf7, 0xfb, 0x1a, 0x1f, 0xce, 0xb0, 0x3a, 0xf3, 0xab, 0x13, 0xf1, 0x8e,
	0x3f, 0x44, 0x20, 0x32, 0xae, 0xc2, 0xb2, 0x89, 0x78, 0x39, 0x62, 0x3b, 0x6c, 0x6f, 0x73, 0xb2,
	0x23, 0xd3, 0xdf, 0xcb, 0x55, 0xc1, 0x69, 0x07, 0x1f, 0x0f, 0x8a, 0xad, 0x5e, 0x78, 0x12, 0xf1,
	0x52, 0x7c, 0xe6, 0xdb, 0xb1, 0x99, 0x2b, 0x82, 0x32, 0x00, 0xfa, 0x18, 0xf4, 0xf5, 0xc3, 0x97,
	0x08, 0x48, 0xa3, 0xb5, 0xd6, 0xf2, 0xd5, 0x92, 0xe5, 0x94, 0x82, 0xfd, 0xd8, 0xd2, 0x45, 0x07,
	0x17, 0x89, 0x3d, 0x1e, 0x14, 0xc3, 0xf8, 0xbf, 0x81, 0x28, 0xf8, 0x90, 0x00, 0xe9, 0x26, 0x91,
	0x1b, 0xf7, 0xf5, 0xd6, 0xfd, 0x99, 0xec, 0x22, 0x94, 0x67, 0x80, 0xf4, 0xbe, 0x83, 0xfe, 0xb9,
	0x3e, 0xa2, 0xdb, 0x9f, 0x0f, 0x1f, 0xf0, 0xf5, 0x1a, 0xab, 0xc9, 0x0f, 0xc6, 0x5f, 0xae, 0xee,
	0x39, 0x53, 0x4e, 0x55, 0x50, 0x83, 0xa3, 0xd3, 0xd4, 0xb4, 0xf8, 0xc0, 0x37, 0xa7, 0xbe, 0xae,
	0xa3, 0x33, 0x5a, 0x11, 0x88, 0x17, 0xfd, 0x3a, 0x77, 0x64, 0x3e, 0x7e, 0xbe, 0x04, 0xdc, 0x0e,
	0x7e, 0x8f, 0xbd, 0x61, 0xe2, 0x88, 0x6f, 0x9d, 0x78, 0x6b, 0xf4, 0xd5, 0x81, 0x85, 0x40, 0x28,
	0xb2, 0xbb, 0x52, 0x4f, 0xf3, 0xf1, 0x63, 0x99, 0xfa, 0x97, 0x7d, 0xff, 0xf2, 0xed, 0x75, 0xff,
	0x87, 0xa3, 0x9f, 0x8b, 0x8c, 0xfd, 0x5a, 0x64, 0xec, 0xf7, 0x22, 0x63, 0xdf, 0xfe, 0x64, 0x83,
	0x4f, 0xdd, 0x85, 0x3c, 0xdf, 0x68, 0xc9, 0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x42,
	0x20, 0xd2, 0xd6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdmissionControlManagementServiceClient is the client API for AdmissionControlManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type AdmissionControlManagementServiceClient interface {
	Communicate(ctx context.Context, opts ...grpc.CallOption) (AdmissionControlManagementService_CommunicateClient, error)
	PolicyAlerts(ctx context.Context, in *AdmissionControlAlerts, opts ...grpc.CallOption) (*types.Empty, error)
}

type admissionControlManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdmissionControlManagementServiceClient(cc grpc.ClientConnInterface) AdmissionControlManagementServiceClient {
	return &admissionControlManagementServiceClient{cc}
}

func (c *admissionControlManagementServiceClient) Communicate(ctx context.Context, opts ...grpc.CallOption) (AdmissionControlManagementService_CommunicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AdmissionControlManagementService_serviceDesc.Streams[0], "/sensor.AdmissionControlManagementService/Communicate", opts...)
	if err != nil {
		return nil, err
	}
	x := &admissionControlManagementServiceCommunicateClient{stream}
	return x, nil
}

type AdmissionControlManagementService_CommunicateClient interface {
	Send(*MsgFromAdmissionControl) error
	Recv() (*MsgToAdmissionControl, error)
	grpc.ClientStream
}

type admissionControlManagementServiceCommunicateClient struct {
	grpc.ClientStream
}

func (x *admissionControlManagementServiceCommunicateClient) Send(m *MsgFromAdmissionControl) error {
	return x.ClientStream.SendMsg(m)
}

func (x *admissionControlManagementServiceCommunicateClient) Recv() (*MsgToAdmissionControl, error) {
	m := new(MsgToAdmissionControl)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *admissionControlManagementServiceClient) PolicyAlerts(ctx context.Context, in *AdmissionControlAlerts, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/sensor.AdmissionControlManagementService/PolicyAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdmissionControlManagementServiceServer is the server API for AdmissionControlManagementService service.
type AdmissionControlManagementServiceServer interface {
	Communicate(AdmissionControlManagementService_CommunicateServer) error
	PolicyAlerts(context.Context, *AdmissionControlAlerts) (*types.Empty, error)
}

// UnimplementedAdmissionControlManagementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdmissionControlManagementServiceServer struct {
}

func (*UnimplementedAdmissionControlManagementServiceServer) Communicate(srv AdmissionControlManagementService_CommunicateServer) error {
	return status.Errorf(codes.Unimplemented, "method Communicate not implemented")
}
func (*UnimplementedAdmissionControlManagementServiceServer) PolicyAlerts(ctx context.Context, req *AdmissionControlAlerts) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyAlerts not implemented")
}

func RegisterAdmissionControlManagementServiceServer(s *grpc.Server, srv AdmissionControlManagementServiceServer) {
	s.RegisterService(&_AdmissionControlManagementService_serviceDesc, srv)
}

func _AdmissionControlManagementService_Communicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AdmissionControlManagementServiceServer).Communicate(&admissionControlManagementServiceCommunicateServer{stream})
}

type AdmissionControlManagementService_CommunicateServer interface {
	Send(*MsgToAdmissionControl) error
	Recv() (*MsgFromAdmissionControl, error)
	grpc.ServerStream
}

type admissionControlManagementServiceCommunicateServer struct {
	grpc.ServerStream
}

func (x *admissionControlManagementServiceCommunicateServer) Send(m *MsgToAdmissionControl) error {
	return x.ServerStream.SendMsg(m)
}

func (x *admissionControlManagementServiceCommunicateServer) Recv() (*MsgFromAdmissionControl, error) {
	m := new(MsgFromAdmissionControl)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AdmissionControlManagementService_PolicyAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdmissionControlAlerts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdmissionControlManagementServiceServer).PolicyAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensor.AdmissionControlManagementService/PolicyAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdmissionControlManagementServiceServer).PolicyAlerts(ctx, req.(*AdmissionControlAlerts))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdmissionControlManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sensor.AdmissionControlManagementService",
	HandlerType: (*AdmissionControlManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PolicyAlerts",
			Handler:    _AdmissionControlManagementService_PolicyAlerts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Communicate",
			Handler:       _AdmissionControlManagementService_Communicate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internalapi/sensor/admission_control_iservice.proto",
}

func (m *MsgFromAdmissionControl) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFromAdmissionControl) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFromAdmissionControl) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *MsgToAdmissionControl) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgToAdmissionControl) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgToAdmissionControl) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Msg != nil {
		{
			size := m.Msg.Size()
			i -= size
			if _, err := m.Msg.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgToAdmissionControl_SettingsPush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgToAdmissionControl_SettingsPush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SettingsPush != nil {
		{
			size, err := m.SettingsPush.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdmissionControlIservice(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *MsgToAdmissionControl_UpdateResourceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgToAdmissionControl_UpdateResourceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.UpdateResourceRequest != nil {
		{
			size, err := m.UpdateResourceRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdmissionControlIservice(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *MsgToAdmissionControl_TestNotifierRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgToAdmissionControl_TestNotifierRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TestNotifierRequest != nil {
		{
			size, err := m.TestNotifierRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdmissionControlIservice(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func encodeVarintAdmissionControlIservice(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdmissionControlIservice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgFromAdmissionControl) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MsgToAdmissionControl) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Msg != nil {
		n += m.Msg.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MsgToAdmissionControl_SettingsPush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SettingsPush != nil {
		l = m.SettingsPush.Size()
		n += 1 + l + sovAdmissionControlIservice(uint64(l))
	}
	return n
}
func (m *MsgToAdmissionControl_UpdateResourceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpdateResourceRequest != nil {
		l = m.UpdateResourceRequest.Size()
		n += 1 + l + sovAdmissionControlIservice(uint64(l))
	}
	return n
}
func (m *MsgToAdmissionControl_TestNotifierRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TestNotifierRequest != nil {
		l = m.TestNotifierRequest.Size()
		n += 1 + l + sovAdmissionControlIservice(uint64(l))
	}
	return n
}

func sovAdmissionControlIservice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdmissionControlIservice(x uint64) (n int) {
	return sovAdmissionControlIservice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgFromAdmissionControl) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmissionControlIservice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgFromAdmissionControl: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFromAdmissionControl: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAdmissionControlIservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmissionControlIservice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgToAdmissionControl) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmissionControlIservice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgToAdmissionControl: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgToAdmissionControl: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettingsPush", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmissionControlIservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAdmissionControlIservice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmissionControlIservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AdmissionControlSettings{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &MsgToAdmissionControl_SettingsPush{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateResourceRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmissionControlIservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAdmissionControlIservice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmissionControlIservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AdmCtrlUpdateResourceRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &MsgToAdmissionControl_UpdateResourceRequest{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestNotifierRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmissionControlIservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAdmissionControlIservice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmissionControlIservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &central.TestNotifierRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Msg = &MsgToAdmissionControl_TestNotifierRequest{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmissionControlIservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmissionControlIservice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAdmissionControlIservice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdmissionControlIservice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAdmissionControlIservice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAdmissionControlIservice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAdmissionControlIservice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdmissionControlIservice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdmissionControlIservice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdmissionControlIservice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdmissionControlIservice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdmissionControlIservice = fmt.Errorf("proto: unexpected end of group")
)
