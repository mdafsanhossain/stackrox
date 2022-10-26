package resources

import (
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	networkPolicyConversion "github.com/stackrox/rox/pkg/protoconv/networkpolicy"
	"github.com/stackrox/rox/sensor/common/store"
	"github.com/stackrox/rox/sensor/kubernetes/eventpipeline/message"
	networkingV1 "k8s.io/api/networking/v1"
)

// networkPolicyDispatcher handles network policy resource events.
type networkPolicyDispatcher struct {
	netpolStore     store.NetworkPolicyStore
	deploymentStore *DeploymentStore
}

func newNetworkPolicyDispatcher(networkPolicyStore store.NetworkPolicyStore, deploymentStore *DeploymentStore) *networkPolicyDispatcher {
	return &networkPolicyDispatcher{
		netpolStore:     networkPolicyStore,
		deploymentStore: deploymentStore,
	}
}

// ProcessEvent processes a network policy resource event and returns the sensor events to generate.
func (h *networkPolicyDispatcher) ProcessEvent(obj, old interface{}, action central.ResourceAction) *message.ResourceEvent {
	np := obj.(*networkingV1.NetworkPolicy)

	roxNetpol := networkPolicyConversion.KubernetesNetworkPolicyWrap{NetworkPolicy: np}.ToRoxNetworkPolicy()

	var reprocessingIds []string
	if features.NetworkPolicySystemPolicy.Enabled() {
		var roxOldNetpol *storage.NetworkPolicy
		if oldNp, ok := old.(*networkingV1.NetworkPolicy); ok && oldNp != nil {
			roxOldNetpol = networkPolicyConversion.KubernetesNetworkPolicyWrap{NetworkPolicy: oldNp}.ToRoxNetworkPolicy()
		}
		sel := h.getSelector(roxNetpol, roxOldNetpol)
		if action == central.ResourceAction_REMOVE_RESOURCE {
			h.netpolStore.Delete(roxNetpol.GetId(), roxNetpol.GetNamespace())
		} else {
			h.netpolStore.Upsert(roxNetpol)
		}

		reprocessingIds = h.updateDeploymentsFromStore(roxNetpol, sel)
	}

	return wrapOutputMessage(
		[]*central.SensorEvent{
			{
				Id:     string(np.UID),
				Action: action,
				Resource: &central.SensorEvent_NetworkPolicy{
					NetworkPolicy: roxNetpol,
				},
			},
		}, nil, reprocessingIds)
}

func (h *networkPolicyDispatcher) getSelector(np, oldNp *storage.NetworkPolicy) selector {
	newsel := createSelector(np.GetSpec().GetPodSelector().GetMatchLabels(), emptyMatchesEverything())
	if oldNp != nil {
		oldsel := createSelector(oldNp.GetSpec().GetPodSelector().GetMatchLabels(), emptyMatchesEverything())
		return or(oldsel, newsel)
	}
	return newsel
}

func (h *networkPolicyDispatcher) updateDeploymentsFromStore(np *storage.NetworkPolicy, sel selector) []string {
	deployments := h.deploymentStore.getMatchingDeployments(np.GetNamespace(), sel)
	idsRequireReprocessing := make([]string, 0, len(deployments))
	for _, deploymentWrap := range deployments {
		idsRequireReprocessing = append(idsRequireReprocessing, deploymentWrap.GetId())
	}
	return idsRequireReprocessing
}
