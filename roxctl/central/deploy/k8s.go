package deploy

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/renderer"
	"github.com/stackrox/rox/pkg/roxctl"
	"github.com/stackrox/rox/pkg/roxctl/defaults"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stackrox/rox/roxctl/common/flags"
	"github.com/stackrox/rox/roxctl/common/util"
)

type persistentFlagsWrapper struct {
	*pflag.FlagSet
}

func (w *persistentFlagsWrapper) UInt32Var(p *uint32, name string, value uint32, usage string, groups ...string) {
	w.FlagSet.Uint32Var(p, name, value, usage)
	utils.Must(w.SetAnnotation(name, groupAnnotationKey, groups))
}

func (w *persistentFlagsWrapper) StringVar(p *string, name, value, usage string, groups ...string) {
	w.StringVarP(p, name, "", value, usage, groups...)
}

func (w *persistentFlagsWrapper) StringVarP(p *string, name, shorthand, value, usage string, groups ...string) {
	w.FlagSet.StringVarP(p, name, shorthand, value, usage)
	utils.Must(w.SetAnnotation(name, groupAnnotationKey, groups))
}

func (w *persistentFlagsWrapper) BoolVar(p *bool, name string, value bool, usage string, groups ...string) {
	w.FlagSet.BoolVar(p, name, value, usage)
	utils.Must(w.SetAnnotation(name, groupAnnotationKey, groups))
}

func (w *persistentFlagsWrapper) Var(value pflag.Value, name, usage string, groups ...string) {
	w.FlagSet.Var(value, name, usage)
	utils.Must(w.SetAnnotation(name, groupAnnotationKey, groups))
}

func orchestratorCommand(shortName, longName string) *cobra.Command {
	c := &cobra.Command{
		Use: shortName,
		Annotations: map[string]string{
			categoryAnnotation: "Enter orchestrator",
		},
		RunE: util.RunENoArgs(func(*cobra.Command) error {
			return errors.New("storage type must be specified")
		}),
	}
	if !roxctl.InMainImage() {
		c.PersistentFlags().Var(newOutputDir(&cfg.OutputDir), "output-dir", "the directory to output the deployment bundle to")
	}
	return c
}

func k8sBasedOrchestrator(k8sConfig *renderer.K8sConfig, shortName, longName string, cluster storage.ClusterType) *cobra.Command {
	c := orchestratorCommand(shortName, longName)
	c.PersistentPreRun = func(*cobra.Command, []string) {
		cfg.K8sConfig = k8sConfig
		cfg.ClusterType = cluster
	}

	c.AddCommand(externalVolume())
	c.AddCommand(hostPathVolume())
	c.AddCommand(noVolume())

	flagWrap := &persistentFlagsWrapper{FlagSet: c.PersistentFlags()}

	// Adds k8s specific flags
	flagWrap.StringVarP(&k8sConfig.MainImage, "main-image", "i", defaults.MainImage(), "main image to use", "central")
	flagWrap.BoolVar(&k8sConfig.OfflineMode, "offline", false, "whether to run StackRox in offline mode, which avoids reaching out to the Internet", "central")

	// Monitoring Flags

	flagWrap.StringVar(&k8sConfig.Monitoring.Endpoint, "monitoring-endpoint", "monitoring.stackrox:443", "monitoring endpoint", "monitoring", "monitoring-type=on-prem")
	utils.Must(flagWrap.MarkHidden("monitoring-endpoint"))

	flagWrap.Var(&monitoringWrapper{Monitoring: &k8sConfig.Monitoring.Type}, "monitoring-type", "where to host the monitoring (on-prem, none)", "monitoring")
	utils.Must(flagWrap.MarkHidden("monitoring-type"))

	flagWrap.StringVarP(&k8sConfig.Monitoring.Password, "monitoring-password", "p", "", "a monitoring password (default: autogenerated)", "monitoring", "monitoring-type=on-prem")
	utils.Must(
		flagWrap.SetAnnotation("monitoring-password", flags.PasswordKey, []string{"true"}))
	utils.Must(flagWrap.MarkHidden("monitoring-password"))

	flagWrap.StringVar(&k8sConfig.MonitoringImage, "monitoring-image", "", "monitoring image to use (default: same repository as main)", "monitoring", "monitoring-type=on-prem")
	utils.Must(flagWrap.MarkHidden("monitoring-image"))

	// Monitoring Persistence flags
	flagWrap.Var(&flags.PersistenceTypeWrapper{PersistenceType: &k8sConfig.Monitoring.PersistenceType}, "monitoring-persistence-type", "monitoring persistence type (none, hostpath, pvc)", "monitoring", "monitoring-type=on-prem")
	utils.Must(flagWrap.MarkHidden("monitoring-persistence-type"))

	flagWrap.StringVar(&k8sConfig.Monitoring.External.Name, "monitoring-persistence-name", "monitoring-db", "external volume name", "monitoring", "monitoring-type=on-prem", "monitoring-persistence-type=pvc")
	utils.Must(flagWrap.MarkHidden("monitoring-persistence-name"))

	flagWrap.StringVar(&k8sConfig.Monitoring.External.StorageClass, "monitoring-persistence-storage-class", "", "monitoring storage class name (optional if you have a default StorageClass configured)", "monitoring", "monitoring-type=on-prem", "monitoring-persistence-type=pvc")
	utils.Must(flagWrap.MarkHidden("monitoring-persistence-storage-class"))

	flagWrap.StringVar(&k8sConfig.Monitoring.HostPath.HostPath, "monitoring-persistence-hostpath", "/var/lib/stackrox/monitoring", "monitoring path on the host", "monitoring", "monitoring-type=on-prem", "monitoring-persistence-type=hostpath")
	utils.Must(flagWrap.MarkHidden("monitoring-persistence-hostpath"))

	flagWrap.StringVar(&k8sConfig.Monitoring.HostPath.NodeSelectorKey, "monitoring-node-selector-key", "", "monitoring node selector key (e.g. kubernetes.io/hostname)", "monitoring", "monitoring-type=on-prem", "monitoring-persistence-type=hostpath")
	utils.Must(flagWrap.MarkHidden("monitoring-node-selector-key"))

	flagWrap.StringVar(&k8sConfig.Monitoring.HostPath.NodeSelectorValue, "monitoring-node-selector-value", "", "monitoring node selector value", "monitoring", "monitoring-type=on-prem", "monitoring-persistence-type=hostpath")
	utils.Must(flagWrap.MarkHidden("monitoring-node-selector-value"))

	flagWrap.StringVar(&k8sConfig.ScannerImage, "scanner-image", defaults.ScannerImage(), "Scanner image to use", "scanner")
	flagWrap.StringVar(&k8sConfig.ScannerDBImage, "scanner-db-image", defaults.ScannerDBImage(), "Scanner DB image to use", "scanner")

	flagWrap.BoolVar(&k8sConfig.EnableTelemetry, "enable-telemetry", true, "whether to enable telemetry", "central")

	return c
}

func newK8sConfig() *renderer.K8sConfig {
	return &renderer.K8sConfig{
		Monitoring: renderer.MonitoringConfig{
			HostPath: &renderer.HostPathPersistence{},
			External: &renderer.ExternalPersistence{},
		},
	}
}

func k8s() *cobra.Command {
	k8sConfig := newK8sConfig()
	c := k8sBasedOrchestrator(k8sConfig, "k8s", "Kubernetes", storage.ClusterType_KUBERNETES_CLUSTER)
	flagWrap := &persistentFlagsWrapper{FlagSet: c.PersistentFlags()}

	flagWrap.Var(&loadBalancerWrapper{LoadBalancerType: &k8sConfig.LoadBalancerType}, "lb-type", "the method of exposing Central (lb, np, none)", "central")

	flagWrap.Var(&fileFormatWrapper{DeploymentFormat: &k8sConfig.DeploymentFormat}, "output-format", "the deployment tool to use (kubectl, helm)", "central")

	flagWrap.Var(&loadBalancerWrapper{LoadBalancerType: &k8sConfig.Monitoring.LoadBalancerType}, "monitoring-lb-type", "the method of exposing Monitoring (lb, np, none)", "monitoring", "monitoring-type=on-prem")
	utils.Must(flagWrap.MarkHidden("monitoring-lb-type"))

	return c
}

func openshift() *cobra.Command {
	k8sConfig := newK8sConfig()
	c := k8sBasedOrchestrator(k8sConfig, "openshift", "Openshift", storage.ClusterType_OPENSHIFT_CLUSTER)

	flagWrap := &persistentFlagsWrapper{FlagSet: c.PersistentFlags()}

	flagWrap.Var(&loadBalancerWrapper{LoadBalancerType: &k8sConfig.LoadBalancerType}, "lb-type", "the method of exposing Central (route, lb, np, none)", "central")

	flagWrap.Var(&loadBalancerWrapper{LoadBalancerType: &k8sConfig.Monitoring.LoadBalancerType}, "monitoring-lb-type", "the method of exposing Monitoring (route, lb, np, none)", "monitoring", "monitoring-type=on-prem")
	utils.Must(flagWrap.MarkHidden("monitoring-lb-type"))

	return c
}
