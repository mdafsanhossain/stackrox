package detector

import (
	"context"
	"time"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/expiringcache"
	"github.com/stackrox/rox/pkg/images/types"
)

const (
	scanTimeout = 3 * time.Minute
)

type scanResult struct {
	action     central.ResourceAction
	deployment *storage.Deployment
	images     []*storage.Image
}

type imageCacheKey struct {
	id, name string
}

type imageChanResult struct {
	image        *storage.Image
	containerIdx int
}

type enricher struct {
	imageSvc       v1.ImageServiceClient
	scanResultChan chan scanResult
	isSyncing      *concurrency.Flag

	imageCache expiringcache.Cache
	stopSig    concurrency.Signal
}

type cacheValue struct {
	signal concurrency.Signal
	image  *storage.Image
}

func (c *cacheValue) waitAndGet() *storage.Image {
	c.signal.Wait()
	return c.image
}

func (c *cacheValue) scanAndSet(svc v1.ImageServiceClient, ci *storage.ContainerImage, useSaved bool) {
	defer c.signal.Signal()

	ctx, cancel := context.WithTimeout(context.Background(), scanTimeout)
	defer cancel()
	scannedImage, err := svc.ScanImageInternal(ctx, &v1.ScanImageInternalRequest{
		Image:    ci,
		UseSaved: useSaved,
	})
	if err != nil {
		c.image = types.ToImage(ci)
		return
	}
	c.image = scannedImage.GetImage()
}

func newEnricher(isSyncing *concurrency.Flag) *enricher {
	return &enricher{
		scanResultChan: make(chan scanResult),
		isSyncing:      isSyncing,

		imageCache: expiringcache.NewExpiringCache(env.ReprocessInterval.DurationSetting()),
		stopSig:    concurrency.NewSignal(),
	}
}

type cacheKeyProvider interface {
	GetId() string
	GetName() *storage.ImageName
}

func getImageCacheKey(provider cacheKeyProvider) imageCacheKey {
	return imageCacheKey{
		id:   provider.GetId(),
		name: provider.GetName().GetFullName(),
	}
}

func (e *enricher) getImageFromCache(key imageCacheKey) (*storage.Image, bool) {
	value, _ := e.imageCache.Get(key).(*cacheValue)
	if value == nil {
		return nil, false
	}
	return value.waitAndGet(), true
}

func (e *enricher) runScan(containerIdx int, ci *storage.ContainerImage) imageChanResult {
	key := getImageCacheKey(ci)

	// Fast path
	img, ok := e.getImageFromCache(key)
	if ok {
		return imageChanResult{
			image:        img,
			containerIdx: containerIdx,
		}
	}

	newValue := &cacheValue{
		signal: concurrency.NewSignal(),
	}
	value := e.imageCache.GetOrSet(key, newValue).(*cacheValue)
	if newValue == value {
		value.scanAndSet(e.imageSvc, ci, e.isSyncing.Get())
	}
	return imageChanResult{
		image:        value.waitAndGet(),
		containerIdx: containerIdx,
	}
}

func (e *enricher) runImageScanAsync(imageChan chan<- imageChanResult, containerIdx int, ci *storage.ContainerImage) {
	go func() {
		// unguarded send (push to channel outside of a select) is allowed because the imageChan is a buffered channel of exact size
		imageChan <- e.runScan(containerIdx, ci)
	}()
}

func (e *enricher) getImages(deployment *storage.Deployment) []*storage.Image {
	imageChan := make(chan imageChanResult, len(deployment.GetContainers()))
	for idx, container := range deployment.GetContainers() {
		e.runImageScanAsync(imageChan, idx, container.GetImage())
	}
	images := make([]*storage.Image, len(deployment.GetContainers()))
	for i := 0; i < len(deployment.GetContainers()); i++ {
		imgResult := <-imageChan
		images[imgResult.containerIdx] = imgResult.image
	}

	return images
}

func (e *enricher) blockingScan(deployment *storage.Deployment, action central.ResourceAction) {
	select {
	case <-e.stopSig.Done():
		return
	case e.scanResultChan <- scanResult{
		action:     action,
		deployment: deployment,
		images:     e.getImages(deployment),
	}:
	}
}

func (e *enricher) outputChan() <-chan scanResult {
	return e.scanResultChan
}

func (e *enricher) stop() {
	e.stopSig.Signal()
}
