package store

import (
	"os"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	ptypes "github.com/gogo/protobuf/types"
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/bolthelper"
	"github.com/stretchr/testify/suite"
)

func TestClusterStore(t *testing.T) {
	suite.Run(t, new(ClusterStoreTestSuite))
}

type ClusterStoreTestSuite struct {
	suite.Suite

	db *bolt.DB

	store Store
}

func (suite *ClusterStoreTestSuite) SetupSuite() {
	db, err := bolthelper.NewTemp("cluster_test.db")
	if err != nil {
		suite.FailNow("Failed to make BoltDB", err.Error())
	}

	suite.db = db
	suite.store = New(db)
}

func (suite *ClusterStoreTestSuite) TeardownSuite() {
	suite.db.Close()
	os.Remove(suite.db.Path())
}

func (suite *ClusterStoreTestSuite) TestClusters() {
	checkin1 := time.Now()
	checkin2 := time.Now().Add(-1 * time.Hour)
	ts1, err := ptypes.TimestampProto(checkin1)
	suite.NoError(err)
	ts2, err := ptypes.TimestampProto(checkin2)
	suite.NoError(err)

	clusters := []*v1.Cluster{
		{
			Name:        "cluster1",
			MainImage:   "test-dtr.example.com/main",
			LastContact: ts1,
		},
		{
			Name:        "cluster2",
			MainImage:   "docker.io/stackrox/main",
			LastContact: ts2,
		},
	}

	// Test Add
	for _, b := range clusters {
		id, err := suite.store.AddCluster(b)
		suite.NoError(err)
		suite.NotEmpty(id)

		// Add the timestamp in the second list.
		t, err := ptypes.TimestampFromProto(b.GetLastContact())
		suite.NoError(err)
		err = suite.store.UpdateClusterContactTime(b.GetId(), t)
		suite.NoError(err)
	}

	for _, b := range clusters {
		got, exists, err := suite.store.GetCluster(b.GetId())
		suite.NoError(err)
		suite.True(exists)
		suite.Equal(got, b)
	}

	// Test Update
	for _, b := range clusters {
		b.MainImage = b.MainImage + "/main"
	}

	for _, b := range clusters {
		suite.NoError(suite.store.UpdateCluster(b))
	}

	for _, b := range clusters {
		got, exists, err := suite.store.GetCluster(b.GetId())
		suite.NoError(err)
		suite.True(exists)
		suite.Equal(got, b)
	}

	// Test Count
	count, err := suite.store.CountClusters()
	suite.NoError(err)
	suite.Equal(len(clusters), count)

	// Test Remove
	for _, b := range clusters {
		suite.NoError(suite.store.RemoveCluster(b.GetId()))
	}

	for _, b := range clusters {
		_, exists, err := suite.store.GetCluster(b.GetId())
		suite.NoError(err)
		suite.False(exists)
	}
}
