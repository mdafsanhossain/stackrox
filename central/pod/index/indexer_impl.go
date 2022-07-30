// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	"bytes"
	"time"

	bleve "github.com/blevesearch/bleve"
	metrics "github.com/stackrox/rox/central/metrics"
	mappings "github.com/stackrox/rox/central/pod/mappings"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/auxpb"
	storage "github.com/stackrox/rox/generated/storage"
	batcher "github.com/stackrox/rox/pkg/batcher"
	ops "github.com/stackrox/rox/pkg/metrics"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

const batchSize = 5000

const resourceName = "Pod"

type indexerImpl struct {
	index bleve.Index
}

type podWrapper struct {
	*storage.Pod `json:"pod"`
	Type         string `json:"type"`
}

func (b *indexerImpl) AddPod(pod *storage.Pod) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Add, "Pod")
	if err := b.index.Index(pod.GetId(), &podWrapper{
		Pod:  pod,
		Type: v1.SearchCategory_PODS.String(),
	}); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) AddPods(pods []*storage.Pod) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.AddMany, "Pod")
	batchManager := batcher.New(len(pods), batchSize)
	for {
		start, end, ok := batchManager.Next()
		if !ok {
			break
		}
		if err := b.processBatch(pods[start:end]); err != nil {
			return err
		}
	}
	return nil
}

func (b *indexerImpl) processBatch(pods []*storage.Pod) error {
	batch := b.index.NewBatch()
	for _, pod := range pods {
		if err := batch.Index(pod.GetId(), &podWrapper{
			Pod:  pod,
			Type: v1.SearchCategory_PODS.String(),
		}); err != nil {
			return err
		}
	}
	return b.index.Batch(batch)
}

func (b *indexerImpl) Count(q *auxpb.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "Pod")
	return blevesearch.RunCountRequest(v1.SearchCategory_PODS, q, b.index, mappings.OptionsMap, opts...)
}

func (b *indexerImpl) DeletePod(id string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Remove, "Pod")
	if err := b.index.Delete(id); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) DeletePods(ids []string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.RemoveMany, "Pod")
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
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "Pod")
	return blevesearch.RunSearchRequest(v1.SearchCategory_PODS, q, b.index, mappings.OptionsMap, opts...)
}
