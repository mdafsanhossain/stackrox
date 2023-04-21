package manager

import (
	"fmt"
	"strings"
	"time"

	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/pkg/alert"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/reflectutils"
	"github.com/stackrox/rox/pkg/sensor/hash"
	"github.com/stackrox/rox/pkg/stringutils"
	"github.com/stackrox/rox/pkg/sync"
)

// Deduper is an interface to deduping logic used to determine whether an event should be processed
type Deduper interface {
	// GetSuccessfulHashes returns a map of key to hashes that were successfully processed by Central
	// and thus can be persisted in the database as being processed
	GetSuccessfulHashes() map[string]uint64
	// ShouldProcess takes a message and determines if it should be processed
	ShouldProcess(msg *central.MsgFromSensor) bool
	// MarkSuccessful marks the message if necessary as being successfully processed, so it can be committed to the database
	// It will promote the message from the received map to the successfully processed map
	MarkSuccessful(msg *central.MsgFromSensor)
	// ProcessSync processes the Sensor sync message and reconciles the successfully processed and received maps
	ProcessSync()
}

// NewDeduper creates a new deduper from the passed existing hashes
func NewDeduper(existingHashes map[string]uint64) Deduper {
	process := make(map[string]*entry)
	for k, v := range existingHashes {
		process[k] = &entry{
			val:         v,
			preexisting: true,
		}
	}
	return &deduperImpl{
		received:              make(map[string]*entry),
		successfullyProcessed: process,
		hasher:                hash.NewHasher(),
	}
}

type entry struct {
	val         uint64
	create      time.Time
	preexisting bool
	processed   bool
}

type deduperImpl struct {
	hashLock sync.RWMutex
	// received map contains messages that have been received but not successfully processed
	received map[string]*entry
	// successfully processed map contains hashes of objects that have been successfully processed
	successfullyProcessed map[string]*entry

	hasher *hash.Hasher
}

// skipDedupe signifies that a message from Sensor can be deduped and thus can be stored in a hash
func skipDedupe(msg *central.MsgFromSensor) bool {
	eventMsg, ok := msg.Msg.(*central.MsgFromSensor_Event)
	if !ok {
		return true
	}
	if eventMsg.Event.GetProcessIndicator() != nil {
		return true
	}
	if alert.IsRuntimeAlertResult(msg.GetEvent().GetAlertResults()) {
		return true
	}
	// This can occur for a very short-lived deployment where alerts are not generated
	// but the deployment is being removed
	if alert.IsAlertResultResolved(msg.GetEvent().GetAlertResults()) {
		return true
	}
	if eventMsg.Event.GetReprocessDeployment() != nil {
		return true
	}
	return false
}

func (d *deduperImpl) shouldReprocess(hashKey string, hash uint64) bool {
	prevValue, ok := d.getValueNoLock(hashKey)
	if !ok {
		// This implies that a REMOVE event has been processed before this event
		// Note: we may want to handle alerts specifically because we should insert them as already resolved for completeness
		return false
	}
	// This implies that no new event was processed after the initial processing of the current message
	return prevValue == hash
}

func getIDFromKey(key string) string {
	return stringutils.GetAfter(key, ":")
}

func buildKey(typ, id string) string {
	return fmt.Sprintf("%s:%s", typ, id)
}

func getKey(msg *central.MsgFromSensor) string {
	event := msg.GetEvent()
	return buildKey(reflectutils.Type(event.GetResource()), event.GetId())
}

// MarkSuccessful marks a message as successfully processed
func (d *deduperImpl) MarkSuccessful(msg *central.MsgFromSensor) {
	// If the object isn't eligible for deduping then do not mark it as being successfully processed
	// because it does not exist in the received map
	if skipDedupe(msg) {
		return
	}
	key := getKey(msg)
	// If we are removing, then we do not need to mark it as successful as there is nothing more
	// to potentially dedupe. We do need to remove it from successfully processed as ShouldProcess is
	// evaluated as objects come into the queue and an object may have been successfully processed after
	if msg.GetEvent().GetAction() == central.ResourceAction_REMOVE_RESOURCE {
		concurrency.WithLock(&d.hashLock, func() {
			delete(d.successfullyProcessed, key)
		})
		return
	}

	d.hashLock.Lock()
	defer d.hashLock.Unlock()

	val, ok := d.received[key]
	// Only remove from this map if the hash matches as received could contain a more recent event
	if ok && val.val == msg.GetEvent().GetSensorHash() {
		delete(d.received, key)
	}
	d.successfullyProcessed[key] = &entry{
		val:       msg.GetEvent().GetSensorHash(),
		processed: true,
	}
}

func (d *deduperImpl) getValueNoLock(key string) (uint64, bool) {
	if prevValue, ok := d.received[key]; ok {
		return prevValue.val, ok
	}
	prevValue, ok := d.successfullyProcessed[key]
	if !ok {
		return 0, false
	}
	return prevValue.val, true
}

func (d *deduperImpl) ProcessSync() {
	// Reconcile successfully processed map with received map. Any keys that exist in successfully processed
	// but do not exist in received, can be dropped from successfully processed
	d.hashLock.Lock()
	defer d.hashLock.Unlock()

	alertResource := reflectutils.Type((*central.SensorEvent_AlertResults)(nil))
	deploymentResource := reflectutils.Type((*central.SensorEvent_Deployment)(nil))

	for k, v := range d.successfullyProcessed {
		// Ignore alerts in the first pass because they are not reconciled at the same time
		if strings.HasPrefix(k, alertResource) {
			continue
		}
		if !v.processed {
			if val, ok := d.received[k]; ok && val.processed {
				continue
			}
			log.Infof("Reconciling key %v : %v", k, len(d.successfullyProcessed))
			delete(d.successfullyProcessed, k)
			log.Infof("Reconciled key %v : %v", k, len(d.successfullyProcessed))
			// If a deployment is being removed due to reconciliation, then we will need to remove the alerts too
			if strings.HasPrefix(k, deploymentResource) {
				alertKey := buildKey(alertResource, getIDFromKey(k))
				log.Infof("Reconciling key %v : %v", alertKey, len(d.successfullyProcessed))
				delete(d.successfullyProcessed, k)
				log.Infof("Reconciled key %v : %v", alertKey, len(d.successfullyProcessed))
			}
		}
		// Mark it now as being not processed because reconciliation is completed
		v.processed = false
	}
}

// ShouldProcess determines if a message should be processed or if it should be deduped and dropped
func (d *deduperImpl) ShouldProcess(msg *central.MsgFromSensor) bool {
	if skipDedupe(msg) {
		return true
	}
	log.Infof("ShouldProcess: %T %v %v", msg.GetEvent().GetResource(), msg.GetEvent().GetId(), msg.GetEvent().GetAction())
	event := msg.GetEvent()
	key := getKey(msg)
	switch event.GetAction() {
	case central.ResourceAction_REMOVE_RESOURCE:
		d.hashLock.Lock()
		defer d.hashLock.Unlock()

		delete(d.received, key)
		delete(d.successfullyProcessed, key)
		return true
	case central.ResourceAction_SYNC_RESOURCE:
		// check if element is in successfully processed and mark as processed for syncs so that these are not reconciled away
		if val, ok := d.successfullyProcessed[key]; ok {
			val.processed = true
		}
	}
	// Backwards compatibility with a previous Sensor
	if event.GetSensorHashOneof() == nil {
		// Compute the sensor hash
		hashValue, ok := d.hasher.HashEvent(msg.GetEvent())
		if !ok {
			return true
		}
		event.SensorHashOneof = &central.SensorEvent_SensorHash{
			SensorHash: hashValue,
		}
	}
	// In the reprocessing case, the above will never evaluate to not nil, but it makes testing easier
	if msg.GetProcessingAttempt() > 0 {
		return d.shouldReprocess(key, event.GetSensorHash())
	}

	d.hashLock.Lock()
	defer d.hashLock.Unlock()

	prevValue, ok := d.getValueNoLock(key)
	if ok && prevValue == event.GetSensorHash() {
		return false
	}
	log.Infof("Setting %s for received: %d", key, len(d.received))
	d.received[key] = &entry{
		val:       event.GetSensorHash(),
		processed: true,
	}
	log.Infof("After setting %s for received: %d", key, len(d.received))
	return true
}

// GetSuccessfulHashes returns a copy of the successfullyProcessed map
func (d *deduperImpl) GetSuccessfulHashes() map[string]uint64 {
	d.hashLock.RLock()
	defer d.hashLock.RUnlock()

	copied := make(map[string]uint64, len(d.successfullyProcessed))
	for k, v := range d.successfullyProcessed {
		copied[k] = v.val
	}
	return copied
}
