package service

import (
	"bitbucket.org/stack-rox/apollo/central/service"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/auth/tokenbased/user"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authn"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/allow"
	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ClusterService is the struct that manages the cluster API
type serviceImpl struct{}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterAuthServiceServer(grpcServer, s)
}

// RegisterServiceHandlerFromEndpoint registers this service with the given gRPC Gateway endpoint.
func (s *serviceImpl) RegisterServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, service.ReturnErrorCode(allow.Anonymous().Authorized(ctx, fullMethodName))
}

// GetAuthStatus retrieves the auth status based on the credentials given to the server.
func (s *serviceImpl) GetAuthStatus(ctx context.Context, request *empty.Empty) (*v1.AuthStatus, error) {
	authStatus, err := tokenAuthStatus(ctx)
	if err == nil {
		return authStatus, nil
	}

	authStatus, err = tlsAuthStatus(ctx)
	if err == nil {
		return authStatus, nil
	}

	return nil, status.Error(codes.Unauthenticated, "not authenticated")
}

func tokenAuthStatus(ctx context.Context) (*v1.AuthStatus, error) {
	identity, err := authn.FromTokenBasedIdentityContext(ctx)
	if err != nil {
		return nil, err
	}
	exp, err := types.TimestampProto(identity.Expiration())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "expiration time: %s", err)
	}
	var url string
	if asUserIdentity, ok := identity.(user.Identity); ok {
		url = asUserIdentity.AuthProvider().RefreshURL()
	}
	return &v1.AuthStatus{
		Id:         &v1.AuthStatus_UserId{UserId: identity.ID()},
		Expires:    exp,
		RefreshUrl: url,
	}, nil
}

func tlsAuthStatus(ctx context.Context) (*v1.AuthStatus, error) {
	id, err := authn.FromTLSContext(ctx)
	switch {
	case err == authn.ErrNoContext:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	case err != nil:
		return nil, status.Error(codes.Internal, err.Error())
	}
	exp, err := types.TimestampProto(id.Expiration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "expiration time: %s", err)
	}
	return &v1.AuthStatus{
		Id:      &v1.AuthStatus_ServiceId{ServiceId: id.Identity.V1()},
		Expires: exp,
	}, nil
}
