// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	"bytes"
	"time"

	bleve "github.com/blevesearch/bleve"
	metrics "github.com/stackrox/rox/central/metrics"
	mappings "github.com/stackrox/rox/central/namespace/index/mappings"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/auxpb"
	storage "github.com/stackrox/rox/generated/storage"
	batcher "github.com/stackrox/rox/pkg/batcher"
	ops "github.com/stackrox/rox/pkg/metrics"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

const batchSize = 5000

const resourceName = "NamespaceMetadata"

type indexerImpl struct {
	index bleve.Index
}

type namespaceMetadataWrapper struct {
	*storage.NamespaceMetadata `json:"namespace_metadata"`
	Type                       string `json:"type"`
}

func (b *indexerImpl) AddNamespaceMetadata(namespacemetadata *storage.NamespaceMetadata) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Add, "NamespaceMetadata")
	if err := b.index.Index(namespacemetadata.GetId(), &namespaceMetadataWrapper{
		NamespaceMetadata: namespacemetadata,
		Type:              v1.SearchCategory_NAMESPACES.String(),
	}); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) AddNamespaceMetadatas(namespacemetadatas []*storage.NamespaceMetadata) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.AddMany, "NamespaceMetadata")
	batchManager := batcher.New(len(namespacemetadatas), batchSize)
	for {
		start, end, ok := batchManager.Next()
		if !ok {
			break
		}
		if err := b.processBatch(namespacemetadatas[start:end]); err != nil {
			return err
		}
	}
	return nil
}

func (b *indexerImpl) processBatch(namespacemetadatas []*storage.NamespaceMetadata) error {
	batch := b.index.NewBatch()
	for _, namespacemetadata := range namespacemetadatas {
		if err := batch.Index(namespacemetadata.GetId(), &namespaceMetadataWrapper{
			NamespaceMetadata: namespacemetadata,
			Type:              v1.SearchCategory_NAMESPACES.String(),
		}); err != nil {
			return err
		}
	}
	return b.index.Batch(batch)
}

func (b *indexerImpl) Count(q *auxpb.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "NamespaceMetadata")
	return blevesearch.RunCountRequest(v1.SearchCategory_NAMESPACES, q, b.index, mappings.OptionsMap, opts...)
}

func (b *indexerImpl) DeleteNamespaceMetadata(id string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Remove, "NamespaceMetadata")
	if err := b.index.Delete(id); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) DeleteNamespaceMetadatas(ids []string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.RemoveMany, "NamespaceMetadata")
	batch := b.index.NewBatch()
	for _, id := range ids {
		batch.Delete(id)
	}
	if err := b.index.Batch(batch); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return b.index.SetInternal([]byte(resourceName), []byte("old"))
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	data, err := b.index.GetInternal([]byte(resourceName))
	if err != nil {
		return false, err
	}
	return !bytes.Equal([]byte("old"), data), nil
}

func (b *indexerImpl) Search(q *auxpb.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "NamespaceMetadata")
	return blevesearch.RunSearchRequest(v1.SearchCategory_NAMESPACES, q, b.index, mappings.OptionsMap, opts...)
}
