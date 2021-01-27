package m53tom54

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations/rocksdbmigration"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/rocksdb"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stackrox/rox/pkg/testutils/rocksdbtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tecbot/gorocksdb"
)

func TestExecWebhookMigration(t *testing.T) {
	suite.Run(t, new(execWebhookTestSuite))
}

type execWebhookTestSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator

	db *rocksdb.RocksDB
}

func (suite *execWebhookTestSuite) SetupTest() {
	suite.envIsolator = envisolator.NewEnvIsolator(suite.T())
	suite.envIsolator.Setenv(features.K8sEventDetection.EnvVar(), "true")

	suite.db = rocksdbtest.RocksDBForT(suite.T())
}

func (suite *execWebhookTestSuite) TearDownTest() {
	rocksdbtest.TearDownRocksDB(suite.db)
	suite.envIsolator.RestoreAll()

}

func (suite *execWebhookTestSuite) TestMigrateClustersWithExecWebhooks() {
	if !features.K8sEventDetection.Enabled() {
		suite.T().Skipf("feature flag %s not enabled", features.K8sEventDetection.EnvVar())
	}

	clusters := []*storage.Cluster{
		{
			Id:   "1",
			Type: storage.ClusterType_OPENSHIFT_CLUSTER,
		},
		{
			Id:                        "2",
			Type:                      storage.ClusterType_OPENSHIFT_CLUSTER,
			AdmissionControllerEvents: true,
		},
		{
			Id:   "3",
			Type: storage.ClusterType_KUBERNETES_CLUSTER,
		},
		{
			Id:                        "4",
			Type:                      storage.ClusterType_KUBERNETES_CLUSTER,
			AdmissionControllerEvents: true,
		},
	}

	wb := gorocksdb.NewWriteBatch()
	for _, c := range clusters {
		bytes, err := proto.Marshal(c)
		suite.NoError(err)

		wb.Put(rocksdbmigration.GetPrefixedKey(clustersPrefix, []byte(c.Id)), bytes)
	}
	err := suite.db.Write(gorocksdb.NewDefaultWriteOptions(), wb)
	suite.NoError(err)

	// Migrate the data
	suite.NoError(migrateExecWebhook(suite.db.DB))

	expected := []*storage.Cluster{
		{
			Id:                        "1",
			Type:                      storage.ClusterType_OPENSHIFT_CLUSTER,
			AdmissionControllerEvents: false,
		},
		{
			Id:                        "2",
			Type:                      storage.ClusterType_OPENSHIFT_CLUSTER,
			AdmissionControllerEvents: true,
		},
		{
			Id:                        "3",
			Type:                      storage.ClusterType_KUBERNETES_CLUSTER,
			AdmissionControllerEvents: true,
		},
		{
			Id:                        "4",
			Type:                      storage.ClusterType_KUBERNETES_CLUSTER,
			AdmissionControllerEvents: true,
		},
	}
	readOpts := gorocksdb.NewDefaultReadOptions()
	it := suite.db.NewIterator(readOpts)
	defer it.Close()

	migratedClusters := make([]*storage.Cluster, 0, len(expected))
	for it.Seek(clustersPrefix); it.ValidForPrefix(clustersPrefix); it.Next() {
		cluster := &storage.Cluster{}
		if err := proto.Unmarshal(it.Value().Data(), cluster); err != nil {
			suite.NoError(err)
		}
		migratedClusters = append(migratedClusters, cluster)
	}

	assert.ElementsMatch(suite.T(), expected, migratedClusters)
}
