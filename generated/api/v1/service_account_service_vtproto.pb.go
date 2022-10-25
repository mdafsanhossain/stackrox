// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.3.1-0.20220817155510-0ae748fd2007
// source: api/v1/service_account_service.proto

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
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ListServiceAccountResponse) CloneVT() *ListServiceAccountResponse {
	if m == nil {
		return (*ListServiceAccountResponse)(nil)
	}
	r := &ListServiceAccountResponse{}
	if rhs := m.SaAndRoles; rhs != nil {
		tmpContainer := make([]*ServiceAccountAndRoles, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.SaAndRoles = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListServiceAccountResponse) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (m *ServiceAccountAndRoles) CloneVT() *ServiceAccountAndRoles {
	if m == nil {
		return (*ServiceAccountAndRoles)(nil)
	}
	r := &ServiceAccountAndRoles{}
	if rhs := m.ServiceAccount; rhs != nil {
		if vtpb, ok := interface{}(rhs).(interface {
			CloneVT() *storage.ServiceAccount
		}); ok {
			r.ServiceAccount = vtpb.CloneVT()
		} else {
			r.ServiceAccount = proto.Clone(rhs).(*storage.ServiceAccount)
		}
	}
	if rhs := m.ClusterRoles; rhs != nil {
		tmpContainer := make([]*storage.K8SRole, len(rhs))
		for k, v := range rhs {
			if vtpb, ok := interface{}(v).(interface{ CloneVT() *storage.K8SRole }); ok {
				tmpContainer[k] = vtpb.CloneVT()
			} else {
				tmpContainer[k] = proto.Clone(v).(*storage.K8SRole)
			}
		}
		r.ClusterRoles = tmpContainer
	}
	if rhs := m.ScopedRoles; rhs != nil {
		tmpContainer := make([]*ScopedRoles, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ScopedRoles = tmpContainer
	}
	if rhs := m.DeploymentRelationships; rhs != nil {
		tmpContainer := make([]*SADeploymentRelationship, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.DeploymentRelationships = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ServiceAccountAndRoles) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (m *GetServiceAccountResponse) CloneVT() *GetServiceAccountResponse {
	if m == nil {
		return (*GetServiceAccountResponse)(nil)
	}
	r := &GetServiceAccountResponse{
		SaAndRole: m.SaAndRole.CloneVT(),
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetServiceAccountResponse) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (m *SADeploymentRelationship) CloneVT() *SADeploymentRelationship {
	if m == nil {
		return (*SADeploymentRelationship)(nil)
	}
	r := &SADeploymentRelationship{
		Id:   m.Id,
		Name: m.Name,
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SADeploymentRelationship) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (this *ListServiceAccountResponse) EqualVT(that *ListServiceAccountResponse) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if len(this.SaAndRoles) != len(that.SaAndRoles) {
		return false
	}
	for i, vx := range this.SaAndRoles {
		vy := that.SaAndRoles[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ServiceAccountAndRoles{}
			}
			if q == nil {
				q = &ServiceAccountAndRoles{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ServiceAccountAndRoles) EqualVT(that *ServiceAccountAndRoles) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if equal, ok := interface{}(this.ServiceAccount).(interface {
		EqualVT(*storage.ServiceAccount) bool
	}); ok {
		if !equal.EqualVT(that.ServiceAccount) {
			return false
		}
	} else if !proto.Equal(this.ServiceAccount, that.ServiceAccount) {
		return false
	}
	if len(this.ClusterRoles) != len(that.ClusterRoles) {
		return false
	}
	for i, vx := range this.ClusterRoles {
		vy := that.ClusterRoles[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &storage.K8SRole{}
			}
			if q == nil {
				q = &storage.K8SRole{}
			}
			if equal, ok := interface{}(p).(interface{ EqualVT(*storage.K8SRole) bool }); ok {
				if !equal.EqualVT(q) {
					return false
				}
			} else if !proto.Equal(p, q) {
				return false
			}
		}
	}
	if len(this.ScopedRoles) != len(that.ScopedRoles) {
		return false
	}
	for i, vx := range this.ScopedRoles {
		vy := that.ScopedRoles[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ScopedRoles{}
			}
			if q == nil {
				q = &ScopedRoles{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.DeploymentRelationships) != len(that.DeploymentRelationships) {
		return false
	}
	for i, vx := range this.DeploymentRelationships {
		vy := that.DeploymentRelationships[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &SADeploymentRelationship{}
			}
			if q == nil {
				q = &SADeploymentRelationship{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetServiceAccountResponse) EqualVT(that *GetServiceAccountResponse) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if !this.SaAndRole.EqualVT(that.SaAndRole) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SADeploymentRelationship) EqualVT(that *SADeploymentRelationship) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceAccountServiceClient is the client API for ServiceAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceAccountServiceClient interface {
	GetServiceAccount(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetServiceAccountResponse, error)
	ListServiceAccounts(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListServiceAccountResponse, error)
}

type serviceAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceAccountServiceClient(cc grpc.ClientConnInterface) ServiceAccountServiceClient {
	return &serviceAccountServiceClient{cc}
}

func (c *serviceAccountServiceClient) GetServiceAccount(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetServiceAccountResponse, error) {
	out := new(GetServiceAccountResponse)
	err := c.cc.Invoke(ctx, "/v1.ServiceAccountService/GetServiceAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAccountServiceClient) ListServiceAccounts(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListServiceAccountResponse, error) {
	out := new(ListServiceAccountResponse)
	err := c.cc.Invoke(ctx, "/v1.ServiceAccountService/ListServiceAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAccountServiceServer is the server API for ServiceAccountService service.
// All implementations must embed UnimplementedServiceAccountServiceServer
// for forward compatibility
type ServiceAccountServiceServer interface {
	GetServiceAccount(context.Context, *ResourceByID) (*GetServiceAccountResponse, error)
	ListServiceAccounts(context.Context, *RawQuery) (*ListServiceAccountResponse, error)
	mustEmbedUnimplementedServiceAccountServiceServer()
}

// UnimplementedServiceAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceAccountServiceServer struct {
}

func (UnimplementedServiceAccountServiceServer) GetServiceAccount(context.Context, *ResourceByID) (*GetServiceAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceAccount not implemented")
}
func (UnimplementedServiceAccountServiceServer) ListServiceAccounts(context.Context, *RawQuery) (*ListServiceAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServiceAccounts not implemented")
}
func (UnimplementedServiceAccountServiceServer) mustEmbedUnimplementedServiceAccountServiceServer() {}

// UnsafeServiceAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceAccountServiceServer will
// result in compilation errors.
type UnsafeServiceAccountServiceServer interface {
	mustEmbedUnimplementedServiceAccountServiceServer()
}

func RegisterServiceAccountServiceServer(s grpc.ServiceRegistrar, srv ServiceAccountServiceServer) {
	s.RegisterService(&ServiceAccountService_ServiceDesc, srv)
}

func _ServiceAccountService_GetServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountServiceServer).GetServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ServiceAccountService/GetServiceAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountServiceServer).GetServiceAccount(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAccountService_ListServiceAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountServiceServer).ListServiceAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ServiceAccountService/ListServiceAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountServiceServer).ListServiceAccounts(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceAccountService_ServiceDesc is the grpc.ServiceDesc for ServiceAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ServiceAccountService",
	HandlerType: (*ServiceAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceAccount",
			Handler:    _ServiceAccountService_GetServiceAccount_Handler,
		},
		{
			MethodName: "ListServiceAccounts",
			Handler:    _ServiceAccountService_ListServiceAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/service_account_service.proto",
}

func (m *ListServiceAccountResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListServiceAccountResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListServiceAccountResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.SaAndRoles) > 0 {
		for iNdEx := len(m.SaAndRoles) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.SaAndRoles[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ServiceAccountAndRoles) MarshalVT() (dAtA []byte, err error) {
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

func (m *ServiceAccountAndRoles) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ServiceAccountAndRoles) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.DeploymentRelationships) > 0 {
		for iNdEx := len(m.DeploymentRelationships) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DeploymentRelationships[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ScopedRoles) > 0 {
		for iNdEx := len(m.ScopedRoles) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ScopedRoles[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClusterRoles) > 0 {
		for iNdEx := len(m.ClusterRoles) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.ClusterRoles[iNdEx]).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.ClusterRoles[iNdEx])
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
	}
	if m.ServiceAccount != nil {
		if vtmsg, ok := interface{}(m.ServiceAccount).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ServiceAccount)
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

func (m *GetServiceAccountResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *GetServiceAccountResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *GetServiceAccountResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.SaAndRole != nil {
		size, err := m.SaAndRole.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SADeploymentRelationship) MarshalVT() (dAtA []byte, err error) {
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

func (m *SADeploymentRelationship) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *SADeploymentRelationship) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarint(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListServiceAccountResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SaAndRoles) > 0 {
		for _, e := range m.SaAndRoles {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ServiceAccountAndRoles) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServiceAccount != nil {
		if size, ok := interface{}(m.ServiceAccount).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ServiceAccount)
		}
		n += 1 + l + sov(uint64(l))
	}
	if len(m.ClusterRoles) > 0 {
		for _, e := range m.ClusterRoles {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.ScopedRoles) > 0 {
		for _, e := range m.ScopedRoles {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DeploymentRelationships) > 0 {
		for _, e := range m.DeploymentRelationships {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *GetServiceAccountResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SaAndRole != nil {
		l = m.SaAndRole.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *SADeploymentRelationship) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ListServiceAccountResponse) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListServiceAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListServiceAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SaAndRoles", wireType)
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
			m.SaAndRoles = append(m.SaAndRoles, &ServiceAccountAndRoles{})
			if err := m.SaAndRoles[len(m.SaAndRoles)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ServiceAccountAndRoles) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: ServiceAccountAndRoles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceAccountAndRoles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAccount", wireType)
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
			if m.ServiceAccount == nil {
				m.ServiceAccount = &storage.ServiceAccount{}
			}
			if unmarshal, ok := interface{}(m.ServiceAccount).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.ServiceAccount); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterRoles", wireType)
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
			m.ClusterRoles = append(m.ClusterRoles, &storage.K8SRole{})
			if unmarshal, ok := interface{}(m.ClusterRoles[len(m.ClusterRoles)-1]).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.ClusterRoles[len(m.ClusterRoles)-1]); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScopedRoles", wireType)
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
			m.ScopedRoles = append(m.ScopedRoles, &ScopedRoles{})
			if err := m.ScopedRoles[len(m.ScopedRoles)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeploymentRelationships", wireType)
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
			m.DeploymentRelationships = append(m.DeploymentRelationships, &SADeploymentRelationship{})
			if err := m.DeploymentRelationships[len(m.DeploymentRelationships)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GetServiceAccountResponse) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetServiceAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetServiceAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SaAndRole", wireType)
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
			if m.SaAndRole == nil {
				m.SaAndRole = &ServiceAccountAndRoles{}
			}
			if err := m.SaAndRole.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *SADeploymentRelationship) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: SADeploymentRelationship: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SADeploymentRelationship: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
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
