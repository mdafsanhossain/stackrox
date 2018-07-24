package service

import (
	"context"

	alertDataStore "bitbucket.org/stack-rox/apollo/central/alert/datastore"
	clusterDataStore "bitbucket.org/stack-rox/apollo/central/cluster/datastore"
	deploymentDataStore "bitbucket.org/stack-rox/apollo/central/deployment/datastore"
	imageDataStore "bitbucket.org/stack-rox/apollo/central/image/datastore"
	"bitbucket.org/stack-rox/apollo/central/role/resources"
	"bitbucket.org/stack-rox/apollo/central/service"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/auth/permissions"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/perrpc"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/authz/user"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// To access summaries, we require users to have view access to every summarized resource.
	// We could consider allowing people to get summaries for just the things they have access to,
	// but that requires non-trivial refactoring, so we'll do it if we feel the need later.
	// This variable is package-level to facilitate the unit test that asserts
	// that it covers all the summarized categories.
	// The keys are matched to fields in the v1.SummaryCountsResponse struct.
	summaryTypeToResource = map[string]permissions.Resource{
		"NumAlerts":      resources.Alert,
		"NumClusters":    resources.Cluster,
		"NumDeployments": resources.Deployment,
		"NumImages":      resources.Image,
	}
)

// SearchService provides APIs for search.
type serviceImpl struct {
	alerts      alertDataStore.DataStore
	clusters    clusterDataStore.DataStore
	deployments deploymentDataStore.DataStore
	images      imageDataStore.DataStore

	authorizer authz.Authorizer
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterSummaryServiceServer(grpcServer, s)
}

// RegisterServiceHandlerFromEndpoint registers this service with the given gRPC Gateway endpoint.
func (s *serviceImpl) RegisterServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterSummaryServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *serviceImpl) initializeAuthorizer() {
	requiredPermissions := make([]permissions.Permission, 0, len(summaryTypeToResource))
	for _, resource := range summaryTypeToResource {
		requiredPermissions = append(requiredPermissions, permissions.View(resource))
	}
	s.authorizer = perrpc.FromMap(
		map[authz.Authorizer][]string{
			user.With(requiredPermissions...): {
				"/v1.SummaryService/GetSummaryCounts",
			},
		},
	)
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, service.ReturnErrorCode(s.authorizer.Authorized(ctx, fullMethodName))
}

// GetSummaryCounts returns the global counts of alerts, clusters, deployments, and images.
func (s *serviceImpl) GetSummaryCounts(context.Context, *empty.Empty) (*v1.SummaryCountsResponse, error) {
	alerts, err := s.alerts.CountAlerts()
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	clusters, err := s.clusters.CountClusters()
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	deployments, err := s.deployments.CountDeployments()
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	images, err := s.images.CountImages()
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.SummaryCountsResponse{
		NumAlerts:      int64(alerts),
		NumClusters:    int64(clusters),
		NumDeployments: int64(deployments),
		NumImages:      int64(images),
	}, nil
}
