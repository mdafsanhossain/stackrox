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

	v1.ImageIntegrationServiceServer
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
