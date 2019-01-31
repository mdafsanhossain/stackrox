package data

import (
	"github.com/stackrox/rox/central/compliance/framework"
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/internalapi/compliance"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/set"
)

var (
	log = logging.LoggerForModule()
)

type repository struct {
	cluster     *storage.Cluster
	nodes       map[string]*storage.Node
	deployments map[string]*storage.Deployment

	alerts             []*storage.ListAlert
	networkPolicies    map[string]*storage.NetworkPolicy
	networkGraph       *v1.NetworkGraph
	policies           map[string]*storage.Policy
	images             []*storage.ListImage
	imageIntegrations  []*storage.ImageIntegration
	processIndicators  []*storage.ProcessIndicator
	networkFlows       []*storage.NetworkFlow
	notifiers          []*storage.Notifier
	categoryToPolicies map[string]set.StringSet // maps categories to policy set

	hostScrape map[string]*compliance.ComplianceReturn
}

func (r *repository) Cluster() *storage.Cluster {
	return r.cluster
}

func (r *repository) Nodes() map[string]*storage.Node {
	return r.nodes
}

func (r *repository) Deployments() map[string]*storage.Deployment {
	return r.deployments
}

func (r *repository) NetworkPolicies() map[string]*storage.NetworkPolicy {
	return r.networkPolicies
}

func (r *repository) NetworkGraph() *v1.NetworkGraph {
	return r.networkGraph
}

func (r *repository) Policies() map[string]*storage.Policy {
	return r.policies
}

func (r *repository) PolicyCategories() map[string]set.StringSet {
	return r.categoryToPolicies
}

func (r *repository) Images() []*storage.ListImage {
	return r.images
}

func (r *repository) ImageIntegrations() []*storage.ImageIntegration {
	return r.imageIntegrations
}

func (r *repository) ProcessIndicators() []*storage.ProcessIndicator {
	return r.processIndicators
}

func (r *repository) NetworkFlows() []*storage.NetworkFlow {
	return r.networkFlows
}

func (r *repository) Notifiers() []*storage.Notifier {
	return r.notifiers
}

func (r *repository) Alerts() []*storage.ListAlert {
	return r.alerts
}

func (r *repository) HostScraped() map[string]*compliance.ComplianceReturn {
	return r.hostScrape
}

func newRepository(domain framework.ComplianceDomain, scrapeResults map[string]*compliance.ComplianceReturn, factory *factory) (*repository, error) {
	r := &repository{}
	if err := r.init(domain, scrapeResults, factory); err != nil {
		return nil, err
	}
	return r, nil
}

func nodesByID(nodes []*storage.Node) map[string]*storage.Node {
	result := make(map[string]*storage.Node, len(nodes))
	for _, node := range nodes {
		result[node.GetId()] = node
	}
	return result
}

func deploymentsByID(deployments []*storage.Deployment) map[string]*storage.Deployment {
	result := make(map[string]*storage.Deployment, len(deployments))
	for _, deployment := range deployments {
		result[deployment.GetId()] = deployment
	}
	return result
}

func networkPoliciesByID(policies []*storage.NetworkPolicy) map[string]*storage.NetworkPolicy {
	result := make(map[string]*storage.NetworkPolicy, len(policies))
	for _, policy := range policies {
		result[policy.GetId()] = policy
	}
	return result
}

func policiesByName(policies []*storage.Policy) map[string]*storage.Policy {
	result := make(map[string]*storage.Policy, len(policies))
	for _, policy := range policies {
		result[policy.GetName()] = policy
	}
	return result
}

func policyCategories(policies []*storage.Policy) map[string]set.StringSet {
	result := make(map[string]set.StringSet, len(policies))
	for _, policy := range policies {
		if policy.Disabled {
			continue
		}
		for _, category := range policy.Categories {
			policySet, ok := result[category]
			if !ok {
				policySet = set.NewStringSet()
			}
			policySet.Add(policy.Name)
			result[category] = policySet
		}
	}
	return result
}

func expandFile(parent *compliance.File) map[string]*compliance.File {
	expanded := make(map[string]*compliance.File)
	for _, child := range parent.GetChildren() {
		childExpanded := expandFile(child)
		for k, v := range childExpanded {
			expanded[k] = v
		}
	}
	expanded[parent.GetPath()] = parent
	return expanded
}

func (r *repository) init(domain framework.ComplianceDomain, scrapeResults map[string]*compliance.ComplianceReturn, f *factory) error {
	r.cluster = domain.Cluster().Cluster()
	r.nodes = nodesByID(framework.Nodes(domain))

	deployments := framework.Deployments(domain)
	r.deployments = deploymentsByID(deployments)

	clusterID := r.cluster.GetId()
	networkPolicies, err := f.networkPoliciesStore.GetNetworkPolicies(&v1.GetNetworkPoliciesRequest{
		ClusterId: clusterID,
	})
	if err != nil {
		return err
	}
	r.networkPolicies = networkPoliciesByID(networkPolicies)

	r.networkGraph = f.networkGraphEvaluator.GetGraph(deployments, networkPolicies)

	policies, err := f.policyStore.GetPolicies()
	if err != nil {
		return err
	}

	r.policies = policiesByName(policies)
	r.categoryToPolicies = policyCategories(policies)

	r.images, err = f.imageStore.ListImages()
	if err != nil {
		return err
	}

	r.imageIntegrations, err = f.imageIntegrationStore.GetImageIntegrations(
		&v1.GetImageIntegrationsRequest{},
	)
	if err != nil {
		return err
	}

	r.processIndicators, err = f.processIndicatorStore.GetProcessIndicators()
	if err != nil {
		return err
	}

	flowStore := f.networkFlowStore.GetFlowStore(domain.Cluster().ID())
	r.networkFlows, _, err = flowStore.GetAllFlows()
	if err != nil {
		return err
	}

	r.notifiers, err = f.notifierStore.GetNotifiers(&v1.GetNotifiersRequest{})
	if err != nil {
		return err
	}

	r.alerts, err = f.alertStore.GetAlertStore()
	if err != nil {
		return err
	}

	// Flatten the files so we can do direct lookups on the nested values
	for _, n := range scrapeResults {
		totalNodeFiles := make(map[string]*compliance.File)
		for path, file := range n.GetFiles() {
			expanded := expandFile(file)
			for k, v := range expanded {
				totalNodeFiles[k] = v
			}
			totalNodeFiles[path] = file
		}
		n.Files = totalNodeFiles
	}

	r.hostScrape = scrapeResults

	return nil
}
