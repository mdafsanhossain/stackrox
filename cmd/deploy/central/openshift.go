package central

import (
	"github.com/stackrox/rox/generated/api/v1"
	kubernetesPkg "github.com/stackrox/rox/pkg/kubernetes"
)

func init() {
	Deployers[v1.ClusterType_OPENSHIFT_CLUSTER] = newOpenshift()
}

type openshift struct{}

func newOpenshift() deployer {
	return &openshift{}
}

var openshiftMonitoringOnPrem = []string{
	"openshift/monitoring/monitoring.sh",
	"openshift/monitoring/monitoring-rbac.yaml",
	"kubernetes/monitoring/monitoring.yaml",
	"kubernetes/monitoring/influxdb.conf",
}

func (o *openshift) Render(c Config) ([]*v1.File, error) {
	injectImageTags(&c)

	var err error
	c.K8sConfig.Registry, err = kubernetesPkg.GetResolvedRegistry(c.K8sConfig.MainImage)
	if err != nil {
		return nil, err
	}
	c.K8sConfig.MonitoringImage = generateMonitoringImage(c.K8sConfig.MainImage)

	filenames := []string{
		"kubernetes/central.yaml",
		"kubernetes/ca-setup.sh",
		"kubernetes/delete-ca.sh",
		"kubernetes/np.yaml",

		"openshift/central.sh",
		"openshift/central-rbac.yaml",
		"openshift/clairify.sh",
		"openshift/clairify.yaml",
		"openshift/image-setup.sh",
		"openshift/port-forward.sh",
		"openshift/route-setup.sh",
	}

	if c.K8sConfig.MonitoringType.OnPrem() {
		filenames = append(filenames, openshiftMonitoringOnPrem...)
		filenames = append(filenames, monitoringClient...)
	} else if c.K8sConfig.MonitoringType.StackRoxHosted() {
		filenames = append(filenames, monitoringClient...)
	}

	return renderFilenames(filenames, &c, "/data/assets/docker-auth.sh")
}

func (o *openshift) Instructions() string {
	return `To deploy:
  1. Unzip the deployment bundle.
  2. If you need to add additional trusted CAs, run ca-setup.sh.
  3. If monitoring is on-prem, run monitoring/monitoring.sh
  4. Run central.sh.
  5. If you want to run the StackRox Clairify scanner, run clairify.sh.
  6. Expose Central:
       a. Using a Route:        ./route-setup.sh
       b. Using a NodePort:     oc create -f np.yaml
       c. Using a port forward: ./port-forward.sh 8443`
}
