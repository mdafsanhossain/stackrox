// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.3.1-0.20220817155510-0ae748fd2007
// source: internalapi/sensor/cert_distribution_iservice.proto

package sensor

import (
	context "context"
	fmt "fmt"
	storage "github.com/stackrox/rox/generated/storage"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *FetchCertificateRequest) CloneVT() *FetchCertificateRequest {
	if m == nil {
		return (*FetchCertificateRequest)(nil)
	}
	r := &FetchCertificateRequest{
		ServiceType:         m.ServiceType,
		ServiceAccountToken: m.ServiceAccountToken,
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *FetchCertificateRequest) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (m *FetchCertificateResponse) CloneVT() *FetchCertificateResponse {
	if m == nil {
		return (*FetchCertificateResponse)(nil)
	}
	r := &FetchCertificateResponse{
		PemCert: m.PemCert,
		PemKey:  m.PemKey,
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *FetchCertificateResponse) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (this *FetchCertificateRequest) EqualVT(that *FetchCertificateRequest) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.ServiceType != that.ServiceType {
		return false
	}
	if this.ServiceAccountToken != that.ServiceAccountToken {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *FetchCertificateResponse) EqualVT(that *FetchCertificateResponse) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.PemCert != that.PemCert {
		return false
	}
	if this.PemKey != that.PemKey {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CertDistributionServiceClient is the client API for CertDistributionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertDistributionServiceClient interface {
	FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error)
}

type certDistributionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertDistributionServiceClient(cc grpc.ClientConnInterface) CertDistributionServiceClient {
	return &certDistributionServiceClient{cc}
}

func (c *certDistributionServiceClient) FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error) {
	out := new(FetchCertificateResponse)
	err := c.cc.Invoke(ctx, "/sensor.CertDistributionService/FetchCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertDistributionServiceServer is the server API for CertDistributionService service.
// All implementations must embed UnimplementedCertDistributionServiceServer
// for forward compatibility
type CertDistributionServiceServer interface {
	FetchCertificate(context.Context, *FetchCertificateRequest) (*FetchCertificateResponse, error)
	mustEmbedUnimplementedCertDistributionServiceServer()
}

// UnimplementedCertDistributionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCertDistributionServiceServer struct {
}

func (UnimplementedCertDistributionServiceServer) FetchCertificate(context.Context, *FetchCertificateRequest) (*FetchCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCertificate not implemented")
}
func (UnimplementedCertDistributionServiceServer) mustEmbedUnimplementedCertDistributionServiceServer() {
}

// UnsafeCertDistributionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertDistributionServiceServer will
// result in compilation errors.
type UnsafeCertDistributionServiceServer interface {
	mustEmbedUnimplementedCertDistributionServiceServer()
}

func RegisterCertDistributionServiceServer(s grpc.ServiceRegistrar, srv CertDistributionServiceServer) {
	s.RegisterService(&CertDistributionService_ServiceDesc, srv)
}

func _CertDistributionService_FetchCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertDistributionServiceServer).FetchCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensor.CertDistributionService/FetchCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertDistributionServiceServer).FetchCertificate(ctx, req.(*FetchCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertDistributionService_ServiceDesc is the grpc.ServiceDesc for CertDistributionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertDistributionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensor.CertDistributionService",
	HandlerType: (*CertDistributionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchCertificate",
			Handler:    _CertDistributionService_FetchCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalapi/sensor/cert_distribution_iservice.proto",
}

func (m *FetchCertificateRequest) MarshalVT() (dAtA []byte, err error) {
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

func (m *FetchCertificateRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *FetchCertificateRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.ServiceAccountToken) > 0 {
		i -= len(m.ServiceAccountToken)
		copy(dAtA[i:], m.ServiceAccountToken)
		i = encodeVarint(dAtA, i, uint64(len(m.ServiceAccountToken)))
		i--
		dAtA[i] = 0x12
	}
	if m.ServiceType != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ServiceType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FetchCertificateResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *FetchCertificateResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *FetchCertificateResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.PemKey) > 0 {
		i -= len(m.PemKey)
		copy(dAtA[i:], m.PemKey)
		i = encodeVarint(dAtA, i, uint64(len(m.PemKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PemCert) > 0 {
		i -= len(m.PemCert)
		copy(dAtA[i:], m.PemCert)
		i = encodeVarint(dAtA, i, uint64(len(m.PemCert)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FetchCertificateRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServiceType != 0 {
		n += 1 + sov(uint64(m.ServiceType))
	}
	l = len(m.ServiceAccountToken)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *FetchCertificateResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PemCert)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.PemKey)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *FetchCertificateRequest) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: FetchCertificateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchCertificateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceType", wireType)
			}
			m.ServiceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServiceType |= storage.ServiceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAccountToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceAccountToken = string(dAtA[iNdEx:postIndex])
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
func (m *FetchCertificateResponse) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: FetchCertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchCertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PemCert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PemCert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PemKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PemKey = string(dAtA[iNdEx:postIndex])
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
