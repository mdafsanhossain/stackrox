//go:build sql_integration

package datastore

import (
	"context"
	"testing"

	"github.com/blevesearch/bleve"
	"github.com/stackrox/rox/central/globalindex"
	"github.com/stackrox/rox/central/serviceaccount/internal/index"
	"github.com/stackrox/rox/central/serviceaccount/internal/store"
	pgStore "github.com/stackrox/rox/central/serviceaccount/internal/store/postgres"
	"github.com/stackrox/rox/central/serviceaccount/internal/store/rocksdb"
	serviceAccountSearch "github.com/stackrox/rox/central/serviceaccount/search"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/fixtures"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	rocksdbHelper "github.com/stackrox/rox/pkg/rocksdb"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/testutils/rocksdbtest"
	"github.com/stackrox/rox/pkg/uuid"
	"github.com/stretchr/testify/suite"
)

func TestServiceAccountDataStore(t *testing.T) {
	suite.Run(t, new(ServiceAccountDataStoreTestSuite))
}

type ServiceAccountDataStoreTestSuite struct {
	suite.Suite

	pool       *postgres.DB
	db         *rocksdbHelper.RocksDB
	bleveIndex bleve.Index

	indexer   index.Indexer
	searcher  serviceAccountSearch.Searcher
	storage   store.Store
	datastore DataStore

	ctx context.Context
}

func (suite *ServiceAccountDataStoreTestSuite) SetupSuite() {
	var err error
	if env.PostgresDatastoreEnabled.BooleanSetting() {
		pgtestbase := pgtest.ForT(suite.T())
		suite.Require().NotNil(pgtestbase)
		suite.pool = pgtestbase.DB
		suite.storage = pgStore.New(suite.pool)
		suite.indexer = pgStore.NewIndexer(suite.pool)
	} else {
		suite.bleveIndex, err = globalindex.TempInitializeIndices("")
		suite.Require().NoError(err)

		suite.db, err = rocksdbHelper.NewTemp(suite.T().Name())
		suite.Require().NoError(err)

		suite.storage = rocksdb.New(suite.db)
		suite.Require().NoError(err)
		suite.indexer = index.New(suite.bleveIndex)
	}
	suite.searcher = serviceAccountSearch.New(suite.storage, suite.indexer)
	suite.datastore, err = New(suite.storage, suite.indexer, suite.searcher)
	suite.Require().NoError(err)

	suite.ctx = sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
			sac.ResourceScopeKeys(resources.ServiceAccount)))
}

func (suite *ServiceAccountDataStoreTestSuite) TearDownSuite() {
	if env.PostgresDatastoreEnabled.BooleanSetting() {
		suite.pool.Close()
	} else {
		rocksdbtest.TearDownRocksDB(suite.db)
		suite.NoError(suite.bleveIndex.Close())
	}
}

func (suite *ServiceAccountDataStoreTestSuite) assertSearchResults(q *v1.Query, s *storage.ServiceAccount) {
	results, err := suite.datastore.SearchServiceAccounts(suite.ctx, q)
	suite.Require().NoError(err)
	if s != nil {
		suite.Require().Len(results, 1)
		suite.Equal(s.GetId(), results[0].GetId())
	} else {
		suite.Len(results, 0)
	}
}

func (suite *ServiceAccountDataStoreTestSuite) TestServiceAccountsDataStore() {
	sa := fixtures.GetServiceAccount()
	err := suite.datastore.UpsertServiceAccount(suite.ctx, sa)
	suite.Require().NoError(err)

	foundSA, found, err := suite.datastore.GetServiceAccount(suite.ctx, sa.GetId())
	suite.Require().NoError(err)
	suite.True(found)
	suite.Equal(sa, foundSA)

	nonexistentID := uuid.Nil.String()
	_, found, err = suite.datastore.GetServiceAccount(suite.ctx, nonexistentID)
	suite.Require().NoError(err)
	suite.False(found)

	validQ := search.NewQueryBuilder().AddStrings(search.Cluster, sa.GetClusterName()).ProtoQuery()
	suite.assertSearchResults(validQ, sa)

	invalidQ := search.NewQueryBuilder().AddStrings(search.Cluster, nonexistentID).ProtoQuery()
	suite.assertSearchResults(invalidQ, nil)

	err = suite.datastore.RemoveServiceAccount(suite.ctx, sa.GetId())
	suite.Require().NoError(err)

	_, found, err = suite.datastore.GetServiceAccount(suite.ctx, sa.GetId())
	suite.Require().NoError(err)
	suite.False(found)

	suite.assertSearchResults(validQ, nil)
}
