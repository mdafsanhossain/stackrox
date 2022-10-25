// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.3.1-0.20220817155510-0ae748fd2007
// source: api/v1/pod_service.proto

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

func (m *PodsResponse) CloneVT() *PodsResponse {
	if m == nil {
		return (*PodsResponse)(nil)
	}
	r := &PodsResponse{}
	if rhs := m.Pods; rhs != nil {
		tmpContainer := make([]*storage.Pod, len(rhs))
		for k, v := range rhs {
			if vtpb, ok := interface{}(v).(interface{ CloneVT() *storage.Pod }); ok {
				tmpContainer[k] = vtpb.CloneVT()
			} else {
				tmpContainer[k] = proto.Clone(v).(*storage.Pod)
			}
		}
		r.Pods = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PodsResponse) CloneGenericVT() proto.Message {
	return m.CloneVT()
}

func (this *PodsResponse) EqualVT(that *PodsResponse) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if len(this.Pods) != len(that.Pods) {
		return false
	}
	for i, vx := range this.Pods {
		vy := that.Pods[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &storage.Pod{}
			}
			if q == nil {
				q = &storage.Pod{}
			}
			if equal, ok := interface{}(p).(interface{ EqualVT(*storage.Pod) bool }); ok {
				if !equal.EqualVT(q) {
					return false
				}
			} else if !proto.Equal(p, q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PodServiceClient is the client API for PodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PodServiceClient interface {
	// GetPods returns the pods.
	GetPods(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*PodsResponse, error)
}

type podServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPodServiceClient(cc grpc.ClientConnInterface) PodServiceClient {
	return &podServiceClient{cc}
}

func (c *podServiceClient) GetPods(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*PodsResponse, error) {
	out := new(PodsResponse)
	err := c.cc.Invoke(ctx, "/v1.PodService/GetPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodServiceServer is the server API for PodService service.
// All implementations must embed UnimplementedPodServiceServer
// for forward compatibility
type PodServiceServer interface {
	// GetPods returns the pods.
	GetPods(context.Context, *RawQuery) (*PodsResponse, error)
	mustEmbedUnimplementedPodServiceServer()
}

// UnimplementedPodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPodServiceServer struct {
}

func (UnimplementedPodServiceServer) GetPods(context.Context, *RawQuery) (*PodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPods not implemented")
}
func (UnimplementedPodServiceServer) mustEmbedUnimplementedPodServiceServer() {}

// UnsafePodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PodServiceServer will
// result in compilation errors.
type UnsafePodServiceServer interface {
	mustEmbedUnimplementedPodServiceServer()
}

func RegisterPodServiceServer(s grpc.ServiceRegistrar, srv PodServiceServer) {
	s.RegisterService(&PodService_ServiceDesc, srv)
}

func _PodService_GetPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodServiceServer).GetPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PodService/GetPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodServiceServer).GetPods(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// PodService_ServiceDesc is the grpc.ServiceDesc for PodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PodService",
	HandlerType: (*PodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPods",
			Handler:    _PodService_GetPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/pod_service.proto",
}

func (m *PodsResponse) MarshalVT() (dAtA []byte, err error) {
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

func (m *PodsResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *PodsResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Pods) > 0 {
		for iNdEx := len(m.Pods) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.Pods[iNdEx]).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.Pods[iNdEx])
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
	}
	return len(dAtA) - i, nil
}

func (m *PodsResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pods) > 0 {
		for _, e := range m.Pods {
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
	n += len(m.unknownFields)
	return n
}

func (m *PodsResponse) UnmarshalVT(dAtA []byte) error {
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
			return fmt.Errorf("proto: PodsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pods", wireType)
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
			m.Pods = append(m.Pods, &storage.Pod{})
			if unmarshal, ok := interface{}(m.Pods[len(m.Pods)-1]).(interface {
				UnmarshalVT([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.Pods[len(m.Pods)-1]); err != nil {
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
