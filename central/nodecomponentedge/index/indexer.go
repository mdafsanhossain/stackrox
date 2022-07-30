// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	bleve "github.com/blevesearch/bleve"
	"github.com/stackrox/rox/generated/auxpb"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

type Indexer interface {
	AddNodeComponentEdge(nodecomponentedge *storage.NodeComponentEdge) error
	AddNodeComponentEdges(nodecomponentedges []*storage.NodeComponentEdge) error
	Count(q *auxpb.Query, opts ...blevesearch.SearchOption) (int, error)
	DeleteNodeComponentEdge(id string) error
	DeleteNodeComponentEdges(ids []string) error
	MarkInitialIndexingComplete() error
	NeedsInitialIndexing() (bool, error)
	Search(q *auxpb.Query, opts ...blevesearch.SearchOption) ([]search.Result, error)
}

func New(index bleve.Index) Indexer {
	return &indexerImpl{index: index}
}
