// Code generated by pg-bindings generator. DO NOT EDIT.
package postgres

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	metrics "github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/auxpb"
	storage "github.com/stackrox/rox/generated/storage"
	ops "github.com/stackrox/rox/pkg/metrics"
	search "github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/blevesearch"
	"github.com/stackrox/rox/pkg/search/postgres"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

func init() {
	mapping.RegisterCategoryToTable(v1.SearchCategory_COMPLIANCE_RESULTS, schema)
}

// NewIndexer returns new indexer for `storage.ComplianceRunResults`.
func NewIndexer(db *pgxpool.Pool) *indexerImpl {
	return &indexerImpl{
		db: db,
	}
}

type indexerImpl struct {
	db *pgxpool.Pool
}

func (b *indexerImpl) Count(q *auxpb.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "ComplianceRunResults")

	return postgres.RunCountRequest(v1.SearchCategory_COMPLIANCE_RESULTS, q, b.db)
}

func (b *indexerImpl) Search(q *auxpb.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "ComplianceRunResults")

	return postgres.RunSearchRequest(v1.SearchCategory_COMPLIANCE_RESULTS, q, b.db)
}

//// Stubs for satisfying interfaces

func (b *indexerImpl) AddComplianceRunResults(deployment *storage.ComplianceRunResults) error {
	return nil
}

func (b *indexerImpl) AddComplianceRunResultss(_ []*storage.ComplianceRunResults) error {
	return nil
}

func (b *indexerImpl) DeleteComplianceRunResults(id string) error {
	return nil
}

func (b *indexerImpl) DeleteComplianceRunResultss(_ []string) error {
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return nil
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	return false, nil
}
