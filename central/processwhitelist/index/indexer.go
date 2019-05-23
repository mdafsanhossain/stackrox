// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	bleve "github.com/blevesearch/bleve"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
)

type Indexer interface {
	AddWhitelist(whitelist *storage.ProcessWhitelist) error
	AddWhitelists(whitelists []*storage.ProcessWhitelist) error
	DeleteWhitelist(id string) error
	Search(q *v1.Query) ([]search.Result, error)
}

func New(index bleve.Index) Indexer {
	return &indexerImpl{index: index}
}
