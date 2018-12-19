package secrets

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	clusterMocks "github.com/stackrox/rox/central/cluster/datastore/mocks"
	secretMocks "github.com/stackrox/rox/central/secret/datastore/mocks"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/fixtures"
	"github.com/stretchr/testify/suite"
)

func TestPipeline(t *testing.T) {
	suite.Run(t, new(PipelineTestSuite))
}

type PipelineTestSuite struct {
	suite.Suite

	ctx      context.Context
	clusters *clusterMocks.MockDataStore
	secrets  *secretMocks.MockDataStore

	mockCtrl *gomock.Controller
}

func (suite *PipelineTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.clusters = clusterMocks.NewMockDataStore(suite.mockCtrl)
	suite.secrets = secretMocks.NewMockDataStore(suite.mockCtrl)
}

func (suite *PipelineTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *PipelineTestSuite) TestRun() {
	secret := fixtures.GetSecret()

	suite.clusters.EXPECT().GetCluster("clusterid").Return(&storage.Cluster{Id: "clusterid", Name: "clustername"}, true, nil)
	suite.secrets.EXPECT().UpsertSecret(secret).Return(nil)

	pipeline := NewPipeline(suite.clusters, suite.secrets)
	sensorEvent := &central.SensorEvent{
		Id:        "secretid",
		ClusterId: "clusterid",
		Action:    central.ResourceAction_CREATE_RESOURCE,
		Resource: &central.SensorEvent_Secret{
			Secret: secret,
		},
	}
	err := pipeline.Run(sensorEvent, nil)
	suite.NoError(err)
}
