package search

import (
	"context"

	"github.com/stackrox/rox/central/policycategory/index"
	"github.com/stackrox/rox/central/policycategory/store"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/auxpb"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/search"
)

var (
	log = logging.LoggerForModule()
)

// Searcher provides search functionality on existing alerts
//go:generate mockgen-wrapper
type Searcher interface {
	Search(ctx context.Context, q *auxpb.Query) ([]search.Result, error)
	Count(ctx context.Context, q *auxpb.Query) (int, error)
	SearchCategories(ctx context.Context, q *auxpb.Query) ([]*v1.SearchResult, error)
	SearchRawCategories(ctx context.Context, q *auxpb.Query) ([]*storage.PolicyCategory, error)
}

// New returns a new instance of Searcher for the given storage and indexer.
func New(storage store.Store, indexer index.Indexer) Searcher {
	return &searcherImpl{
		storage:  storage,
		indexer:  indexer,
		searcher: formatSearcher(indexer),
	}
}
