package resources

import (
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/sensor/common/selector"
	"github.com/stackrox/rox/sensor/common/store"
	"github.com/stackrox/rox/sensor/common/store/service/servicewrapper"
)

// portExposureReconciler reconciles the port exposures in the deployment store on receiving
// service or route updates.
type portExposureReconciler interface {
	UpdateExposuresForMatchingDeployments(namespace string, sel selector.Selector) []*central.SensorEvent
	UpdateExposureOnServiceCreate(svc servicewrapper.SelectorRouteWrap) []*central.SensorEvent
}

type portExposureReconcilerImpl struct {
	deploymentStore *DeploymentStore
	serviceStore    store.ServiceStore
}

func newPortExposureReconciler(deploymentStore *DeploymentStore, serviceStore store.ServiceStore) portExposureReconciler {
	return &portExposureReconcilerImpl{
		deploymentStore: deploymentStore,
		serviceStore:    serviceStore,
	}
}

func (p *portExposureReconcilerImpl) UpdateExposuresForMatchingDeployments(namespace string, sel selector.Selector) []*central.SensorEvent {
	var events []*central.SensorEvent
	for _, deploymentWrap := range p.deploymentStore.getMatchingDeployments(namespace, sel) {
		if svcs := p.serviceStore.GetMatchingServicesWithRoutes(deploymentWrap.Namespace, deploymentWrap.PodLabels); len(svcs) > 0 || deploymentWrap.anyNonHostPort() {
			cloned := deploymentWrap.Clone()
			cloned.updatePortExposureFromServices(svcs...)
			p.deploymentStore.addOrUpdateDeployment(cloned)
		}

		events = append(events, deploymentWrap.toEvent(central.ResourceAction_UPDATE_RESOURCE))
	}
	return events
}

func (p *portExposureReconcilerImpl) UpdateExposureOnServiceCreate(svc servicewrapper.SelectorRouteWrap) []*central.SensorEvent {
	var events []*central.SensorEvent
	for _, deploymentWrap := range p.deploymentStore.getMatchingDeployments(svc.Namespace, svc.Selector) {
		cloned := deploymentWrap.Clone()
		cloned.updatePortExposure(svc)
		p.deploymentStore.addOrUpdateDeployment(cloned)
		events = append(events, cloned.toEvent(central.ResourceAction_UPDATE_RESOURCE))
	}
	return events
}
