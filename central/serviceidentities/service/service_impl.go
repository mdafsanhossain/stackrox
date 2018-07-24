package service

import (
	"context"

	"bitbucket.org/stack-rox/apollo/central/role/resources"
	"bitbucket.org/stack-rox/apollo/central/service"
	"bitbucket.org/stack-rox/apollo/central/serviceidentities/store"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/auth/permissions"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/perrpc"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/user"
	"bitbucket.org/stack-rox/apollo/pkg/mtls"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authorizer = perrpc.FromMap(map[authz.Authorizer][]string{
		user.With(permissions.View(resources.ServiceIdentity)): {
			"/v1.ServiceIdentityService/GetServiceIdentities",
			"/v1.ServiceIdentityService/GetAuthorities",
		},
		user.With(permissions.Modify(resources.ServiceIdentity)): {
			"/v1.ServiceIdentityService/CreateServiceIdentity",
		},
	})
)

// IdentityService is the struct that manages the Service Identity API
type serviceImpl struct {
	storage store.Store
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterServiceIdentityServiceServer(grpcServer, s)
}

// RegisterServiceHandlerFromEndpoint registers this service with the given gRPC Gateway endpoint.
func (s *serviceImpl) RegisterServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterServiceIdentityServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, service.ReturnErrorCode(authorizer.Authorized(ctx, fullMethodName))
}

// GetServiceIdentities returns the currently defined service identities.
func (s *serviceImpl) GetServiceIdentities(ctx context.Context, _ *empty.Empty) (*v1.ServiceIdentityResponse, error) {
	serviceIdentities, err := s.storage.GetServiceIdentities()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.ServiceIdentityResponse{
		Identities: serviceIdentities,
	}, nil
}

// CreateServiceIdentity generates a new key and certificate for a service.
// The key and certificate are not retained and can not be retrieved except
// in the response to this API call.
func (s *serviceImpl) CreateServiceIdentity(ctx context.Context, request *v1.CreateServiceIdentityRequest) (*v1.CreateServiceIdentityResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "Request must be nonempty")
	}
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "ID must be nonempty")
	}
	if request.GetType() == v1.ServiceType_UNKNOWN_SERVICE {
		return nil, status.Error(codes.InvalidArgument, "Service type must be nonempty")
	}
	cert, key, id, err := mtls.IssueNewCert(mtls.CommonName{ServiceType: request.GetType(), Identifier: request.GetId()}, s.storage)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.CreateServiceIdentityResponse{
		Identity:    id,
		Certificate: string(cert),
		PrivateKey:  string(key),
	}, nil
}

// GetAuthorities returns the authorities currently in use.
func (s *serviceImpl) GetAuthorities(ctx context.Context, request *empty.Empty) (*v1.Authorities, error) {
	ca, err := mtls.CACertPEM()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.Authorities{
		Authorities: []*v1.Authority{
			{
				Certificate: string(ca),
			},
		},
	}, nil
}
