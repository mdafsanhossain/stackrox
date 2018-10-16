package docker

import (
	"fmt"
	"strings"

	manifestV1 "github.com/docker/distribution/manifest/schema1"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/heroku/docker-registry-client/registry"
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/images/utils"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/registries/types"
	"github.com/stackrox/rox/pkg/urlfmt"
)

var (
	log = logging.LoggerForModule()
)

// Creator provides the type and registries.Creator to add to the registries Registry.
func Creator() (string, func(integration *v1.ImageIntegration) (types.ImageRegistry, error)) {
	return "docker", func(integration *v1.ImageIntegration) (types.ImageRegistry, error) {
		reg, err := newRegistry(integration)
		return reg, err
	}
}

// Registry is the basic docker registry implementation
type Registry struct {
	cfg                   Config
	protoImageIntegration *v1.ImageIntegration

	client *registry.Registry

	url      string
	registry string // This is the registry portion of the image
}

// Config is the basic config for the docker registry
type Config struct {
	// Endpoint defines the Docker Registry URL
	Endpoint string
	// Username defines the Username for the Docker Registry
	Username string
	// Password defines the password for the Docker Registry
	Password string
	// Insecure defines if the registry should be insecure
	Insecure bool
}

// NewDockerRegistry creates a new instantiation of the docker registry
// TODO(cgorman) AP-386 - properly put the base docker registry into another pkg
func NewDockerRegistry(cfg Config, integration *v1.ImageIntegration) (*Registry, error) {
	url, err := urlfmt.FormatURL(cfg.Endpoint, urlfmt.HTTPS, urlfmt.NoTrailingSlash)
	if err != nil {
		return nil, err
	}
	// if the registryServer endpoint contains docker.io then the image will be docker.io/namespace/repo:tag
	registryServer := urlfmt.GetServerFromURL(url)
	if strings.Contains(cfg.Endpoint, "docker.io") {
		registryServer = "docker.io"
	}

	var client *registry.Registry
	if cfg.Insecure {
		client, err = registry.NewInsecure(url, cfg.Username, cfg.Password)
	} else {
		client, err = registry.New(url, cfg.Username, cfg.Password)
	}
	if err != nil {
		return nil, err
	}

	// Turn off the logs
	client.Logf = registry.Quiet

	return &Registry{
		url:                   url,
		registry:              registryServer,
		client:                client,
		cfg:                   cfg,
		protoImageIntegration: integration,
	}, nil
}

func newRegistry(integration *v1.ImageIntegration) (*Registry, error) {
	dockerConfig, ok := integration.IntegrationConfig.(*v1.ImageIntegration_Docker)
	if !ok {
		return nil, fmt.Errorf("Docker configuration required")
	}
	cfg := Config{
		Endpoint: dockerConfig.Docker.GetEndpoint(),
		Username: dockerConfig.Docker.GetUsername(),
		Password: dockerConfig.Docker.GetPassword(),
		Insecure: dockerConfig.Docker.GetInsecure(),
	}
	return NewDockerRegistry(cfg, integration)
}

// Match decides if the image is contained within this registry
func (r *Registry) Match(image *v1.Image) bool {
	return r.registry == image.GetName().GetRegistry()
}

// Global returns whether or not this registry is available from all clusters
func (r *Registry) Global() bool {
	return len(r.protoImageIntegration.GetClusters()) == 0
}

// Metadata returns the metadata via this registries implementation
func (r *Registry) Metadata(image *v1.Image) (*v1.ImageMetadata, error) {
	log.Infof("Getting metadata for image %s", image.GetName().GetFullName())
	if image == nil {
		return nil, nil
	}

	remote := image.GetName().GetRemote()

	digest, manifestType, err := r.client.ManifestDigest(remote, utils.Reference(image))
	if err != nil {
		return nil, fmt.Errorf("Failed to get the manifest digest : %s", err)
	}

	// If the image ID is empty, then populate with the digest from the manifest
	// This only applies in a situation with CI client
	if image.GetId() == "" {
		image.Id = digest.String()
	}

	switch manifestType {
	case manifestV1.MediaTypeManifest:
		v1Metadata, err := r.handleV1Manifest(remote, digest.String())
		if err != nil {
			return nil, err
		}
		return &v1.ImageMetadata{
			V1: v1Metadata,
		}, nil
	case manifestV1.MediaTypeSignedManifest:
		v1Metadata, err := r.handleV1SignedManifest(remote, digest.String())
		if err != nil {
			return nil, err
		}
		return &v1.ImageMetadata{
			V1: v1Metadata,
		}, nil
	case registry.MediaTypeManifestList:
		return r.handleV2ManifestList(remote, digest.String())
	case schema2.MediaTypeManifest:
		return r.handleV2Manifest(remote, digest.String())
	default:
		return nil, fmt.Errorf("unknown manifest type '%s'", manifestType)
	}
}

// Test tests the current registry and makes sure that it is working properly
func (r *Registry) Test() error {
	return r.client.Ping()
}

// Config returns the configuration of the docker registry
func (r *Registry) Config() *types.Config {
	return &types.Config{
		Username:         r.cfg.Username,
		Password:         r.cfg.Password,
		Insecure:         r.cfg.Insecure,
		URL:              r.url,
		RegistryHostname: r.registry,
	}
}
