// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.3.1-0.20220817155510-0ae748fd2007
// source: api/v1/credential_expiry_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *GetCertExpiry_Request) CloneVT() *GetCertExpiry_Request {
	if m == nil {
		return (*GetCertExpiry_Request)(nil)
	}
	r := &GetCertExpiry_Request{
		Component: m.Component,
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetCertExpiry_Request) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (m *GetCertExpiry_Response) CloneVT() *GetCertExpiry_Response {
	if m == nil {
		return (*GetCertExpiry_Response)(nil)
	}
	r := &GetCertExpiry_Response{}
	if rhs := m.Expiry; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *timestamppb.Timestamp }); ok {
			r.Expiry = vtpb.CloneVT()
		} else {
			r.Expiry = proto.Clone(rhs).(*timestamppb.Timestamp)
		}
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetCertExpiry_Response) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (m *GetCertExpiry) CloneVT() *GetCertExpiry {
	if m == nil {
		return (*GetCertExpiry)(nil)
	}
	r := &GetCertExpiry{}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetCertExpiry) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (this *GetCertExpiry_Request) EqualVT(that *GetCertExpiry_Request) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Component != that.Component {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetCertExpiry_Response) EqualVT(that *GetCertExpiry_Response) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if equal, ok := interface{}(this.Expiry).(interface {
		EqualVT(*timestamppb.Timestamp) bool
	}); ok {
		if !equal.EqualVT(that.Expiry) {
			return false
		}
	} else if !proto.Equal(this.Expiry, that.Expiry) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetCertExpiry) EqualVT(that *GetCertExpiry) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CredentialExpiryServiceClient is the client API for CredentialExpiryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CredentialExpiryServiceClient interface {
	// GetCertExpiry returns information related to the expiry component mTLS certificate.
	GetCertExpiry(ctx context.Context, in *GetCertExpiry_Request, opts ...grpc.CallOption) (*GetCertExpiry_Response, error)
}

type credentialExpiryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCredentialExpiryServiceClient(cc grpc.ClientConnInterface) CredentialExpiryServiceClient {
	return &credentialExpiryServiceClient{cc}
}

func (c *credentialExpiryServiceClient) GetCertExpiry(ctx context.Context, in *GetCertExpiry_Request, opts ...grpc.CallOption) (*GetCertExpiry_Response, error) {
	out := new(GetCertExpiry_Response)
	err := c.cc.Invoke(ctx, "/v1.CredentialExpiryService/GetCertExpiry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialExpiryServiceServer is the server API for CredentialExpiryService service.
// All implementations must embed UnimplementedCredentialExpiryServiceServer
// for forward compatibility
type CredentialExpiryServiceServer interface {
	// GetCertExpiry returns information related to the expiry component mTLS certificate.
	GetCertExpiry(context.Context, *GetCertExpiry_Request) (*GetCertExpiry_Response, error)
	mustEmbedUnimplementedCredentialExpiryServiceServer()
}

// UnimplementedCredentialExpiryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCredentialExpiryServiceServer struct {
}

func (UnimplementedCredentialExpiryServiceServer) GetCertExpiry(context.Context, *GetCertExpiry_Request) (*GetCertExpiry_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertExpiry not implemented")
}
func (UnimplementedCredentialExpiryServiceServer) mustEmbedUnimplementedCredentialExpiryServiceServer() {
}

// UnsafeCredentialExpiryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CredentialExpiryServiceServer will
// result in compilation errors.
type UnsafeCredentialExpiryServiceServer interface {
	mustEmbedUnimplementedCredentialExpiryServiceServer()
}

func RegisterCredentialExpiryServiceServer(s grpc.ServiceRegistrar, srv CredentialExpiryServiceServer) {
	s.RegisterService(&CredentialExpiryService_ServiceDesc, srv)
}

func _CredentialExpiryService_GetCertExpiry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertExpiry_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialExpiryServiceServer).GetCertExpiry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CredentialExpiryService/GetCertExpiry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialExpiryServiceServer).GetCertExpiry(ctx, req.(*GetCertExpiry_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CredentialExpiryService_ServiceDesc is the grpc.ServiceDesc for CredentialExpiryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CredentialExpiryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CredentialExpiryService",
	HandlerType: (*CredentialExpiryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCertExpiry",
			Handler:    _CredentialExpiryService_GetCertExpiry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/credential_expiry_service.proto",
}

func (m *GetCertExpiry_Request) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCertExpiry_Request) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *GetCertExpiry_Request) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Component != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Component))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetCertExpiry_Response) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCertExpiry_Response) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *GetCertExpiry_Response) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Expiry != nil {
		if vtmsg, ok := interface{}(m.Expiry).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Expiry)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetCertExpiry) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCertExpiry) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *GetCertExpiry) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	return len(dAtA) - i, nil
}

func (m *GetCertExpiry_Request) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Component != 0 {
		n += 1 + sov(uint64(m.Component))
	}
	n += len(m.unknownFields)
	return n
}

func (m *GetCertExpiry_Response) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Expiry != nil {
		if size, ok := interface{}(m.Expiry).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Expiry)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *GetCertExpiry) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *GetCertExpiry_Request) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: GetCertExpiry_Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCertExpiry_Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Component", wireType)
			}
			m.Component = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Component |= GetCertExpiry_Component(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetCertExpiry_Response) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: GetCertExpiry_Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCertExpiry_Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiry == nil {
				m.Expiry = &timestamppb.Timestamp{}
			}
			if unmarshal, ok := interface{}(m.Expiry).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.Expiry); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetCertExpiry) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: GetCertExpiry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCertExpiry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
