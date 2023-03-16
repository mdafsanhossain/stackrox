package nodeinventory

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	clusterDataStore "github.com/stackrox/rox/central/cluster/datastore"
	"github.com/stackrox/rox/central/enrichment"
	countMetrics "github.com/stackrox/rox/central/metrics"
	nodeDatastore "github.com/stackrox/rox/central/node/datastore"
	riskManager "github.com/stackrox/rox/central/risk/manager"
	"github.com/stackrox/rox/central/sensor/service/common"
	"github.com/stackrox/rox/central/sensor/service/pipeline"
	"github.com/stackrox/rox/central/sensor/service/pipeline/reconciliation"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/nodes/enricher"
)

var (
	log = logging.LoggerForModule()
)

// GetPipeline returns an instantiation of this particular pipeline
func GetPipeline() pipeline.Fragment {
	return newPipeline(clusterDataStore.Singleton(), nodeDatastore.Singleton(), enrichment.NodeEnricherSingleton(), riskManager.Singleton())
}

// newPipeline returns a new instance of Pipeline.
func newPipeline(clusters clusterDataStore.DataStore, nodes nodeDatastore.DataStore, enricher enricher.NodeEnricher, riskManager riskManager.Manager) pipeline.Fragment {
	return &pipelineImpl{
		clusterStore:  clusters,
		nodeDatastore: nodes,
		enricher:      enricher,
		riskManager:   riskManager,
	}
}

type pipelineImpl struct {
	clusterStore  clusterDataStore.DataStore
	nodeDatastore nodeDatastore.DataStore
	enricher      enricher.NodeEnricher
	riskManager   riskManager.Manager
}

func (p *pipelineImpl) Reconcile(ctx context.Context, clusterID string, storeMap *reconciliation.StoreMap) error {
	return nil
}

func (p *pipelineImpl) Match(msg *central.MsgFromSensor) bool {
	return msg.GetEvent().GetNodeInventory() != nil
}

// Run runs the pipeline template on the input and returns the output.
func (p *pipelineImpl) Run(ctx context.Context, clusterID string, msg *central.MsgFromSensor, _ common.MessageInjector) error {
	defer countMetrics.IncrementResourceProcessedCounter(pipeline.ActionToOperation(msg.GetEvent().GetAction()), metrics.NodeInventory)

	// Sanitize input.
	event := msg.GetEvent()
	ninv := event.GetNodeInventory()
	if ninv == nil {
		return errors.Errorf("unexpected resource type %T for node inventory", event.GetResource())
	}
	invStr := fmt.Sprintf("for node %s (id: %s)", ninv.GetNodeName(), ninv.GetNodeId())
	log.Infof("received node inventory %s", invStr)
	log.Debugf("node inventory %s contains %d packages to scan from %d content sets", invStr,
		len(ninv.GetComponents().GetRhelComponents()), len(ninv.GetComponents().GetRhelContentSets()))
	if event.GetAction() != central.ResourceAction_UNSET_ACTION_RESOURCE {
		log.Errorf("node inventory %s with unsupported action: %s", invStr, event.GetAction())
		return nil
	}
	ninv = ninv.Clone()

	// Read the node from the database, if not found we fail.
	node, found, err := p.nodeDatastore.GetNode(ctx, ninv.GetNodeId())
	if err != nil {
		log.Errorf("fetching node (id: %q) from the database: %v", ninv.GetNodeId(), err)
		return errors.WithMessagef(err, "fetching node: %s", ninv.GetNodeId())
	}
	if !found {
		log.Errorf("fetching node (id: %q) from the database: node does not exist", ninv.GetNodeId())
		return errors.WithMessagef(err, "node does not exist: %s", ninv.GetNodeId())
	}
	log.Debugf("node %s found, enriching with node inventory", nodeDatastore.NodeString(node))

	// Call Scanner to enrich the node inventory and attach the results to the node object.
	err = p.enricher.EnrichNodeWithInventory(node, ninv)
	if err != nil {
		log.Errorf("enriching node %s: %v", nodeDatastore.NodeString(node), err)
		return errors.WithMessagef(err, "enrinching node %s", nodeDatastore.NodeString(node))
	}
	log.Debugf("node inventory for node %s has been scanned and contains %d results",
		nodeDatastore.NodeString(node), len(node.GetScan().GetComponents()))

	// Update the whole node in the database with the new and previous information.
	err = p.riskManager.CalculateRiskAndUpsertNode(node)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (p *pipelineImpl) OnFinish(_ string) {}
