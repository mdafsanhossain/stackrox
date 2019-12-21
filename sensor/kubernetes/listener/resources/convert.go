package resources

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	openshift_appsv1 "github.com/openshift/api/apps/v1"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/containers"
	"github.com/stackrox/rox/pkg/env"
	imageUtils "github.com/stackrox/rox/pkg/images/utils"
	"github.com/stackrox/rox/pkg/kubernetes"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/protoconv/k8s"
	"github.com/stackrox/rox/pkg/protoconv/resources"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stackrox/rox/pkg/uuid"
	"k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/intstr"
	v1listers "k8s.io/client-go/listers/core/v1"
)

const (
	k8sStandalonePodType = "StaticPods"
	kubeSystemNamespace  = "kube-system"
	labelMaxLength       = 63
	trailingUIDLen       = 5
)

var (
	log = logging.LoggerForModule()

	k8sComponentLabelKeys = []string{
		"component",
		"k8s-app",
	}
)

func getK8sComponentID(component string) string {
	u, err := uuid.FromString(env.ClusterID.Setting())
	if err != nil {
		log.Error(err)
		return ""
	}
	return uuid.NewV5(u, component).String()
}

type deploymentWrap struct {
	*storage.Deployment
	registryOverride string
	original         interface{}
	portConfigs      map[portRef]*storage.PortConfig
	pods             []*v1.Pod
	podSelector      labels.Selector
}

// This checks if a reflect value is a Zero value, which means the field did not exist
func doesFieldExist(value reflect.Value) bool {
	return !reflect.DeepEqual(value, reflect.Value{})
}

func newDeploymentEventFromResource(obj interface{}, action *central.ResourceAction, deploymentType string,
	lister v1listers.PodLister, namespaceStore *namespaceStore, registryOverride string) *deploymentWrap {
	wrap := newWrap(obj, deploymentType, registryOverride)
	if wrap == nil {
		return nil
	}
	if ok, err := wrap.populateNonStaticFields(obj, action, lister, namespaceStore); err != nil {
		// Panic on dev because we should always be able to parse the deployments
		utils.Should(err)
		return nil
	} else if !ok {
		return nil
	}
	return wrap
}

func newWrap(obj interface{}, kind, registryOverride string) *deploymentWrap {
	deployment, err := resources.NewDeploymentFromStaticResource(obj, kind, registryOverride)
	if err != nil || deployment == nil {
		return nil
	}
	return &deploymentWrap{
		Deployment:       deployment,
		registryOverride: registryOverride,
	}
}

func (w *deploymentWrap) populateK8sComponentIfNecessary(o *v1.Pod) *metav1.LabelSelector {
	if o.Namespace == kubeSystemNamespace {
		for _, labelKey := range k8sComponentLabelKeys {
			value, ok := o.Labels[labelKey]
			if !ok {
				continue
			}
			w.Id = getK8sComponentID(value)
			w.Name = fmt.Sprintf("static-%s-pods", value)
			w.Type = k8sStandalonePodType
			return &metav1.LabelSelector{
				MatchLabels: map[string]string{
					labelKey: value,
				},
			}
		}
	}
	return nil
}

func checkIfNewPodSpecRequired(podSpec *v1.PodSpec, pods []*v1.Pod) bool {
	containerSet := set.NewStringSet()
	for _, c := range podSpec.Containers {
		containerSet.Add(c.Name)
	}
	var updated bool
	for _, p := range pods {
		if p.GetDeletionTimestamp() != nil {
			continue
		}
		for _, c := range p.Spec.Containers {
			if containerSet.Contains(c.Name) {
				continue
			}
			updated = true
			containerSet.Add(c.Name)
			podSpec.Containers = append(podSpec.Containers, c)
		}
	}
	return updated
}

func (w *deploymentWrap) populateNonStaticFields(obj interface{}, action *central.ResourceAction, lister v1listers.PodLister, namespaceStore *namespaceStore) (bool, error) {
	w.original = obj
	objValue := reflect.Indirect(reflect.ValueOf(obj))
	spec := objValue.FieldByName("Spec")
	if !doesFieldExist(spec) {
		return false, fmt.Errorf("obj %+v does not have a Spec field", objValue)
	}

	var (
		podSpec       v1.PodSpec
		podLabels     map[string]string
		labelSelector *metav1.LabelSelector
		err           error
	)

	switch o := obj.(type) {
	case *openshift_appsv1.DeploymentConfig:
		if o.Spec.Template == nil {
			return false, fmt.Errorf("spec obj %+v does not have a Template field or is not a pointer pod spec", spec)
		}
		podLabels = o.Spec.Template.Labels
		podSpec = o.Spec.Template.Spec

		labelSelector, err = w.getLabelSelector(spec)
		if err != nil {
			return false, errors.Wrap(err, "error getting label selector")
		}

	// Pods don't have the abstractions that higher level objects have so maintain it's lifecycle independently
	case *v1.Pod:
		if o.Status.Phase == v1.PodSucceeded || o.Status.Phase == v1.PodFailed {
			*action = central.ResourceAction_REMOVE_RESOURCE
		}

		// Standalone Pods do not have a PodTemplate, like the other deployment
		// types do. So, we need to directly access the Pod's Spec field,
		// instead of looking for it inside a PodTemplate.
		podLabels = o.Labels
		labelSelector = w.populateK8sComponentIfNecessary(o)
	case *v1beta1.CronJob:
		// Cron jobs have a Job spec that then have a Pod Template underneath
		podLabels = o.Spec.JobTemplate.Spec.Template.GetLabels()
		podSpec = o.Spec.JobTemplate.Spec.Template.Spec
		labelSelector = o.Spec.JobTemplate.Spec.Selector
	default:
		podTemplate, err := resources.SpecToPodTemplateSpec(spec)
		if err != nil {
			return false, errors.Wrapf(err, "spec obj %+v cannot be converted to a pod template spec", spec)
		}
		podLabels = podTemplate.Labels
		podSpec = podTemplate.Spec

		labelSelector, err = w.getLabelSelector(spec)
		if err != nil {
			return false, errors.Wrap(err, "error getting label selector")
		}
	}

	labelSel, err := k8s.ToRoxLabelSelector(labelSelector)
	if err != nil {
		log.Warnf("Could not convert label selector: %v", err)
	}

	w.PodLabels = podLabels
	w.LabelSelector = labelSel
	w.AutomountServiceAccountToken = true
	if podSpec.AutomountServiceAccountToken != nil {
		w.AutomountServiceAccountToken = *podSpec.AutomountServiceAccountToken
	}

	w.populateNamespaceID(namespaceStore)

	if labelSelector == nil {
		labelSelector = &metav1.LabelSelector{
			MatchLabels: podLabels,
		}
	}

	if *action != central.ResourceAction_REMOVE_RESOURCE {
		// If we have a standalone pod, we cannot use the labels to try and select that pod so we must directly populate the pod data
		// We need to special case kube-proxy because we are consolidating it into a deployment
		if pod, ok := obj.(*v1.Pod); ok && w.Type != k8sStandalonePodType {
			w.populatePorts()
			w.populateDataFromPods(pod)
		} else {
			pods, err := w.getPods(labelSelector, lister)
			if err != nil {
				return false, err
			}
			if updated := checkIfNewPodSpecRequired(&podSpec, pods); updated {
				resources.NewDeploymentWrap(w.Deployment, w.registryOverride).PopulateDeploymentFromPodSpec(podSpec)
			}
			w.populatePorts()
			w.populateDataFromPods(pods...)
		}
	} else {
		w.populatePorts()
	}
	return true, nil
}

func (w *deploymentWrap) GetDeployment() *storage.Deployment {
	if w == nil {
		return nil
	}
	return w.Deployment
}

func matchesOwnerName(name, topLevelType string, p *v1.Pod) bool {
	// Edge case that happens for Standalone Pods
	if len(p.GetOwnerReferences()) == 0 {
		return true
	}
	var numExpectedDashes int
	switch topLevelType {
	case kubernetes.CronJob, kubernetes.Deployment, kubernetes.DeploymentConfig: // 2 dash in pod
		// nginx-deployment-86d59dd769-7gmsk we want nginx-deployment
		numExpectedDashes = 2
	case kubernetes.DaemonSet, kubernetes.StatefulSet, kubernetes.ReplicationController, kubernetes.ReplicaSet, kubernetes.Job: // 1 dash in pod
		// nginx-deployment-7gmsk we want nginx-deployment
		numExpectedDashes = 1
	default:
		log.Warnf("Currently do not handle top level owner type %q. Adding to top level object %q", topLevelType, name)
		// By default if we can't parse, then we'll hit the mis-attribution edge case, but I'd rather do that
		// then miss the pods altogether
		return true
	}
	spl := strings.Split(p.GetName(), "-")
	if len(spl) > numExpectedDashes {
		if name == strings.Join(spl[:len(spl)-numExpectedDashes], "-") {
			return true
		}
		// We were able to parse the name, but didn't find a match
		if len(p.GetName()) < labelMaxLength {
			return false
		}
	} else if len(p.GetName()) < labelMaxLength {
		// We should have been able to parse the name normally as it was < max length, but we were not
		log.Debugf("Could not parse pod %q with top level owner type %q", p.GetName(), topLevelType)
		return false
	}

	if name[0:len(name)-trailingUIDLen] == p.GetName()[0:len(p.GetName())-trailingUIDLen] {
		return true
	}

	if name == strings.Join(spl[:len(spl)-1], "-") {
		return true
	}

	log.Debugf("Could not parse pod %q with owner type %q", p.GetName(), topLevelType)
	return false
}

// Do cheap filtering on pod name based on name of higher level object (deployment, daemonset, etc)
func filterOnName(name, topLevelType string, pods []*v1.Pod) []*v1.Pod {
	filteredPods := pods[:0]
	for _, p := range pods {
		if matchesOwnerName(name, topLevelType, p) {
			filteredPods = append(filteredPods, p)
		}
	}
	return filteredPods
}

func (w *deploymentWrap) getPods(labelSelector *metav1.LabelSelector, lister v1listers.PodLister) ([]*v1.Pod, error) {
	compiledLabelSelector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, errors.Wrap(err, "could not compile label selector")
	}
	w.podSelector = compiledLabelSelector
	pods, err := lister.Pods(w.Namespace).List(w.podSelector)
	if err != nil {
		return nil, err
	}
	return filterOnName(w.Name, w.Type, pods), nil
}

func (w *deploymentWrap) populateDataFromPods(pods ...*v1.Pod) {
	w.pods = pods
	w.populateImageIDs(pods...)
	w.populateContainerInstances(pods...)
}

func (w *deploymentWrap) populateContainerInstances(pods ...*v1.Pod) {
	for _, p := range pods {
		for i, instance := range containerInstances(p) {
			// This check that the size is not greater is necessary, because pods can be in terminating as a deployment is updated
			// The deployment will still be managing the pods, but we want to take the new pod(s) as the source of truth
			if i >= len(w.Containers) {
				break
			}
			w.Containers[i].Instances = append(w.Containers[i].Instances, instance)
		}
	}
	// Create a stable ordering
	for _, c := range w.Containers {
		sort.SliceStable(c.Instances, func(i, j int) bool { return c.Instances[i].InstanceId.Id < c.Instances[j].InstanceId.Id })
	}
}

func (w *deploymentWrap) populateImageIDs(pods ...*v1.Pod) {
	// All containers have a container status
	// The downside to this is that if different pods have different versions then we will miss that fact that pods are running
	// different versions and clobber it. I've added a log to illustrate the clobbering so we can see how often it happens

	// Sort the w.Deployment.Containers by name and p.Status.ContainerStatuses by name
	// This is because the order is not guaranteed
	sort.SliceStable(w.Deployment.Containers, func(i, j int) bool {
		return w.Deployment.GetContainers()[i].Name < w.Deployment.GetContainers()[j].Name
	})

	// Sort the pods by time created as that pod will be most likely to have the most updated spec
	sort.SliceStable(pods, func(i, j int) bool {
		return pods[j].CreationTimestamp.Before(&pods[i].CreationTimestamp)
	})

	for _, p := range pods {
		sort.SliceStable(p.Status.ContainerStatuses, func(i, j int) bool {
			return p.Status.ContainerStatuses[i].Name < p.Status.ContainerStatuses[j].Name
		})
		sort.SliceStable(p.Spec.Containers, func(i, j int) bool {
			return p.Spec.Containers[i].Name < p.Spec.Containers[j].Name
		})
		for i, c := range p.Status.ContainerStatuses {
			if i >= len(w.Deployment.Containers) || i >= len(p.Spec.Containers) {
				// This should not happened, but could happen if w.Deployment.Containers and container status are out of sync
				break
			}

			// If there already is an image ID for the image then that implies that the name of the image was fully qualified
			// with an image digest. e.g. stackrox.io/main@sha256:xyz
			if w.Deployment.Containers[i].Image.GetId() != "" {
				continue
			}

			parsedName, err := imageUtils.GenerateImageFromStringWithOverride(p.Spec.Containers[i].Image, w.registryOverride)
			if err != nil {
				// This error will only happen if we could not parse the image, this is possible if the image in kubernetes is malformed
				// e.g. us.gcr.io/$PROJECT/xyz:latest is an example that we have seen
				continue
			}

			// If the pod spec image doesn't match the top level image, then it is an old spec so we should ignore its digest
			if parsedName.GetName().GetFullName() != w.Containers[i].Image.GetName().GetFullName() {
				continue
			}

			if digest := imageUtils.ExtractImageDigest(c.ImageID); digest != "" {
				w.Deployment.Containers[i].Image.Id = digest
				w.Deployment.Containers[i].Image.NotPullable = !imageUtils.IsPullable(c.ImageID)
			}
		}
	}
}

func (w *deploymentWrap) getLabelSelector(spec reflect.Value) (*metav1.LabelSelector, error) {
	s := spec.FieldByName("Selector")
	if !doesFieldExist(s) {
		return nil, nil
	}

	// Selector is of map type for replication controller
	if labelMap, ok := s.Interface().(map[string]string); ok {
		return &metav1.LabelSelector{
			MatchLabels: labelMap,
		}, nil
	}

	// All other resources uses labelSelector.
	if ls, ok := s.Interface().(*metav1.LabelSelector); ok {
		return ls, nil
	}

	return nil, fmt.Errorf("unable to get label selector for %+v", spec.Type())
}

func (w *deploymentWrap) populateNamespaceID(namespaceStore *namespaceStore) {
	if namespaceID, found := namespaceStore.lookupNamespaceID(w.GetNamespace()); found {
		w.NamespaceId = namespaceID
	} else {
		log.Errorf("no namespace ID found for namespace %s and deployment %q", w.GetNamespace(), w.GetName())
	}
}

func (w *deploymentWrap) populatePorts() {
	w.portConfigs = make(map[portRef]*storage.PortConfig)
	for _, c := range w.GetContainers() {
		for _, p := range c.GetPorts() {
			w.portConfigs[portRef{Port: intstr.FromInt(int(p.ContainerPort)), Protocol: v1.Protocol(p.Protocol)}] = p
			if p.Name != "" {
				w.portConfigs[portRef{Port: intstr.FromString(p.Name), Protocol: v1.Protocol(p.Protocol)}] = p
			}
		}
	}
}

func (w *deploymentWrap) toEvent(action central.ResourceAction) *central.SensorEvent {
	return &central.SensorEvent{
		Id:     w.GetId(),
		Action: action,
		Resource: &central.SensorEvent_Deployment{
			Deployment: w.Deployment,
		},
	}
}

func filterHostExposure(exposureInfos []*storage.PortConfig_ExposureInfo) (filtered []*storage.PortConfig_ExposureInfo, level storage.PortConfig_ExposureLevel) {
	for _, exposureInfo := range exposureInfos {
		if exposureInfo.GetLevel() != storage.PortConfig_HOST {
			continue
		}
		filtered = append(filtered, exposureInfo)
		level = storage.PortConfig_HOST
	}
	return
}

func (w *deploymentWrap) resetPortExposure() {
	for _, portCfg := range w.portConfigs {
		portCfg.ExposureInfos, portCfg.Exposure = filterHostExposure(portCfg.ExposureInfos)
	}
}

func (w *deploymentWrap) updatePortExposureFromStore(store *serviceStore) {
	w.resetPortExposure()

	svcs := store.getMatchingServices(w.Namespace, w.PodLabels)
	for _, svc := range svcs {
		w.updatePortExposure(svc)
	}
}

func (w *deploymentWrap) updatePortExposure(svc *serviceWrap) {
	if !svc.selector.Matches(labels.Set(w.PodLabels)) {
		return
	}

	for ref, exposureInfo := range svc.exposure() {
		portCfg := w.portConfigs[ref]
		if portCfg == nil {
			if ref.Port.Type == intstr.String {
				// named ports MUST be defined in the pod spec
				continue
			}
			portCfg = &storage.PortConfig{
				ContainerPort: ref.Port.IntVal,
				Protocol:      string(ref.Protocol),
			}
			w.Ports = append(w.Ports, portCfg)
			w.portConfigs[ref] = portCfg
		}

		portCfg.ExposureInfos = append(portCfg.ExposureInfos, exposureInfo)

		if containers.CompareExposureLevel(portCfg.Exposure, exposureInfo.GetLevel()) < 0 {
			portCfg.Exposure = exposureInfo.GetLevel()
		}
	}
}
