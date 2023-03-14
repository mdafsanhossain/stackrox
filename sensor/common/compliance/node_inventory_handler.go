package compliance

import (
	"github.com/stackrox/rox/generated/internalapi/sensor"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/sensor/common"
)

// MessageToComplianceWithAddress adds the hostname to sensor.MsgToCompliance so we know where to send it to.
type MessageToComplianceWithAddress struct {
	msg       *sensor.MsgToCompliance
	hostname  string
	broadcast bool
}

// nodeInventoryHandler is responsible for handling arriving NodeInventory messages, processing them, and sending them to central
type nodeInventoryHandler interface {
	common.SensorComponent
	Stopped() concurrency.ReadOnlyErrorSignal

	ComplianceC() <-chan *MessageToComplianceWithAddress
}

var _ nodeInventoryHandler = (*nodeInventoryHandlerImpl)(nil)

// NewNodeInventoryHandler returns a new instance of a NodeInventoryHandler
func NewNodeInventoryHandler(ch <-chan *storage.NodeInventory, matcher NodeIDMatcher) nodeInventoryHandler {
	return &nodeInventoryHandlerImpl{
		inventories:  ch,
		toCentral:    nil,
		centralReady: concurrency.NewSignal(),
		toCompliance: nil,
		lock:         &sync.Mutex{},
		stopper:      concurrency.NewStopper(),
		nodeMatcher:  matcher,
	}
}
