// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.3.1-0.20220817155510-0ae748fd2007
// source: api/v1/auth_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	storage "github.com/stackrox/rox/generated/storage"
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

func (m *UserAttribute) CloneVT() *UserAttribute {
	if m == nil {
		return (*UserAttribute)(nil)
	}
	r := &UserAttribute{
		Key: m.Key,
	}
	if rhs := m.Values; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.Values = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *UserAttribute) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (m *AuthStatus) CloneVT() *AuthStatus {
	if m == nil {
		return (*AuthStatus)(nil)
	}
	r := &AuthStatus{
		RefreshUrl: m.RefreshUrl,
	}
	if m.Id != nil {
		r.Id = m.Id.(interface{ CloneVT() isAuthStatus_Id }).CloneVT()
	}
	if rhs := m.Expires; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *timestamppb.Timestamp }); ok {
			r.Expires = vtpb.CloneVT()
		} else {
			r.Expires = proto.Clone(rhs).(*timestamppb.Timestamp)
		}
	}
	if rhs := m.AuthProvider; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *storage.AuthProvider }); ok {
			r.AuthProvider = vtpb.CloneVT()
		} else {
			r.AuthProvider = proto.Clone(rhs).(*storage.AuthProvider)
		}
	}
	if rhs := m.UserInfo; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface{ CloneVT() *storage.UserInfo }); ok {
			r.UserInfo = vtpb.CloneVT()
		} else {
			r.UserInfo = proto.Clone(rhs).(*storage.UserInfo)
		}
	}
	if rhs := m.UserAttributes; rhs != nil {
		tmpContainer := make([]*UserAttribute, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.UserAttributes = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *AuthStatus) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (m *AuthStatus_UserId) CloneVT() isAuthStatus_Id {
	if m == nil {
		return (*AuthStatus_UserId)(nil)
	}
	r := &AuthStatus_UserId{
		UserId: m.UserId,
	}
	return r
}

func (m *AuthStatus_ServiceId) CloneVT() isAuthStatus_Id {
	if m == nil {
		return (*AuthStatus_ServiceId)(nil)
	}
	r := &AuthStatus_ServiceId{}
	if rhs := m.ServiceId; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface {
			CloneVT() *storage.ServiceIdentity
		}); ok {
			r.ServiceId = vtpb.CloneVT()
		} else {
			r.ServiceId = proto.Clone(rhs).(*storage.ServiceIdentity)
		}
	}
	return r
}

func (this *UserAttribute) EqualVT(that *UserAttribute) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Key != that.Key {
		return false
	}
	if len(this.Values) != len(that.Values) {
		return false
	}
	for i, vx := range this.Values {
		vy := that.Values[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AuthStatus) EqualVT(that *AuthStatus) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Id == nil && that.Id != nil {
		return false
	} else if this.Id != nil {
		if that.Id == nil {
			return false
		}
		if !this.Id.(interface{ EqualVT(isAuthStatus_Id) bool }).EqualVT(that.Id) {
			return false
		}
	}
	if equal, ok := interface{}(this.Expires).(interface {
		EqualVT(*timestamppb.Timestamp) bool
	}); ok {
		if !equal.EqualVT(that.Expires) {
			return false
		}
	} else if !proto.Equal(this.Expires, that.Expires) {
		return false
	}
	if this.RefreshUrl != that.RefreshUrl {
		return false
	}
	if equal, ok := interface{}(this.AuthProvider).(interface {
		EqualVT(*storage.AuthProvider) bool
	}); ok {
		if !equal.EqualVT(that.AuthProvider) {
			return false
		}
	} else if !proto.Equal(this.AuthProvider, that.AuthProvider) {
		return false
	}
	if equal, ok := interface{}(this.UserInfo).(interface{ EqualVT(*storage.UserInfo) bool }); ok {
		if !equal.EqualVT(that.UserInfo) {
			return false
		}
	} else if !proto.Equal(this.UserInfo, that.UserInfo) {
		return false
	}
	if len(this.UserAttributes) != len(that.UserAttributes) {
		return false
	}
	for i, vx := range this.UserAttributes {
		vy := that.UserAttributes[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &UserAttribute{}
			}
			if q == nil {
				q = &UserAttribute{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AuthStatus_UserId) EqualVT(thatIface isAuthStatus_Id) bool {
	that, ok := thatIface.(*AuthStatus_UserId)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.UserId != that.UserId {
		return false
	}
	return true
}

func (this *AuthStatus_ServiceId) EqualVT(thatIface isAuthStatus_Id) bool {
	that, ok := thatIface.(*AuthStatus_ServiceId)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.ServiceId, that.ServiceId; p != q {
		if p == nil {
			p = &storage.ServiceIdentity{}
		}
		if q == nil {
			q = &storage.ServiceIdentity{}
		}
		if equal, ok := interface{}(p).(interface {
			EqualVT(*storage.ServiceIdentity) bool
		}); ok {
			if !equal.EqualVT(q) {
				return false
			}
		} else if !proto.Equal(p, q) {
			return false
		}
	}
	return true
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	GetAuthStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthStatus, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GetAuthStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthStatus, error) {
	out := new(AuthStatus)
	err := c.cc.Invoke(ctx, "/v1.AuthService/GetAuthStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	GetAuthStatus(context.Context, *Empty) (*AuthStatus, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) GetAuthStatus(context.Context, *Empty) (*AuthStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthStatus not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_GetAuthStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAuthStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AuthService/GetAuthStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAuthStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthStatus",
			Handler:    _AuthService_GetAuthStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/auth_service.proto",
}

func (m *UserAttribute) MarshalVT() (dAtA []byte, err error) {
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

func (m *UserAttribute) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *UserAttribute) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Values[iNdEx])
			copy(dAtA[i:], m.Values[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.Values[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarint(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthStatus) MarshalVT() (dAtA []byte, err error) {
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

func (m *AuthStatus) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *AuthStatus) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.Id.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.UserAttributes) > 0 {
		for iNdEx := len(m.UserAttributes) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.UserAttributes[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.UserInfo != nil {
		if vtmsg, ok := interface{}(m.UserInfo).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.UserInfo)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.AuthProvider != nil {
		if vtmsg, ok := interface{}(m.AuthProvider).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.AuthProvider)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RefreshUrl) > 0 {
		i -= len(m.RefreshUrl)
		copy(dAtA[i:], m.RefreshUrl)
		i = encodeVarint(dAtA, i, uint64(len(m.RefreshUrl)))
		i--
		dAtA[i] = 0x22
	}
	if m.Expires != nil {
		if vtmsg, ok := interface{}(m.Expires).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Expires)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func (m *AuthStatus_UserId) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *AuthStatus_UserId) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.UserId)
	copy(dAtA[i:], m.UserId)
	i = encodeVarint(dAtA, i, uint64(len(m.UserId)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *AuthStatus_ServiceId) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *AuthStatus_ServiceId) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ServiceId != nil {
		if vtmsg, ok := interface{}(m.ServiceId).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ServiceId)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *UserAttribute) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *AuthStatus) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Id.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.Expires != nil {
		if size, ok := interface{}(m.Expires).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Expires)
		}
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.RefreshUrl)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.AuthProvider != nil {
		if size, ok := interface{}(m.AuthProvider).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.AuthProvider)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.UserInfo != nil {
		if size, ok := interface{}(m.UserInfo).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.UserInfo)
		}
		n += 1 + l + sov(uint64(l))
	}
	if len(m.UserAttributes) > 0 {
		for _, e := range m.UserAttributes {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *AuthStatus_UserId) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *AuthStatus_ServiceId) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServiceId != nil {
		if size, ok := interface{}(m.ServiceId).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ServiceId)
		}
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *UserAttribute) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: UserAttribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserAttribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
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
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
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
func (m *AuthStatus) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: AuthStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
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
			m.Id = &AuthStatus_UserId{UserId: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
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
			if oneof, ok := m.Id.(*AuthStatus_ServiceId); ok {
				if unmarshal, ok := interface{}(oneof.ServiceId).(interface {
					UnmarshalVT([]byte) error
				}); ok {
					if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
						return err
					}
				} else {
					if err := proto.Unmarshal(dAtA[iNdEx:postIndex], oneof.ServiceId); err != nil {
						return err
					}
				}
			} else {
				v := &storage.ServiceIdentity{}
				if unmarshal, ok := interface{}(v).(interface {
					UnmarshalVT([]byte) error
				}); ok {
					if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
						return err
					}
				} else {
					if err := proto.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
						return err
					}
				}
				m.Id = &AuthStatus_ServiceId{ServiceId: v}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
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
			if m.Expires == nil {
				m.Expires = &timestamppb.Timestamp{}
			}
			if unmarshal, ok := interface{}(m.Expires).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.Expires); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefreshUrl", wireType)
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
			m.RefreshUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthProvider", wireType)
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
			if m.AuthProvider == nil {
				m.AuthProvider = &storage.AuthProvider{}
			}
			if unmarshal, ok := interface{}(m.AuthProvider).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.AuthProvider); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserInfo", wireType)
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
			if m.UserInfo == nil {
				m.UserInfo = &storage.UserInfo{}
			}
			if unmarshal, ok := interface{}(m.UserInfo).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.UserInfo); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAttributes", wireType)
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
			m.UserAttributes = append(m.UserAttributes, &UserAttribute{})
			if err := m.UserAttributes[len(m.UserAttributes)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
