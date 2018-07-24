package service

import (
	"fmt"

	"bitbucket.org/stack-rox/apollo/central/detection"
	"bitbucket.org/stack-rox/apollo/central/imageintegration/datastore"
	"bitbucket.org/stack-rox/apollo/central/role/resources"
	"bitbucket.org/stack-rox/apollo/central/service"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/auth/permissions"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authn"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/or"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/perrpc"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/user"
	"bitbucket.org/stack-rox/apollo/pkg/secrets"
	"bitbucket.org/stack-rox/apollo/pkg/sources"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authorizer = or.SensorOrAuthorizer(perrpc.FromMap(map[authz.Authorizer][]string{
		user.With(permissions.View(resources.ImageIntegration)): {
			"/v1.ImageIntegrationService/GetImageIntegration",
			"/v1.ImageIntegrationService/GetImageIntegrations",
		},
		user.With(permissions.Modify(resources.ImageIntegration)): {
			"/v1.ImageIntegrationService/PostImageIntegration",
			"/v1.ImageIntegrationService/PutImageIntegration",
			"/v1.ImageIntegrationService/TestImageIntegration",
			"/v1.ImageIntegrationService/DeleteImageIntegration",
		},
	}))
)

// ImageIntegrationService is the struct that manages the ImageIntegration API
type serviceImpl struct {
	datastore datastore.DataStore
	detector  *detection.Detector
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterImageIntegrationServiceServer(grpcServer, s)
}

// RegisterServiceHandlerFromEndpoint registers this service with the given gRPC Gateway endpoint.
func (s *serviceImpl) RegisterServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterImageIntegrationServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, service.ReturnErrorCode(authorizer.Authorized(ctx, fullMethodName))
}

func scrubImageIntegration(i *v1.ImageIntegration) {
	i.Config = secrets.ScrubSecretsFromMap(i.Config)
	secrets.ScrubSecretsFromStruct(i)
}

// GetImageIntegration retrieves the integration based on the id passed
func (s *serviceImpl) GetImageIntegration(ctx context.Context, request *v1.ResourceByID) (*v1.ImageIntegration, error) {
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "Image integration id must be provided")
	}
	integration, exists, err := s.datastore.GetImageIntegration(request.GetId())
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Image integration %s not found", request.GetId()))
	}
	scrubImageIntegration(integration)
	return integration, nil
}

// GetImageIntegrations retrieves all image integrations that matches the request filters
func (s *serviceImpl) GetImageIntegrations(ctx context.Context, request *v1.GetImageIntegrationsRequest) (*v1.GetImageIntegrationsResponse, error) {
	integrations, err := s.datastore.GetImageIntegrations(request)
	if err != nil {
		return nil, err
	}
	identity, err := authn.FromTLSContext(ctx)
	switch {
	case err == authn.ErrNoContext:
		log.Debugf("No authentication context provided")
	case err != nil:
		log.Warnf("Error getting client identity: %s", err)
	case err == nil && identity.Name.ServiceType == v1.ServiceType_SENSOR_SERVICE:
		return &v1.GetImageIntegrationsResponse{Integrations: integrations}, nil
	}
	// Remove secrets for other API accessors.
	for _, i := range integrations {
		scrubImageIntegration(i)
	}
	return &v1.GetImageIntegrationsResponse{Integrations: integrations}, nil
}

// PutImageIntegration updates an image integration in the system
func (s *serviceImpl) PutImageIntegration(ctx context.Context, request *v1.ImageIntegration) (*empty.Empty, error) {
	// creates and validates the configuration
	source, err := sources.NewImageIntegration(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.datastore.UpdateImageIntegration(request); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	s.detector.UpdateImageIntegration(source)
	return &empty.Empty{}, nil
}

// PostImageIntegration inserts a new image integration into the system if it doesn't already exist
func (s *serviceImpl) PostImageIntegration(ctx context.Context, request *v1.ImageIntegration) (*v1.ImageIntegration, error) {
	if request.GetId() != "" {
		return nil, status.Error(codes.InvalidArgument, "Id field should be empty when posting a new image integration")
	}
	// creates and validates the configuration
	source, err := sources.NewImageIntegration(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	id, err := s.datastore.AddImageIntegration(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	request.Id = id
	s.detector.UpdateImageIntegration(source)
	return request, nil
}

// TestImageIntegration tests to see if the config is setup properly
func (s *serviceImpl) TestImageIntegration(ctx context.Context, request *v1.ImageIntegration) (*empty.Empty, error) {
	source, err := sources.NewImageIntegration(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := source.Test(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &empty.Empty{}, nil
}

// DeleteImageIntegration deletes an integration from the system
func (s *serviceImpl) DeleteImageIntegration(ctx context.Context, request *v1.ResourceByID) (*empty.Empty, error) {
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "Image integration id must be provided")
	}
	if err := s.datastore.RemoveImageIntegration(request.GetId()); err != nil {
		return nil, service.ReturnErrorCode(err)
	}
	s.detector.RemoveImageIntegration(request.GetId())
	return &empty.Empty{}, nil
}
