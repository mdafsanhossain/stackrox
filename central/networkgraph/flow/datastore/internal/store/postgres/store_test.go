//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/fixtures/fixtureconsts"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/timestamp"
	"github.com/stretchr/testify/suite"
)

const (
	clusterID = fixtureconsts.Cluster1

	flowsCountStmt = "select count(*) from network_flows"
)

type NetworkflowStoreSuite struct {
	suite.Suite
	store  FlowStore
	ctx    context.Context
	testDB *pgtest.TestPostgres
}

func TestNetworkflowStore(t *testing.T) {
	suite.Run(t, new(NetworkflowStoreSuite))
}

func (s *NetworkflowStoreSuite) SetupSuite() {
	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	} else {
		s.T().Setenv(env.PostgresDatastoreEnabled.EnvVar(), "true")
	}

	s.ctx = context.Background()

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB, clusterID)
	s.testDB.DB.Exec(s.ctx, "truncate table network_flows;")
}

func (s *NetworkflowStoreSuite) TearDownSuite() {
	//s.testDB.Teardown(s.T())
}

func getTimestamp(seconds int64) *types.Timestamp {
	return &types.Timestamp{
		Seconds: seconds,
	}
}

func (s *NetworkflowStoreSuite) TestStore() {
	secondCluster := fixtureconsts.Cluster2
	store2 := New(s.testDB.DB, secondCluster)

	networkFlow := &storage.NetworkFlow{
		Props: &storage.NetworkFlowProperties{
			SrcEntity:  &storage.NetworkEntityInfo{Type: storage.NetworkEntityInfo_DEPLOYMENT, Id: "a"},
			DstEntity:  &storage.NetworkEntityInfo{Type: storage.NetworkEntityInfo_DEPLOYMENT, Id: "b"},
			DstPort:    1,
			L4Protocol: storage.L4Protocol_L4_PROTOCOL_TCP,
		},
		LastSeenTimestamp: getTimestamp(1),
		ClusterId:         clusterID,
	}
	zeroTs := timestamp.MicroTS(0)

	foundNetworkFlows, _, err := s.store.GetAllFlows(s.ctx, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, 0)

	// Adding the same thing twice to ensure that we only retrieve 1 based on serial Flow_Id implementation
	s.NoError(s.store.UpsertFlows(s.ctx, []*storage.NetworkFlow{networkFlow}, zeroTs))
	networkFlow.LastSeenTimestamp = getTimestamp(2)
	s.NoError(s.store.UpsertFlows(s.ctx, []*storage.NetworkFlow{networkFlow}, zeroTs))
	foundNetworkFlows, _, err = s.store.GetAllFlows(s.ctx, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, 1)
	s.Equal(networkFlow, foundNetworkFlows[0])

	// Check the get all flows by since time
	foundNetworkFlows, _, err = s.store.GetAllFlows(s.ctx, getTimestamp(3))
	s.NoError(err)
	s.Len(foundNetworkFlows, 0)

	s.NoError(s.store.RemoveFlow(s.ctx, networkFlow.GetProps()))
	foundNetworkFlows, _, err = s.store.GetAllFlows(s.ctx, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, 0)

	s.NoError(s.store.UpsertFlows(s.ctx, []*storage.NetworkFlow{networkFlow}, zeroTs))

	err = s.store.RemoveFlowsForDeployment(s.ctx, networkFlow.GetProps().GetSrcEntity().GetId())
	s.NoError(err)

	foundNetworkFlows, _, err = s.store.GetAllFlows(s.ctx, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, 0)

	var networkFlows []*storage.NetworkFlow
	flowCount := 1000
	for i := 0; i < flowCount; i++ {
		networkFlow := &storage.NetworkFlow{}
		s.NoError(testutils.FullInit(networkFlow, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		networkFlow.ClusterId = clusterID
		networkFlows = append(networkFlows, networkFlow)

		for j := 0; j < 100; j++ {
			updatedFlow := &storage.NetworkFlow{
				Props: &storage.NetworkFlowProperties{
					DstPort:    uint32(j),
					L4Protocol: networkFlow.Props.L4Protocol,
					DstEntity: &storage.NetworkEntityInfo{
						Id:   networkFlow.Props.DstEntity.Id,
						Type: networkFlow.Props.DstEntity.Type,
					},
					SrcEntity: &storage.NetworkEntityInfo{
						Id:   networkFlow.Props.SrcEntity.Id,
						Type: networkFlow.Props.SrcEntity.Type,
					},
				},
				LastSeenTimestamp: networkFlow.LastSeenTimestamp,
				ClusterId:         networkFlow.ClusterId,
			}
			networkFlows = append(networkFlows, updatedFlow)
		}
	}

	log.Infof("SHREWS -- about to insert a crap ton of flows -- %d", len(networkFlows))
	s.NoError(s.store.UpsertFlows(s.ctx, networkFlows, zeroTs))
	log.Infof("SHREWS -- crap ton of flows inserted")

	foundNetworkFlows, _, err = s.store.GetAllFlows(s.ctx, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, flowCount)

	// Make sure store for second cluster does not find any flows
	foundNetworkFlows, _, err = store2.GetAllFlows(s.ctx, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, 0)

	// Add a flow to the second cluster
	networkFlow.ClusterId = secondCluster
	s.NoError(store2.UpsertFlows(s.ctx, []*storage.NetworkFlow{networkFlow}, zeroTs))

	foundNetworkFlows, _, err = store2.GetAllFlows(s.ctx, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, 1)

	pred := func(props *storage.NetworkFlowProperties) bool {
		return true
	}
	flowPredicate := func(flow *storage.NetworkFlow) bool {
		return true
	}
	foundNetworkFlows, _, err = store2.GetMatchingFlows(s.ctx, pred, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, 1)

	err = store2.RemoveMatchingFlows(s.ctx, pred, flowPredicate)
	s.NoError(err)

	// Store 2 flows should be removed.
	foundNetworkFlows, _, err = store2.GetAllFlows(s.ctx, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, 0)

	// Store 1 flows should remain
	foundNetworkFlows, _, err = s.store.GetAllFlows(s.ctx, nil)
	s.NoError(err)
	s.Len(foundNetworkFlows, flowCount)
}
