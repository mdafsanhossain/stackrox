package service

import (
	"context"

	clusterDatastore "github.com/stackrox/rox/central/cluster/datastore"
	"github.com/stackrox/rox/central/enrichanddetect"
	"github.com/stackrox/rox/central/imageintegration/datastore"
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/grpc"
	"github.com/stackrox/rox/pkg/images/integration"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/registries"
	"github.com/stackrox/rox/pkg/scanners"
)

var (
	log = logging.LoggerForModule()
)

// Service provides the interface to the microservice that serves alert data.
type Service interface {
	grpc.APIService

	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)

	GetImageIntegration(ctx context.Context, request *v1.ResourceByID) (*v1.ImageIntegration, error)
	GetImageIntegrations(ctx context.Context, request *v1.GetImageIntegrationsRequest) (*v1.GetImageIntegrationsResponse, error)
	PutImageIntegration(ctx context.Context, request *v1.ImageIntegration) (*v1.Empty, error)
	PostImageIntegration(ctx context.Context, request *v1.ImageIntegration) (*v1.ImageIntegration, error)
	TestImageIntegration(ctx context.Context, request *v1.ImageIntegration) (*v1.Empty, error)
	DeleteImageIntegration(ctx context.Context, request *v1.ResourceByID) (*v1.Empty, error)
}

// New returns a new Service instance using the given DataStore.
func New(registryFactory registries.Factory,
	scannerFactory scanners.Factory,
	toNotify integration.ToNotify,
	datastore datastore.DataStore,
	clusterDatastore clusterDatastore.DataStore,
	enrichAndDetectLoop enrichanddetect.Loop) Service {
	return &serviceImpl{
		registryFactory:     registryFactory,
		scannerFactory:      scannerFactory,
		toNotify:            toNotify,
		datastore:           datastore,
		clusterDatastore:    clusterDatastore,
		enrichAndDetectLoop: enrichAndDetectLoop,
	}
}
