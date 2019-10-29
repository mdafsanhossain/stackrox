package datastore

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/ranking"
	"github.com/stackrox/rox/central/risk/datastore/internal/index"
	"github.com/stackrox/rox/central/risk/datastore/internal/search"
	"github.com/stackrox/rox/central/risk/datastore/internal/store"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	pkgSearch "github.com/stackrox/rox/pkg/search"
)

// DataStore is an intermediary to RiskStorage.
//go:generate mockgen-wrapper
type DataStore interface {
	Search(ctx context.Context, q *v1.Query) ([]pkgSearch.Result, error)
	SearchRawRisks(ctx context.Context, q *v1.Query) ([]*storage.Risk, error)

	GetRisk(ctx context.Context, subjectID string, subjectType storage.RiskSubjectType) (*storage.Risk, bool, error)
	GetRiskByIndicators(ctx context.Context, subjectID string, subjectType storage.RiskSubjectType, riskIndicatorNames []string) (risk *storage.Risk, err error)
	UpsertRisk(ctx context.Context, risk *storage.Risk) error
	RemoveRisk(ctx context.Context, subjectID string, subjectType storage.RiskSubjectType) error
}

// New returns a new instance of DataStore using the input store, indexer, and searcher.
func New(store store.Store, indexer index.Indexer, searcher search.Searcher) (DataStore, error) {

	d := &datastoreImpl{
		storage:  store,
		indexer:  indexer,
		searcher: searcher,
		subjectTypeToRanker: map[string]*ranking.Ranker{
			storage.RiskSubjectType_CLUSTER.String():         ranking.ClusterRanker(),
			storage.RiskSubjectType_NAMESPACE.String():       ranking.NamespaceRanker(),
			storage.RiskSubjectType_DEPLOYMENT.String():      ranking.DeploymentRanker(),
			storage.RiskSubjectType_IMAGE.String():           ranking.ImageRanker(),
			storage.RiskSubjectType_IMAGE_COMPONENT.String(): ranking.ImageComponentRanker(),
		},
	}
	if err := d.buildIndex(); err != nil {
		return nil, errors.Wrap(err, "failed to build index from existing store")
	}
	return d, nil

}
